package models

import (
	"Derzhavnaya/internal/db"
	"time"
)

type MenuItem struct {
	ID        string
	Position  int16
	Label     string
	Icon      string
	Url       string
	IsActive  bool
	CreatedAt time.Time
	Role      string
}

func NewMenuItemFromDB(dbMenuItem db.WebMenuItem) MenuItem {
	return MenuItem{
		ID:        dbMenuItem.ID.String(),
		Position:  dbMenuItem.Position,
		Label:     dbMenuItem.Label,
		Icon:      dbMenuItem.Icon,
		Url:       dbMenuItem.Url,
		IsActive:  dbMenuItem.IsActive,
		CreatedAt: dbMenuItem.CreatedAt,
		Role:      ptrString(dbMenuItem.Role, ""),
	}
}
