package health

import (
	"Derzhavnaya/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(RunChecks),
)

func RunChecks(cfg *config.Config, pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("health check failed: database is not responding: %w", err)
	}

	if err := CheckS3(cfg.S3); err != nil {
		return fmt.Errorf("health check failed: s3 storage is unreachable: %w", err)
	}

	return nil
}
