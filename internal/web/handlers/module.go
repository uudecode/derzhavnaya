package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Handler interface {
	Register(r chi.Router)
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewAuthHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
		fx.Annotate(NewPageHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
	),
)
