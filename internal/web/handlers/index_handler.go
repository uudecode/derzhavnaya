package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/translator"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IndexHandler struct {
	BaseHandler
}

func NewIndexHandler(queries *db.Queries, cfg *config.Config, renderer *render.Engine, trans *translator.Translator) *IndexHandler {
	return &IndexHandler{
		BaseHandler: BaseHandler{
			DB:         queries,
			Renderer:   renderer,
			Cfg:        cfg,
			Translator: trans,
		},
	}
}

func (h *IndexHandler) Register(r chi.Router) {
	r.Get("/", h.Index)
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"MainIconURL": h.Cfg.S3.PublicBaseURL + "/icons/derzhavnaya_main.jpg",
		"CurrentPath": "/",
	}

	h.RenderPage(w, r, "index.html", data)
}
