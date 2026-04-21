package server

import (
	"Derzhavnaya/internal/web/translator"
	"context"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/handlers"
	"Derzhavnaya/web"
)

var Module = fx.Options(
	fx.Provide(NewServer,
		translator.NewTranslator),
	fx.Invoke(StartHTTPServer),
)

type Server struct {
	mux *chi.Mux
	cfg *config.Config
}

type ServerParams struct {
	fx.In

	Cfg        *config.Config
	DB         *db.Queries
	Translator *translator.Translator
	Handlers   []handlers.Handler `group:"handlers"`
}

func NewServer(p ServerParams) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(auth.LoadUser(p.DB))
	r.Use(handlers.LanguageMiddleware)

	r.Use(func(next http.Handler) http.Handler {
		return nosurf.New(next)
	})

	staticFS, _ := fs.Sub(web.Static, "static")
	fileServer := http.FileServer(http.FS(staticFS))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.Static, "static/favicon.ico")
	})

	var pageH *handlers.IndexHandler
	for _, h := range p.Handlers {
		if ph, ok := h.(*handlers.IndexHandler); ok {
			pageH = ph
			break
		}
	}

	for _, h := range p.Handlers {
		if _, ok := h.(*handlers.AuthHandler); ok {
			h.Register(r)
			continue
		}

		r.Group(func(r chi.Router) {
			if pageH != nil {
				r.Use(pageH.LayoutMiddleware)
			}
			h.Register(r)
		})
	}

	return &Server{
		mux: r,
		cfg: p.Cfg,
	}
}

func StartHTTPServer(lc fx.Lifecycle, srv *Server) {
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", srv.cfg.App.Port),
		Handler: srv.mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info().Msgf("Starting server on %s", httpSrv.Addr)
			go func() {
				if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Error().Err(err).Msg("Server failed")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Shutting down server...")
			return httpSrv.Shutdown(ctx)
		},
	})
}
