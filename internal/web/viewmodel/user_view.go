package viewmodel

import "Derzhavnaya/internal/models"

type UserView struct {
	FullName string
	Email    string
	Role     string
}

func NewUserView(u models.User) *UserView {
	if u.ID == "" {
		return nil
	}
	return &UserView{
		FullName: u.FullName,
		Email:    u.Email,
		Role:     u.Role,
	}
}
