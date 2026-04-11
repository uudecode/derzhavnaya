-- +goose Up
-- +goose StatementBegin

CREATE TABLE web.menu_item (
                           id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           position      smallint NOT NULL UNIQUE,
                           label      TEXT NOT NULL UNIQUE,
                           icon   TEXT NOT NULL, 
                           url   TEXT NOT NULL, 
                           is_active BOOLEAN NOT NULL DEFAULT true,
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO web.menu_item (position, icon, label, url) VALUES (1,'🏠','Главная','/');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (2,'📰', 'Новости','/news');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (3,'📸', 'Фотогалерея','/gallery');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (4,'📖', 'Библиотека','/lib');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (5,'🕯', 'Расписание','/schedule');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS web.menu_item;
-- +goose StatementEnd