-- +goose Up
-- +goose StatementBegin

INSERT INTO web.translation (key, lang, value) VALUES ('header.title', 'ru', 'Санкт-Петербургская епархия');
INSERT INTO web.translation (key, lang, value) VALUES ('header.title', 'en', 'Saint Petersburg Diocese of the Russian Orthodox Church');
INSERT INTO web.translation (key, lang, value) VALUES ('header.title', 'fr', 'Diocèse de Saint-Pétersbourg');

INSERT INTO web.translation (key, lang, value) VALUES ('header.login', 'ru', 'Войти');
INSERT INTO web.translation (key, lang, value) VALUES ('header.login', 'en', 'Login');
INSERT INTO web.translation (key, lang, value) VALUES ('header.login', 'fr', 'Se connecter');

INSERT INTO web.translation (key, lang, value) VALUES ('header.logout', 'ru', 'Выйти');
INSERT INTO web.translation (key, lang, value) VALUES ('header.logout', 'en', 'Logout');
INSERT INTO web.translation (key, lang, value) VALUES ('header.logout', 'fr', 'Se déconnecter');


INSERT INTO web.translation (key, lang, value) VALUES ('auth.login', 'ru', 'Войти');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.login', 'en', 'Login');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.login', 'fr', 'Se connecter');

INSERT INTO web.translation (key, lang, value) VALUES ('auth.login_title', 'ru', 'Вход в систему');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.login_title', 'en', 'Login to system');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.login_title', 'fr', 'Connexion');

INSERT INTO web.translation (key, lang, value) VALUES ('auth.email_placeholder', 'ru', 'Email адрес');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.email_placeholder', 'en', 'Email address');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.email_placeholder', 'fr', 'Adresse email');

INSERT INTO web.translation (key, lang, value) VALUES ('auth.password_placeholder', 'ru', 'Пароль');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.password_placeholder', 'en', 'Password');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.password_placeholder', 'fr', 'Mot de passe');

INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.too_many_attempts', 'ru', 'Слишком много попыток. Попробуйте позже.');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.too_many_attempts', 'en', 'Too many attempts. Please try again later.');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.too_many_attempts', 'fr', 'Trop de tentatives. Veuillez réessayer plus tard.');

INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.invalid_credentials', 'ru', 'Неверный логин или пароль');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.invalid_credentials', 'en', 'Invalid username or password');
INSERT INTO web.translation (key, lang, value) VALUES ('auth.error.invalid_credentials', 'fr', 'Nom d''utilisateur ou mot de passe incorrect');

INSERT INTO web.translation (key, lang, value) VALUES ('error.server_error', 'ru', 'Ошибка сервера');
INSERT INTO web.translation (key, lang, value) VALUES ('error.server_error', 'en', 'Server error');
INSERT INTO web.translation (key, lang, value) VALUES ('error.server_error', 'fr', 'Erreur de serveur');

INSERT INTO web.translation (key, lang, value) VALUES ('footer.copyright', 'ru', '© 2001—2026 Храм иконы Божией Матери «Державная»');
INSERT INTO web.translation (key, lang, value) VALUES ('footer.copyright', 'en', '© 2001—2026 Church of the Derzhavnaya Icon of the Mother of God');
INSERT INTO web.translation (key, lang, value) VALUES ('footer.copyright', 'fr', '© 2001—2026 Église de l''icône de la Mère de Dieu «Derzhavnaya»');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.load_more', 'ru', 'Показать еще вопросы');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.load_more', 'en', 'Show more questions');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.load_more', 'fr', 'Afficher plus de questions');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.no_questions', 'ru', 'Вопросов пока нет...');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.no_questions', 'en', 'No questions yet...');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.no_questions', 'fr', 'Aucune question pour l''instant...');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.title', 'ru', 'Вопросы и ответы');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.title', 'en', 'Questions and Answers');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.title', 'fr', 'Questions et réponses');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.description', 'ru', 'Духовные наставления и ответы на вопросы прихожан');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.description', 'en', 'Spiritual guidance and answers to parishioners’ questions');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.description', 'fr', 'Conseils spirituels et réponses aux questions des paroissiens');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.from', 'ru', 'От кого:');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.from', 'en', 'From:');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.from', 'fr', 'De :');

