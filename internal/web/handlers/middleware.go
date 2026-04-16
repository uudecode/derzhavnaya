package handlers

import (
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/types"
	"context"
	"net/http"
)

func (h *PageHandler) LayoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := auth.FromContext(r.Context())

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
