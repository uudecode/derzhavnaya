package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/auth"
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
	user, _ := auth.FromContext(r.Context())
	allMenuItems, _ := h.DB.GetActiveMenuItems(r.Context())
	menuItems := h.filterMenu(allMenuItems, user.Role)

	data := map[string]any{
		"MainIconURL": h.Cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
		"MenuItems":   menuItems,
		"CurrentPath": "/",
	}

	h.Renderer.Render(w, r, "index.html", data)
}

func (h *PageHandler) filterMenu(items []db.WebMenuItem, userRole string) []db.WebMenuItem {
	var filtered []db.WebMenuItem
	for _, item := range items {
		if !item.Role.Valid || item.Role.String == "" {
			filtered = append(filtered, item)
			continue
		}

		if item.Role.String == userRole {
			filtered = append(filtered, item)
			continue
		}

		if userRole == "ADMIN" {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
