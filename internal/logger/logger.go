package logger

import (
	"Derzhavnaya/internal/config"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(Setup),
)

func Setup(cfg *config.Config) {
	level, err := zerolog.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	// Здесь же можно настроить таймстампы или формат вывода (JSON vs Console)
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
