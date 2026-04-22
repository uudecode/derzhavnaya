package handlers

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"Derzhavnaya/internal/web/render"
	"Derzhavnaya/internal/web/translator"
	"Derzhavnaya/internal/web/viewmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type QuestionsHandler struct {
	BaseHandler
}

func NewQuestionsHandler(queries *db.Queries, cfg *config.Config, renderer *render.Engine, trans *translator.Translator) *QuestionsHandler {
	return &QuestionsHandler{
		BaseHandler: BaseHandler{
			DB:         queries,
			Renderer:   renderer,
			Cfg:        cfg,
			Translator: trans,
		},
	}
}

func (h *QuestionsHandler) Register(r chi.Router) {
	r.Get("/talks", h.Talks)
}

func (h *QuestionsHandler) Talks(w http.ResponseWriter, r *http.Request) {
	pageSize := 20
	cursorTimeStr := r.URL.Query().Get("cursor_time")
	cursorIDStr := r.URL.Query().Get("cursor_id")
	var items []db.HramTalk
	var err error

	limitWithNext := int32(pageSize + 1)
	if cursorTimeStr == "" || cursorIDStr == "" {
		items, err = h.DB.GetAnsweredQuestionsFirstPage(r.Context(), limitWithNext)
	} else {
		lastTime, _ := time.Parse(time.RFC3339, cursorTimeStr)
		lastID, _ := strconv.Atoi(cursorIDStr)
		items, err = h.DB.GetAnsweredQuestionsPaginated(r.Context(),
			db.GetAnsweredQuestionsPaginatedParams{
				LastDate: pgtype.Date{
					Time:  lastTime,
					Valid: true,
				},
				LastID:    int32(lastID),
				PageLimit: limitWithNext,
			})
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	hasNext := len(items) > pageSize

	var rawItems []db.HramTalk
	if hasNext {
		rawItems = items[:pageSize]
	} else {
		rawItems = items
	}

	viewItems := make([]viewmodel.QuestionView, len(rawItems))
	for i, item := range rawItems {
		cleanQ := render.UgcPolicy.Sanitize(item.Question)
		cleanA := render.UgcPolicy.Sanitize(item.Answer)

		viewItems[i] = viewmodel.QuestionView{
			ID:       item.ID,
			DataQ:    item.DataQ.Time,
			Name:     item.Name,
			Question: template.HTML(cleanQ),
			Answer:   template.HTML(cleanA),
		}
	}

	var nextCursorTime string
	var nextCursorID int32
	if hasNext {
		lastRaw := rawItems[len(rawItems)-1]
		nextCursorTime = lastRaw.DataQ.Time.Format(time.RFC3339)
		nextCursorID = lastRaw.ID
	}

	data := viewmodel.NewTalkView(viewItems, hasNext, nextCursorTime, nextCursorID, "/talks")

	h.RenderPage(w, r, "talks.html", data)
}
