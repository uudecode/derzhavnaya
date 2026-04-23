package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Translation struct {
	EN string `json:"en"`
	FR string `json:"fr"`
}

type Glossary map[string]map[string]Translation

func RunGlossaryImport(ctx context.Context, pool *pgxpool.Pool, fileName string) error {
	log.Info().Msg("Starting glossary import...")

	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read glossary file: %w", err)
	}

	var g Glossary
	if err := json.Unmarshal(data, &g); err != nil {
		return fmt.Errorf("unmarshal glossary JSON: %w", err)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	stmt := `
		INSERT INTO web.glossary (category, ru_term, en_trans, fr_trans)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (category, ru_term) DO UPDATE
		SET en_trans = EXCLUDED.en_trans,
		    fr_trans = EXCLUDED.fr_trans
	`

	for cat, terms := range g {
		for ru, trans := range terms {
			_, err := tx.Exec(ctx, stmt, cat, ru, trans.EN, trans.FR)
			if err != nil {
				return fmt.Errorf("insert term [%s/%s]: %w", cat, ru, err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	log.Info().Msg("Glossary import completed successfully!")
	return nil
}
