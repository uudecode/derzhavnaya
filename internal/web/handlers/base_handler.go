package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/translator"
	"Derzhavnaya/internal/web/viewmodel"
	"net/http"

	"github.com/justinas/nosurf"
)

type BaseHandler struct {
	DB         *db.Queries
	Renderer   *render.Engine
	Cfg        *config.Config
	Translator *translator.Translator
}

func (h *BaseHandler) RenderPage(w http.ResponseWriter, r *http.Request, templateName string, pageVM any) {
	user, _ := GetUserFromContext(r)
	menu, _ := GetMenuFromContext(r)
	lang := GetLangFromCtx(r)

	base := viewmodel.NewBaseData(
		nosurf.Token(r),
		viewmodel.NewUserView(user),
		viewmodel.NewMenuItemView(menu),
		lang.String(),
		r.URL.Path,
	)

	data := viewmodel.PageContainer{
		Base: base,
		Page: pageVM,
	}

	h.Renderer.Render(w, templateName, data)
}

func (h *BaseHandler) T(r *http.Request, key string) string {
	lang := GetLangFromCtx(r)
	return h.Translator.T(r.Context(), key, lang)
}
