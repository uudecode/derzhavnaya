package server

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/handlers"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/web"

	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/justinas/nosurf"
	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg         *config.Config
	renderer    *render.Engine
	db          *db.Queries
	pool        *pgxpool.Pool
	authHandler *handlers.AuthHandler
	pageHandler *handlers.PageHandler
}

func NewServer(cfg *config.Config, pool *pgxpool.Pool) *Server {
	renderer := render.NewEngine()
	queries := db.New(pool)
	limiter := auth.NewRateLimiter()
	authH := &handlers.AuthHandler{DB: queries, Limiter: limiter, Renderer: renderer}
	pageH := &handlers.PageHandler{DB: queries, Cfg: cfg, Renderer: renderer}
	return &Server{
		cfg:         cfg,
		pool:        pool,
		db:          queries,
		renderer:    renderer,
		authHandler: authH,
		pageHandler: pageH,
	}
}

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(auth.LoadUser(s.db))
	csrfHandler := nosurf.New(r)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // Поставь true, когда будет HTTPS/прод
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(func(next http.Handler) http.Handler {
		return nosurf.New(next)
	})
	staticFS, err := fs.Sub(web.Static, "static")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create static subfs")
	}
	fileServer := http.FileServer(http.FS(staticFS))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.Static, "static/favicon.ico")
	})
	r.Get("/login", s.authHandler.LoginGet)
	r.Post("/login", s.authHandler.LoginPost)
	r.Post("/logout", s.authHandler.LogoutPost)
	r.Get("/", s.pageHandler.Index)

	return r
}
