package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/auth"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/translator"
	"Derzhavnaya/internal/web/viewmodel"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	BaseHandler
	Limiter *auth.RateLimiter
}

func NewAuthHandler(queries *db.Queries, cfg *config.Config, renderer *render.Engine, trans *translator.Translator, limiter *auth.RateLimiter) *AuthHandler {
	return &AuthHandler{
		BaseHandler: BaseHandler{
			DB:         queries,
			Renderer:   renderer,
			Cfg:        cfg,
			Translator: trans,
		},
		Limiter: limiter,
	}
}

func (h *AuthHandler) Register(r chi.Router) {
	r.Get("/login", h.LoginGet)
	r.Post("/login", h.LoginPost)
	r.Post("/logout", h.LogoutPost)
}
func (h *AuthHandler) LoginGet(w http.ResponseWriter, r *http.Request) {
	h.RenderPage(w, r, "login.html", viewmodel.AuthView{
		BackURL: backURL(r),
	})
}

func (h *AuthHandler) LoginPost(w http.ResponseWriter, r *http.Request) {
	renderError := func(msg string, code int, emailAddr string) {
		h.RenderPage(w, r, "login.html", viewmodel.AuthView{
			BackURL: backURL(r),
			Error:   msg,
			Email:   emailAddr,
		})
	}
	ip := auth.GetIP(r)
	if !h.Limiter.Allow(ip) {
		renderError(h.T(r, "auth.error.too_many_attempts"), http.StatusTooManyRequests, "")
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.DB.GetUserByEmail(r.Context(), email)
	if err != nil {
		renderError(h.T(r, "auth.error.invalid_credentials"), http.StatusUnauthorized, "")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		renderError(h.T(r, "auth.error.invalid_credentials"), http.StatusUnauthorized, "")
		return
	}

	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = h.DB.CreateSession(r.Context(), db.CreateSessionParams{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		http.Error(w, h.T(r, "error.server_error"), http.StatusInternalServerError)
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
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func backURL(r *http.Request) string {
	backURL := r.Referer()
	if backURL == "" {
		backURL = "/"
	}
	return backURL
}
