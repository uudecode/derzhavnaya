-- +goose Up
-- +goose StatementBegin
CREATE TABLE web.translation (
                                 id serial PRIMARY KEY,
                                 key text NOT NULL,
                                 lang text NOt NULL,
                                 value text NOT NULL
);

CREATE INDEX idx_translation_key_and_lang ON web.translation(key, lang);

INSERT INTO web.translation (key, lang, value) VALUES ('site.index_title', 'ru', 'Санкт-Петербургская епархия');
INSERT INTO web.translation (key, lang, value) VALUES ('site.index_title', 'en', 'Saint Petersburg Diocese of the Russian Orthodox Church');
INSERT INTO web.translation (key, lang, value) VALUES ('site.index_title', 'fr', 'Diocèse de Saint-Pétersbourg');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS web.translation;
-- +goose StatementEnd