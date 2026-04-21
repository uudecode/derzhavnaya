-- +goose Up
-- +goose StatementBegin

INSERT INTO web.gallery_category (id, code)
VALUES
    (-100, 'sovereign-iconography'),
    (-101, 'our-church'),
    (-102, 'church-shrines'),
    (-103, 'our-feast-days'),
    (-104, 'our-visits'),
    (-105, 'our-pilgrimages'),
    (-106, 'church-and-children'),
    (-107, 'christmas-educational-readings-moscow'),
    (-108, 'misc');


INSERT INTO web.gallery_category_translation
(category_id, locale, title, slug, description)
VALUES
    (-100, 'ru', 'Иконография Державного образа Божией Матери', 'ikonografiya-derzhavnogo-obraza-bozhiej-materi',
     '<p>На этой странице мы помещаем известную нам иконографию иконы Божией Матери "Державная". Кажется, это единственное в своем роде собрание различных списков Державной иконы, которое можно найти в Интернете.</p>
     <p>Часть изображений сделана нами, часть взята с различных сайтов или из иных источников. Качество изображений разное, поэтому заранее просим прощения.</p>
     <p>Мы будем рады, если Вы примете участие в составлении этой странички. Единственное "но": не принимаются изображения икон, написанных для так называемой "Православной Церкви Божией Матери Державная" - секта, более известная под названием "Богородичный центр" ("богородичники").</p>
     <p>О Державной иконе Божией Матери можно прочитать <a href="/svytany_derzhavnaya_history.html">здесь</a>.</p>'),

    (-100, 'en', 'Iconography of the Sovereign Icon of the Mother of God', 'sovereign-iconography',
     '<p>On this page, we present the iconography of the "Sovereign" Icon of the Mother of God known to us. This appears to be a unique collection of various copies of the Sovereign Icon available on the Internet.</p>
     <p>Some images were taken by us, others were sourced from various websites or other sources. The quality of the images varies, and we apologize for this in advance.</p>
     <p>We would be glad if you could participate in compiling this page. The only condition: we do not accept images of icons painted for the so-called "Orthodox Church of the Mother of God the Sovereign" — a sect better known as the "Bogorodichny Center".</p>
     <p>You can read about the Sovereign Icon of the Mother of God <a href="/svytany_derzhavnaya_history.html">here</a>.</p>'),

    (-100, 'fr', 'Iconographie de l’icône souveraine de la Mère de Dieu', 'iconographie-icone-souveraine',
     '<p>Sur cette page, nous présentons l’iconographie de l’icône "Souveraine" de la Mère de Dieu connue de nous. Il semble qu’il s’agisse de la seule collection unique de diverses copies de l’icône Souveraine disponible sur Internet.</p>
     <p>Certaines images ont été réalisées par nous, d’autres proviennent de divers sites ou d’autres sources. La qualité des images varie, et nous vous présentons nos excuses à l’avance.</p>
     <p>Nous serions ravis que vous participiez à la constitution de cette page. La seule condition : nous n’acceptons pas les images d’icônes peintes pour la soi-disant "Église orthodoxe de la Mère de Dieu Souveraine" — une secte plus connue sous le nom de "Centre Bogorodichny".</p>
     <p>Vous pouvez lire plus sur l’icône Souveraine de la Mère de Dieu <a href="/svytany_derzhavnaya_history.html">ici</a>.</p>'),

    (-101, 'ru', 'Наш храм', 'nash-khram', 'Воздвигаем иконостас'),
    (-101, 'en', 'Our Church', 'our-church', 'Erecting the iconostasis'),
    (-101, 'fr', 'Notre église', 'notre-eglise', 'Érection de l’iconostase'),

    (-102, 'ru', 'Святыни храма', 'svyatyni-khrama', 'Святыни храма'),
    (-102, 'en', 'Church Shrines', 'church-shrines', 'Church shrines'),
    (-102, 'fr', 'Sanctuaires de l’église', 'sanctuaires-eglise', 'Sanctuaires de l’église'),

    (-103, 'ru', 'Наши праздники', 'nashi-prazdniki', ''),
    (-103, 'en', 'Our Feast Days', 'our-feast-days', ''),
    (-103, 'fr', 'Nos fêtes', 'nos-fetes', ''),

    (-104, 'ru', 'Наши визиты', 'nashi-vizity', ''),
    (-104, 'en', 'Our Visits', 'our-visits', ''),
    (-104, 'fr', 'Nos visites', 'nos-visites', ''),

    (-105, 'ru', 'Наши паломничества', 'nashi-palomnichestva', ''),
    (-105, 'en', 'Our Pilgrimages', 'our-pilgrimages', ''),
    (-105, 'fr', 'Nos pèlerinages', 'nos-pelerinages', ''),

    (-106, 'ru', 'Храм и дети', 'khram-i-deti', ''),
    (-106, 'en', 'Church and Children', 'church-and-children', ''),
    (-106, 'fr', 'L’église et les enfants', 'eglise-et-enfants', ''),

    (-107, 'ru', 'Международные общеобразовательные Рождественские чтения в Москве', 'rozhdestvenskie-chteniya-v-moskve', ''),
    (-107, 'en', 'International Christmas Educational Readings in Moscow', 'christmas-educational-readings-moscow', ''),
    (-107, 'fr', 'Lectures éducatives de Noël internationales à Moscou', 'lectures-educatives-noel-moscou', ''),

    (-108, 'ru', 'Разное', 'raznoe', ''),
    (-108, 'en', 'Miscellaneous', 'misc', ''),
    (-108, 'fr', 'Divers', 'divers', '');



INSERT INTO web.gallery_category (id, parent_id, code)
VALUES
  (-10400, -104, 'leushino-2004-10-09');

INSERT INTO web.gallery_category_translation
(category_id,  locale, title, slug, description)
VALUES
    (-10400,'ru', 'Храмовый праздник ап. Иоанна Богослова на Леушинском подворье (9 октября 2004 г.)', 'leushino-2004-10-09', ''),
    (-10400,'en', 'Feast Day of St. John the Theologian at Leushino Metochion (October 9, 2004)', 'leushino-2004-10-09', ''),
    (-10400,'fr', 'Fête de l’apôtre Jean le Théologien au métochion de Leushino (9 octobre 2004)', 'leushino-2004-10-09', '')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

    DELETE FROM web.gallery_category_translation;
    DELETE FROM web.gallery_category;

-- +goose StatementEnd