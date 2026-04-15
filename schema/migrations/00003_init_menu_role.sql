-- +goose Up
-- +goose StatementBegin

ALTER TABLE web.menu_item ADD COLUMN role text;

INSERT INTO web.menu_item (position, icon, label, url, role) VALUES (100,'✍️','Редактор новостей','/admin/news', 'admin');
INSERT INTO web.menu_item (position, icon, label, url, role) VALUES (101,'🕊️','Обработка вопросов','/admin/talks', 'admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE web.menu_item DROP COLUMN role;
-- +goose StatementEnd