package render

import (
	"html/template"
	"strings"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"hasPrefix": strings.HasPrefix,
		"contains":  strings.Contains,
	}
}
