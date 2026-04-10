-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS web;

CREATE TABLE web.users (
                           id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           email      TEXT NOT NULL UNIQUE,
                           password   TEXT NOT NULL, -- хеш пароля
                           full_name  TEXT,
                           role       TEXT NOT NULL DEFAULT 'guest', -- упростим: admin, editor, guest
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE web.sessions (
                              id         TEXT PRIMARY KEY, -- случайный токен для куки
                              user_id    UUID NOT NULL REFERENCES web.users(id) ON DELETE CASCADE,
                              expires_at TIMESTAMPTZ NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Создадим индекс для быстрого поиска сессий
CREATE INDEX idx_sessions_user_id ON web.sessions(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS web.sessions;
DROP TABLE IF EXISTS web.users;
DROP SCHEMA IF EXISTS web;
-- +goose StatementEnd