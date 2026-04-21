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
		fx.Annotate(NewIndexHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
		fx.Annotate(NewQuestionsHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
		fx.Annotate(NewGalleryHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
		fx.Annotate(NewLanguageHandler, fx.As(new(Handler)), fx.ResultTags(`group:"handlers"`)),
	),
)
