package viewmodel

import "Derzhavnaya/internal/models"

type MenuItemView struct {
	Label string
	Url   string
	Icon  string
}

func NewMenuItemView(items []models.MenuItem) []MenuItemView {
	views := make([]MenuItemView, 0, len(items))

	for _, item := range items {
		views = append(views, MenuItemView{
			Label: item.Label,
			Url:   item.Url,
			Icon:  item.Icon,
		})
	}
	return views
}
