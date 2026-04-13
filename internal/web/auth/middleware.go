package auth

import (
	"Derzhavnaya/internal/db"
	"context"
	"net/http"
)

// Ключ для контекста (типизированный, чтобы избежать коллизий)
type contextKey string

const UserKey contextKey = "user"

func LoadUser(queries *db.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Достаем куку сессии
			cookie, err := r.Cookie("session_id")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Проверяем сессию в БД (твой запрос GetUserBySessionID)
			// r.Context() передает сигналы отмены (если юзер закрыл вкладку)
			user, err := queries.GetUserBySessionID(r.Context(), cookie.Value)
			if err != nil {
				// Если сессия не найдена или протухла, просто идем дальше
				next.ServeHTTP(w, r)
				return
			}

			// Кладем юзера в "карман" запроса (контекст)
			ctx := context.WithValue(r.Context(), UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Вспомогательная функция для хендлеров, чтобы удобно доставать юзера
func FromContext(ctx context.Context) (db.WebUser, bool) {
	user, ok := ctx.Value(UserKey).(db.WebUser)
	return user, ok
}
