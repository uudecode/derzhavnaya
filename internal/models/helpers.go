package models

import "github.com/jackc/pgx/v5/pgtype"

func ptrString(val pgtype.Text, fallback string) string {
	if val.Valid {
		return val.String
	}
	return fallback
}
