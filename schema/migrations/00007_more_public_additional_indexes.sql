-- +goose Up
-- +goose StatementBegin

CREATE INDEX IF NOT EXISTS idx_hram_talk_pagination2 ON public.hram_talk(flag, data_a DESC, id DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX public.idx_hram_talk_pagination2;

-- +goose StatementEnd