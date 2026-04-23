package main

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/health"
	"Derzhavnaya/internal/logger"
	"Derzhavnaya/internal/tools"
	"Derzhavnaya/internal/translation"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/handlers"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/server"
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "import-glossary":
			cfg, err := config.Load()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to load config")
				return
			}
			pool, err := db.NewDatabasePool(cfg)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create database pool")
				return
			}
			defer pool.Close()
			if err := tools.RunGlossaryImport(context.Background(), pool, "/glossary.json"); err != nil {
				log.Fatal().Err(err).Msg("glossary import failed")
			}
			return
		}
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fx.New(
		config.Module,
		logger.Module,
		health.Module,
		db.Module,
		auth.Module,
		render.Module,
		handlers.Module,
		server.Module,
		translation.Module,
		fx.Invoke(
			translation.StartTranslationWorker,
		),
	)

	app.Run()
}
