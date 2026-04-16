-- +goose Up
-- +goose StatementBegin

CREATE INDEX IF NOT EXISTS idx_hram_talk_pagination ON public.hram_talk(flag, data_q DESC, id DESC);
CREATE INDEX IF NOT EXISTS idx_hram_news_date ON public.hram_news(data DESC);
CREATE INDEX IF NOT EXISTS idx_hram_lib_sub ON public.hram_lib(id_subsection);
CREATE INDEX IF NOT EXISTS idx_hram_mus_sub ON public.hram_mus(id_subsection);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name ON public."user" (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX public.idx_hram_talk_pagination;
DROP INDEX public.idx_hram_news_date;
DROP INDEX public.idx_hram_lib_sub;
DROP INDEX public.idx_hram_mus_sub;
DROP INDEX public.idx_user_name;

-- +goose StatementEnd