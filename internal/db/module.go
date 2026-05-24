package db

import (
	"Derzhavnaya/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDatabasePoolFx),
	fx.Provide(NewQueries),
)

func NewDatabasePoolFx(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := cfg.Database.DSN()

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool structure: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info().Msg("Running database migrations...")
			if err := runMigrations(dsn); err != nil {
				return fmt.Errorf("migrations failed on start: %w", err)
			}

			log.Info().Msg("Pinging database pool...")
			if err := pool.Ping(ctx); err != nil {
				return fmt.Errorf("database ping failed on start: %w", err)
			}

			if cfg.InitialAdmin.Email != "" {
				log.Info().Msg("Seeding initial admin...")
				if err := seedAdmin(ctx, pool, cfg.InitialAdmin); err != nil {
					log.Error().Err(err).Msg("failed to seed initial admin")
				}
			}

			log.Info().Msg("Database is fully ready and initialized.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Closing database pool...")
			pool.Close()
			return nil
		},
	})

	return pool, nil
}

func NewQueries(pool *pgxpool.Pool) *Queries {
	return New(pool)
}
