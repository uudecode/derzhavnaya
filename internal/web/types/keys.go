package types

type contextKey string

const (
	MenuKey contextKey = "menu_items"
	PathKey contextKey = "current_path"
	UserKey contextKey = "user_info"
)
