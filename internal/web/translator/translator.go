package translator

import (
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/types"
	"context"
	"sync"
)

type Translator struct {
	DB    *db.Queries
	cache sync.Map
}

func NewTranslator(queries *db.Queries) *Translator {
	return &Translator{
		DB:    queries,
		cache: sync.Map{},
	}
}
func (t *Translator) T(ctx context.Context, key string, lang types.Locale) string {
	cacheKey := lang.String() + ":" + key
	if val, ok := t.cache.Load(cacheKey); ok {
		return val.(string)
	}

	val, err := t.DB.GetTranslation(ctx, db.GetTranslationParams{
		Key:  key,
		Lang: lang.String(),
	})
	if err != nil {
		return "MISSING_" + key
	}

	t.cache.Store(cacheKey, val.Value)
	return val.Value
}