INSERT INTO web.translation (key, lang, value) VALUES ('talks.priest_answer', 'ru', 'Ответ священника:');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.priest_answer', 'en', 'Priest’s answer:');
INSERT INTO web.translation (key, lang, value) VALUES ('talks.priest_answer', 'fr', 'Réponse du prêtre :');

-- Храм во имя иконы Божией Матери (Верхняя часть заголовка)
INSERT INTO web.translation (key, lang, value) VALUES ('index.church_name', 'ru', 'Храм во имя иконы Божией Матери');
INSERT INTO web.translation (key, lang, value) VALUES ('index.church_name', 'en', 'Church of the icon of the Mother of God');
INSERT INTO web.translation (key, lang, value) VALUES ('index.church_name', 'fr', 'Église de l''icône de la Mère de Dieu');

-- «ДЕРЖАВНАЯ» (Основное название иконы)
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_name', 'ru', '«ДЕРЖАВНАЯ»');
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_name', 'en', '«DERZHAVNAYA»');
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_name', 'fr', '«DERZHAVNAYA»');


-- «Добро пожаловать в наш приход»
INSERT INTO web.translation (key, lang, value) VALUES ('index.welcome', 'ru', 'Добро пожаловать в наш приход');
INSERT INTO web.translation (key, lang, value) VALUES ('index.welcome', 'en', 'Welcome to our parish');
INSERT INTO web.translation (key, lang, value) VALUES ('index.welcome', 'fr', 'Bienvenue dans notre paroisse');

-- Икона «Державная»
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_title', 'ru', 'Икона «Державная»');
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_title', 'en', 'Derzhavnaya Icon');
INSERT INTO web.translation (key, lang, value) VALUES ('index.icon_title', 'fr', 'Icône «Derzhavnaya»');

-- Наш храм — это место молитвы и утешения...
INSERT INTO web.translation (key, lang, value) VALUES ('index.hero_text', 'ru', 'Наш храм — это место молитвы и утешения. Мы рады приветствовать вас на обновленном сайте.');
INSERT INTO web.translation (key, lang, value) VALUES ('index.hero_text', 'en', 'Our church is a place of prayer and comfort. We are glad to welcome you to our updated website.');
INSERT INTO web.translation (key, lang, value) VALUES ('index.hero_text', 'fr', 'Notre église est un lieu de prière et de réconfort. Nous sommes heureux de vous accueillir sur notre site mis à jour.');

-- События прихода
INSERT INTO web.translation (key, lang, value) VALUES ('index.events_title', 'ru', 'События прихода');
INSERT INTO web.translation (key, lang, value) VALUES ('index.events_title', 'en', 'Parish events');
INSERT INTO web.translation (key, lang, value) VALUES ('index.events_title', 'fr', 'Événements de la paroisse');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM web.translation WHERE key = 'header.title';
DELETE FROM web.translation WHERE key = 'header.login';
DELETE FROM web.translation WHERE key = 'header.logout';
DELETE FROM web.translation WHERE key = 'auth.login';
DELETE FROM web.translation WHERE key = 'auth.login_title';
DELETE FROM web.translation WHERE key = 'auth.email_placeholder';
DELETE FROM web.translation WHERE key = 'auth.password_placeholder';
DELETE FROM web.translation WHERE key = 'auth.error.too_many_attempts';
DELETE FROM web.translation WHERE key = 'auth.error.invalid_credentials';
DELETE FROM web.translation WHERE key = 'error.server_error';
DELETE FROM web.translation WHERE key = 'footer.copyright';
DELETE FROM web.translation WHERE key = 'talks.load_more';
DELETE FROM web.translation WHERE key = 'talks.no_questions';
DELETE FROM web.translation WHERE key = 'talks.title';
DELETE FROM web.translation WHERE key = 'talks.description';
DELETE FROM web.translation WHERE key = 'talks.from';
DELETE FROM web.translation WHERE key = 'talks.priest_answer';
DELETE FROM web.translation WHERE key LIKE 'index.%';
-- +goose StatementEnd