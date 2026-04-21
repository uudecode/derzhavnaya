-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS web.menu_item_translation (
                                                                category_id INTEGER NOT NULL REFERENCES web.gallery_category(id) ON DELETE CASCADE,
                                                                locale TEXT NOT NULL CHECK (locale IN ('ru', 'en', 'fr')),
                                                                title TEXT NOT NULL,
                                                                slug TEXT NOT NULL,
                                                                description TEXT NOT NULL DEFAULT '',

                                                                PRIMARY KEY (category_id, locale),
                                                                UNIQUE (locale, slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd