package viewmodel

import (
	"html/template"
	"time"
)

type QuestionView struct {
	ID       int32
	DataQ    time.Time
	Name     string
	Question template.HTML
	Answer   template.HTML
}

type TalkView struct {
	Talks          []QuestionView
	HasNext        bool
	NextCursorTime string
	NextCursorID   int32
	CurrentPath    string
}

func NewTalkView(talks []QuestionView, hasNext bool, nextCursorTime string, nextCursorID int32, currentPath string) TalkView {
	return TalkView{
		Talks:          talks,
		HasNext:        hasNext,
		NextCursorTime: nextCursorTime,
		NextCursorID:   nextCursorID,
		CurrentPath:    currentPath,
	}
}
