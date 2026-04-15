package handlers

import (
	"Derzhavnaya/internal/db"
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestFilterMenu_RealData(t *testing.T) {
	h := &PageHandler{}

	// Данные, максимально приближенные к твоему CSV
	items := []db.WebMenuItem{
		{Label: "Главная", Role: pgtype.Text(sql.NullString{String: "", Valid: false})},                // Для всех
		{Label: "Новости", Role: pgtype.Text(sql.NullString{String: "", Valid: false})},                // Для всех
		{Label: "Редактор новостей", Role: pgtype.Text(sql.NullString{String: "admin", Valid: true})},  // Только админ
		{Label: "Обработка вопросов", Role: pgtype.Text(sql.NullString{String: "admin", Valid: true})}, // Только админ
	}

	tests := []struct {
		name     string
		userRole string
		expected int // Кол-во видимых пунктов
	}{
		{
			name:     "Anonymous user sees only public links",
			userRole: "",
			expected: 2, // Главная, Новости
		},
		{
			name:     "Admin user sees everything",
			userRole: "admin",
			expected: 4, // Все пункты
		},
		{
			name:     "Regular user sees only public links",
			userRole: "user",
			expected: 2, // Обычный юзер не админ, видит только публичное
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := h.filterMenu(items, tt.userRole)
			if len(result) != tt.expected {
				t.Errorf("Role '%s': expected %d items, got %d", tt.userRole, tt.expected, len(result))
			}
		})
	}
}
