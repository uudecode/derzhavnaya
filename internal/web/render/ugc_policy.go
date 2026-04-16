package render

import "github.com/microcosm-cc/bluemonday"

var UgcPolicy *bluemonday.Policy

func init() {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowAttrs("href").OnElements("a")
	p.AllowElements("p", "b", "strong", "i", "em", "br")
	UgcPolicy = p
}
