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
-- +goose StatementEnd