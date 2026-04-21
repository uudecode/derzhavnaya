package render

import (
	"Derzhavnaya/internal/web/translator"
	"Derzhavnaya/internal/web/types"
	"Derzhavnaya/internal/web/viewmodel"
	"context"
	"html/template"
	"strings"
)

func GetFuncMap(trans *translator.Translator) template.FuncMap {
	return template.FuncMap{
		"hasPrefix": strings.HasPrefix,
		"contains":  strings.Contains,
		"T": func(data viewmodel.PageContainer, key string) string {
			lang := types.Locale(data.Base.CurrentLang)
			return trans.T(context.Background(), key, lang)
		},
	}
}
