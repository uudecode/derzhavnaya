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

func NewDatabasePool(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := Prepare(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return pool, nil
}

func NewDatabasePoolFx(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := NewDatabasePool(cfg)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
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
