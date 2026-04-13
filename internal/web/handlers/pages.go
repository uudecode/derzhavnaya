package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"net/http"
)

type PageHandler struct {
	DB       *db.Queries
	Cfg      *config.Config
	Renderer *render.Engine
}

func (h *PageHandler) Index(w http.ResponseWriter, r *http.Request) {
	menuItems, _ := h.DB.GetActiveMenuItems(r.Context())

	data := map[string]any{
		"MainIconURL": h.Cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
		"MenuItems":   menuItems,
		"CurrentPath": "/",
	}

	h.Renderer.Render(w, r, "index.html", data)
}
