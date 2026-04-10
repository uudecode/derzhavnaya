package server

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/web"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg *config.Config
	// сюда же потом добавим пул базы: db *pgxpool.Pool
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	staticFS, err := fs.Sub(web.Static, "static")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create static subfs")
	}
	fileServer := http.FileServer(http.FS(staticFS))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.Static, "static/favicon.ico")
	})
	r.Get("/", s.handleIndex)

	return r
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(web.Templates, "templates/*.html")
	if err != nil {
		log.Error().Err(err).Msg("failed to parse templates")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"MainIconURL": s.cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
	}

	// 2. Явно говорим, какой шаблон основной
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Error().Err(err).Msg("template execution failed")
	}
}
