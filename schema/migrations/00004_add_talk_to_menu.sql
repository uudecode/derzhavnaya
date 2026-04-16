-- +goose Up
-- +goose StatementBegin


INSERT INTO web.menu_item (position, icon, label, url) VALUES (6, '💬', 'Разговор со священником', '/talks');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (7, '🎶', 'Нотный стан', '/music');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (8, '🙏', 'Просим ваших молитв', '/prayers');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (9, '🏺', 'Наши святыни', '/relics');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (10, '📅', 'Православный календарь', '/calendar');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (11, '🏚️', 'Разрушенные храмы', '/ruined');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (12, '📜', 'Мудрые мысли', '/thoughts');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (13, '⏳', 'История храма', '/history');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (14, '☀️', 'Центр Гелиос', '/helios');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (15, '📍', 'Как нас найти', '/contacts');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (16, '⚖️', 'Страница юриста', '/lawyer');
INSERT INTO web.menu_item (position, icon, label, url) VALUES (17, '🎒', 'Паломникам', '/pilgrims');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM web.menu_item WHERE position IN (6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17);
-- +goose StatementEnd