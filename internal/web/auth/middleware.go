package auth

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/models"
	"Derzhavnaya/internal/web/types"
	"context"
	"net/http"
)

func LoadUser(queries *db.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			dbUser, err := queries.GetUserBySessionID(r.Context(), cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user := models.NewUserFromDB(dbUser)
			ctx := context.WithValue(r.Context(), types.UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
