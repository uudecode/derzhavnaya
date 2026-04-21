package translator

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/types"
	"context"
)

type Translator struct {
	DB *db.Queries
}

func NewTranslator(queries *db.Queries) *Translator {
	return &Translator{
		DB: queries,
	}
}
func (t *Translator) T(ctx context.Context, key string, lang types.Locale) string {
	translation, err := t.DB.GetTranslation(ctx, db.GetTranslationParams{
		Key:  key,
		Lang: lang.String(),
	})
	if err != nil {
		return "MISSING_" + key
	}
	return translation.Value
}
