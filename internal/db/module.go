package db

import (
	"Derzhavnaya/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDatabasePool),
	fx.Provide(NewQueries),
)

func NewDatabasePool(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := Prepare(context.Background(), cfg)
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
