package main

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/health"
	"Derzhavnaya/internal/logger"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/handlers"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/server"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func main() {
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
	)

	app.Run()
}
