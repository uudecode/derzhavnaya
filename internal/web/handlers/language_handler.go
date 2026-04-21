package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LanguageHandler struct{}

func NewLanguageHandler() *LanguageHandler {
	return &LanguageHandler{}
}

func (h *LanguageHandler) Register(r chi.Router) {
	r.Get("/set-lang/{lang}", SetLanguage)
}

func SetLanguage(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	http.SetCookie(w, &http.Cookie{
		Name:     "lang",
		Value:    lang,
		Path:     "/",
		MaxAge:   31536000, // 1 год
		HttpOnly: true,
	})
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
