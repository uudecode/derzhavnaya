-- +goose Up
-- +goose StatementBegin
CREATE TABLE web.glossary (
                          id SERIAL PRIMARY KEY,
                          category TEXT NOT NULL,
                          ru_term TEXT NOT NULL,
                          en_trans TEXT,
                          fr_trans TEXT,
                          CONSTRAINT glossary_category_ru_term_uk UNIQUE (category, ru_term)
);

CREATE INDEX glossary_category_idx ON web.glossary (category);
CREATE INDEX glossary_term_idx ON web.glossary (ru_term);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE web.glossary;
-- +goose StatementEnd