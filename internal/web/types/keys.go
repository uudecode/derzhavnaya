package types

type contextKey string

const (
	MenuKey contextKey = "menu_items"
	PathKey contextKey = "current_path"
	LangKey contextKey = "lang"
	UserKey contextKey = "user"
)

type Locale string

const (
	EN Locale = "en"
	RU Locale = "ru"
	FR Locale = "fr"
)

func (l Locale) String() string {
	return string(l)
}
