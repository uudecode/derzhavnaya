-- +goose Up
-- +goose StatementBegin
-- Меню (keys)
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.home', 'ru', 'Главная');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.home', 'en', 'Home');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.home', 'fr', 'Accueil');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.news', 'ru', 'Новости');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.news', 'en', 'News');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.news', 'fr', 'Actualités');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.gallery', 'ru', 'Фотогалерея');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.gallery', 'en', 'Gallery');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.gallery', 'fr', 'Galerie');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lib', 'ru', 'Библиотека');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lib', 'en', 'Library');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lib', 'fr', 'Bibliothèque');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.schedule', 'ru', 'Расписание');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.schedule', 'en', 'Schedule');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.schedule', 'fr', 'Horaires');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_news', 'ru', 'Редактор новостей');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_news', 'en', 'News Editor');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_news', 'fr', 'Éditeur de nouvelles');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_talks', 'ru', 'Обработка вопросов');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_talks', 'en', 'Questions Editor');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.admin_talks', 'fr', 'Gestion des questions');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.talks', 'ru', 'Разговор со священником');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.talks', 'en', 'Talk to a priest');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.talks', 'fr', 'Parler à un prêtre');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.music', 'ru', 'Нотный стан');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.music', 'en', 'Sheet music');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.music', 'fr', 'Partitions');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.prayers', 'ru', 'Просим ваших молитв');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.prayers', 'en', 'Prayer requests');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.prayers', 'fr', 'Demandes de prières');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.relics', 'ru', 'Наши святыни');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.relics', 'en', 'Our relics');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.relics', 'fr', 'Nos reliques');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.calendar', 'ru', 'Православный календарь');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.calendar', 'en', 'Orthodox calendar');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.calendar', 'fr', 'Calendrier orthodoxe');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.ruined', 'ru', 'Разрушенные храмы');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.ruined', 'en', 'Ruined churches');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.ruined', 'fr', 'Églises en ruines');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.thoughts', 'ru', 'Мудрые мысли');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.thoughts', 'en', 'Wise thoughts');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.thoughts', 'fr', 'Pensées sages');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.history', 'ru', 'История храма');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.history', 'en', 'History of the church');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.history', 'fr', 'Histoire de l''église');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.helios', 'ru', 'Центр Гелиос');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.helios', 'en', 'Helios Center');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.helios', 'fr', 'Centre Helios');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.contacts', 'ru', 'Как нас найти');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.contacts', 'en', 'How to find us');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.contacts', 'fr', 'Comment nous trouver');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lawyer', 'ru', 'Страница юриста');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lawyer', 'en', 'Lawyer page');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.lawyer', 'fr', 'Page juridique');

INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.pilgrims', 'ru', 'Паломникам');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.pilgrims', 'en', 'For pilgrims');
INSERT INTO web.translation (key, lang, value) VALUES ('menu.item.pilgrims', 'fr', 'Pour les pèlerins');


UPDATE web.menu_item SET label = 'menu.item.home' WHERE url = '/';
UPDATE web.menu_item SET label = 'menu.item.news' WHERE url = '/news';
UPDATE web.menu_item SET label = 'menu.item.gallery' WHERE url = '/gallery';
UPDATE web.menu_item SET label = 'menu.item.lib' WHERE url = '/lib';
UPDATE web.menu_item SET label = 'menu.item.schedule' WHERE url = '/schedule';
UPDATE web.menu_item SET label = 'menu.item.admin_news' WHERE url = '/admin/news';
UPDATE web.menu_item SET label = 'menu.item.admin_talks' WHERE url = '/admin/talks';
UPDATE web.menu_item SET label = 'menu.item.talks' WHERE url = '/talks';
UPDATE web.menu_item SET label = 'menu.item.music' WHERE url = '/music';
UPDATE web.menu_item SET label = 'menu.item.prayers' WHERE url = '/prayers';
UPDATE web.menu_item SET label = 'menu.item.relics' WHERE url = '/relics';
UPDATE web.menu_item SET label = 'menu.item.calendar' WHERE url = '/calendar';
UPDATE web.menu_item SET label = 'menu.item.ruined' WHERE url = '/ruined';
UPDATE web.menu_item SET label = 'menu.item.thoughts' WHERE url = '/thoughts';
UPDATE web.menu_item SET label = 'menu.item.history' WHERE url = '/history';
UPDATE web.menu_item SET label = 'menu.item.helios' WHERE url = '/helios';
UPDATE web.menu_item SET label = 'menu.item.contacts' WHERE url = '/contacts';
UPDATE web.menu_item SET label = 'menu.item.lawyer' WHERE url = '/lawyer';
UPDATE web.menu_item SET label = 'menu.item.pilgrims' WHERE url = '/pilgrims';


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM web.translation WHERE key LIKE 'menu.item.%';
UPDATE web.menu_item SET label = 'Главная' WHERE url = '/';
UPDATE web.menu_item SET label = 'Новости' WHERE url = '/news';
UPDATE web.menu_item SET label = 'Фотогалерея' WHERE url = '/gallery';
UPDATE web.menu_item SET label = 'Библиотека' WHERE url = '/lib';
UPDATE web.menu_item SET label = 'Расписание' WHERE url = '/schedule';
UPDATE web.menu_item SET label = 'Редактор новостей' WHERE url = '/admin/news';
UPDATE web.menu_item SET label = 'Обработка вопросов' WHERE url = '/admin/talks';
UPDATE web.menu_item SET label = 'Разговор со священником' WHERE url = '/talks';
UPDATE web.menu_item SET label = 'Нотный стан' WHERE url = '/music';
UPDATE web.menu_item SET label = 'Просим ваших молитв' WHERE url = '/prayers';
UPDATE web.menu_item SET label = 'Наши святыни' WHERE url = '/relics';
UPDATE web.menu_item SET label = 'Православный календарь' WHERE url = '/calendar';
UPDATE web.menu_item SET label = 'Разрушенные храмы' WHERE url = '/ruined';
UPDATE web.menu_item SET label = 'Мудрые мысли' WHERE url = '/thoughts';
UPDATE web.menu_item SET label = 'История храма' WHERE url = '/history';
UPDATE web.menu_item SET label = 'Центр Гелиос' WHERE url = '/helios';
UPDATE web.menu_item SET label = 'Как нас найти' WHERE url = '/contacts';
UPDATE web.menu_item SET label = 'Страница юриста' WHERE url = '/lawyer';
UPDATE web.menu_item SET label = 'Паломникам' WHERE url = '/pilgrims';
-- +goose StatementEnd