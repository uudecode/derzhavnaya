package main

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/server"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/health"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load config")
	}

	level, err := zerolog.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	if err := health.CheckPostgres(cfg.Database); err != nil {
		log.Fatal().Err(err).Msg("Postgres check failed")
	}

	if err := health.CheckS3(cfg.S3); err != nil {
		log.Fatal().Err(err).Msg("S3 check failed")
	}

	ctx := context.Background()

	pool, err := db.Prepare(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("database bootstrap failed")
	}
	defer pool.Close()
	srv := server.NewServer(cfg)
	log.Info().Msgf("Starting server on :%d", cfg.App.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), srv.Routes()); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
