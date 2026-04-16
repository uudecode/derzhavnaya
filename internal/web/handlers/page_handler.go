package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/types"
	"context"

	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PageHandler struct {
	DB       *db.Queries
	Cfg      *config.Config
	Renderer *render.Engine
}

func NewPageHandler(queries *db.Queries, cfg *config.Config, renderer *render.Engine) *PageHandler {
	return &PageHandler{
		DB:       queries,
		Cfg:      cfg,
		Renderer: renderer,
	}
}

func (h *PageHandler) Register(r chi.Router) {
	r.Get("/", h.Index)
}

func (h *PageHandler) Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"MainIconURL": h.Cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
		"CurrentPath": "/",
	}

	h.Renderer.Render(w, r, "index.html", data)
}

func (h *PageHandler) MenuMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := auth.FromContext(r.Context())
		all, _ := h.DB.GetActiveMenuItems(r.Context())

		menu := auth.FilterMenuByRole(all, user.Role)

		ctx := context.WithValue(r.Context(), types.MenuKey, menu)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
