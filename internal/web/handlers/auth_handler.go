package handlers

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/render"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB       *db.Queries
	Limiter  *auth.RateLimiter
	Renderer *render.Engine
}

func NewAuthHandler(queries *db.Queries, limiter *auth.RateLimiter, renderer *render.Engine) *AuthHandler {
	return &AuthHandler{
		DB:       queries,
		Limiter:  limiter,
		Renderer: renderer,
	}
}

func (h *AuthHandler) Register(r chi.Router) {
	r.Get("/login", h.LoginGet)
	r.Post("/login", h.LoginPost)
	r.Post("/logout", h.LogoutPost)
}
func (h *AuthHandler) LoginGet(w http.ResponseWriter, r *http.Request) {
	backURL := r.Referer()
	if backURL == "" {
		backURL = "/"
	}

	h.Renderer.Render(w, r, "login.html", map[string]any{
		"BackURL": backURL,
	})
}

func (h *AuthHandler) LoginPost(w http.ResponseWriter, r *http.Request) {
	renderError := func(msg string, code int, emailAddr string) {
		h.Renderer.Render(w, r, "login.html", map[string]any{
			"Error": msg,
			"Email": emailAddr,
		})
	}
	ip := auth.GetIP(r)
	if !h.Limiter.Allow(ip) {
		renderError("Слишком много попыток. Попробуйте позже.", http.StatusTooManyRequests, "")
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// 1. Ищем юзера (GetUserByEmail)
	user, err := h.DB.GetUserByEmail(r.Context(), email)
	if err != nil {
		renderError("Неверный логин или пароль", http.StatusUnauthorized, email)
		return
	}

	// 2. Проверяем хеш пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		renderError("Неверный логин или пароль", http.StatusUnauthorized, email)
		return
	}

	// 3. Создаем сессию (CreateSession)
	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = h.DB.CreateSession(r.Context(), db.CreateSessionParams{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true, // Недоступна для JS
		Secure:   true, // Только HTTPS (Traefik терминирует SSL, так что ок)
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) LogoutPost(w http.ResponseWriter, r *http.Request) {
	// В идеале тут еще надо вызвать удаление из БД по ID из куки
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err = h.DB.DeleteSession(r.Context(), cookie.Value)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete session")
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Удаляет куку немедленно
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
