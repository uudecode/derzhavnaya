package handlers

import (
	"Derzhavnaya/internal/models"
	"Derzhavnaya/internal/web/types"
	"net/http"
)

func GetLangFromCtx(r *http.Request) types.Locale {
	val := r.Context().Value(types.LangKey)
	if val == nil {
		return types.RU
	}
	return val.(types.Locale)
}

func GetUserFromContext(r *http.Request) (models.User, bool) {
	user, ok := r.Context().Value(types.UserKey).(models.User)
	return user, ok
}

func GetMenuFromContext(r *http.Request) ([]models.MenuItem, bool) {
	menu, ok := r.Context().Value(types.MenuKey).([]models.MenuItem)
	return menu, ok
}
