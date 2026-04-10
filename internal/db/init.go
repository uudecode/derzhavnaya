package db

import (
	"Derzhavnaya/schema"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // Нужен для sql.Open
	"github.com/pressly/goose/v3"
)

func RunMigrations(dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(schema.MigrationsFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
