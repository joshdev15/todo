package todo

import (
	"encoding/json"
	"todo/internal/check"

	"github.com/google/uuid"
)

func New(title, text string) *Todo {
	return &Todo{
		Id:    uuid.NewString(),
		Title: title,
		Text:  text,
		Done:  false,
	}
}

type Todo struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Done  bool   `json:"done,omitempty"`
}

func (t *Todo) ToJSON() string {
	bytes, err := json.Marshal(t)
	check.Err(err)
	return string(bytes)
}
