package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/translator"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type GalleryHandler struct {
	BaseHandler
}

func NewGalleryHandler(queries *db.Queries, cfg *config.Config, renderer *render.Engine, trans *translator.Translator) *GalleryHandler {
	return &GalleryHandler{
		BaseHandler: BaseHandler{
			DB:         queries,
			Renderer:   renderer,
			Cfg:        cfg,
			Translator: trans,
		},
	}
}

func (h *GalleryHandler) Register(r chi.Router) {
	r.Get("/gallery*", h.HandleGallery)
}

func (h *GalleryHandler) HandleGallery(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "*")
	segments := strings.Split(strings.Trim(path, "/"), "/")
	log.Info().Msgf("segments %+v", segments)
	data := map[string]interface{}{}
	h.RenderPage(w, r, "gallery.html", data)
}
