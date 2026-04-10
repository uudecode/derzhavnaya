package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // Нужен для goose
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"Derzhavnaya/internal/config"
	"Derzhavnaya/schema"
)

type gooseLogger struct{}

func (l *gooseLogger) Printf(format string, v ...interface{}) {
	log.Info().Msgf(strings.TrimRight(fmt.Sprintf(format, v...), "\n"))
}

func (l *gooseLogger) Fatalf(format string, v ...any) {
	log.Fatal().Msgf(strings.TrimRight(fmt.Sprintf(format, v...), "\n"))
}

func Prepare(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := cfg.Database.DSN()

	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("migrations: %w", err)
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pool connect: %w", err)
	}

	if cfg.InitialAdmin.Email != "" {
		if err := seedAdmin(ctx, pool, cfg.InitialAdmin); err != nil {
			log.Error().Err(err).Msg("failed to seed initial admin")
		}
	}

	return pool, nil
}

func runMigrations(dsn string) error {
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	goose.SetLogger(&gooseLogger{})
	goose.SetBaseFS(schema.MigrationsFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(sqlDB, "migrations")
}

func seedAdmin(ctx context.Context, pool *pgxpool.Pool, adm config.AdminConfig) error {
	queries := New(pool)

	hash, err := bcrypt.GenerateFromPassword([]byte(adm.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return queries.CreateAdminIfNotExist(ctx, CreateAdminIfNotExistParams{
		Email:    adm.Email,
		Password: string(hash),
	})
}
