package translation

import (
	"Derzhavnaya/internal/db"
	"context"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func StartTranslationWorker(lc fx.Lifecycle, queries *db.Queries, ts Translator) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				testText := "Здравствуйте отец Дмитрий. У меня такая беда. Болеет ребенок 13 лет. У нее боли в грудной клетке.Каждую неделю ходим к врачам они только руками разводят. Все анализы, узи, обследование в норме. Думали сердце врачи говорят что нет. Говорят, что это нервное. Выписали нам психотропные таблетки которые я не решаюсь давать ребенку. Каждый день читаю молитвы о болящем. Причащаемся раз в месяц по благословению батюшки. Заказываю записочки о здравии и сорокоуст в разных храмах. Я наверное виновата перед Господом. Все советую отвести ребенка к бабке. Подскажите пожалуйста что мне делать."
				langs := []string{"en", "fr"}

				for _, lang := range langs {
					res, err := ts.Translate(context.Background(), testText, lang)
					if err != nil {
						log.Error().Err(err).Str("lang", lang).Msg("Тестовый перевод не удался")
						continue
					}
					log.Info().Str("lang", lang).Str("result", res).Msg("Успешный тестовый перевод!")
				}
			}()
			//go func() {
			//	ticker := time.NewTicker(1 * time.Minute) // Можно настроить интервал
			//	for range ticker.C {
			//		// 1. Берем задачу из БД через SQLC
			//		// 2. ts.Translate(...)
			//		// 3. Сохраняем результат
			//	}
			//}()
			return nil
		},
	})
}
