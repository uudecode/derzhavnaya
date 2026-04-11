package server

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/web"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg  *config.Config
	tmpl *template.Template
	db   *db.Queries
	pool *pgxpool.Pool
}

func NewServer(cfg *config.Config, pool *pgxpool.Pool) *Server {
	tmpl := template.Must(template.New("").ParseFS(web.Templates, "templates/*.html"))
	return &Server{
		cfg:  cfg,
		tmpl: tmpl,
		pool: pool,
		db:   db.New(pool),
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
	menuItems, err := s.db.GetActiveMenuItems(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch menu")
	}

	data := map[string]any{
		"MainIconURL": s.cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
		"MenuItems":   menuItems,
		"CurrentPath": "/",
	}

	err = s.tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Error().Err(err).Msg("template execution failed")
	}
}
