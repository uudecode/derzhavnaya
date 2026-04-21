package models

import "Derzhavnaya/internal/db"

type User struct {
	ID       string
	FullName string
	Email    string
	Role     string
}

func NewUserFromDB(dbUser db.WebUser) User {
	return User{
		ID:       dbUser.ID.String(),
		FullName: ptrString(dbUser.FullName, "Incognito"),
		Role:     dbUser.Role,
		Email:    dbUser.Email,
	}
}
