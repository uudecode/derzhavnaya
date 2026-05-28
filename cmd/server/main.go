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

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if len(os.Args) > 1 && os.Args[1] == "import-glossary" {
		runGlossaryImportCLI()
		return
	}

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

func runGlossaryImportCLI() {
	app := fx.New(
		config.Module,
		logger.Module,
		db.Module,

		fx.Invoke(func(pool *pgxpool.Pool, sd fx.Shutdowner) {
			log.Info().Msg("Starting glossary import from CLI...")
			if err := tools.RunGlossaryImport(context.Background(), pool, "/glossary.json"); err != nil {
				log.Error().Err(err).Msg("Glossary import failed")
				_ = sd.Shutdown(fx.ExitCode(1))
				return
			}

			log.Info().Msg("Glossary import completed successfully!")
			_ = sd.Shutdown(fx.ExitCode(0))
		}),

		fx.NopLogger,
	)

	app.Run()
}
