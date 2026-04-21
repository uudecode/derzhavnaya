package render

import (
	"Derzhavnaya/internal/web/viewmodel"
	"Derzhavnaya/web"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Engine struct {
	templates map[string]*template.Template
}

func NewEngine() *Engine {
	e := &Engine{
		templates: make(map[string]*template.Template),
	}
	funcMap := GetFuncMap()
	entries, err := fs.ReadDir(web.Templates, "templates")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read templates directory")
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || name == "base.html" || name == "partials.html" {
			continue
		}

		t, err := template.New("base").Funcs(funcMap).ParseFS(web.Templates,
			"templates/base.html",
			"templates/partials.html",
			"templates/"+name)

		if err != nil {
			log.Fatal().Err(err).Msgf("failed to parse template: %s", name)
		}
		e.templates[name] = t
		log.Debug().Msgf("Template %s cached", name)
	}
	return e
}

func (e *Engine) Render(w http.ResponseWriter, name string, data viewmodel.PageContainer) {

	t, ok := e.templates[name]
	if !ok {
		log.Error().Msgf("template %s not found in cache", name)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Error().Err(err).Msgf("failed to execute template %s", name)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
