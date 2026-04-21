package handlers

import (
	"Derzhavnaya/internal/models"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/types"
	"context"
	"net/http"
)

func (h *IndexHandler) LayoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(types.UserKey).(models.User)
		all, err := h.DB.GetActiveMenuItems(r.Context())
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		menu := auth.FilterMenuByRole(all, user.Role)
		ctx := context.WithValue(r.Context(), types.MenuKey, menu)

		currentPath := r.URL.Path
		ctx = context.WithValue(ctx, types.PathKey, currentPath)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		langCookie, err := r.Cookie(string(types.LangKey))
		lang := types.RU

		if err == nil {
			parsedLang := types.Locale(langCookie.Value)

			switch parsedLang {
			case types.EN, types.RU, types.FR:
				lang = parsedLang
			default:
				lang = types.RU
			}
		}

		ctx := context.WithValue(r.Context(), types.LangKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
