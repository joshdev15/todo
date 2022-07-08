package todo

import (
	"encoding/json"
	"todo/internal/check"

	"github.com/google/uuid"
)

// New creates a new Todo
// example: newTodo := todo.New("My todo", "This is my todo")
func New(title, text string) *Todo {
	return &Todo{
		Id:    uuid.NewString(),
		Title: title,
		Text:  text,
		Done:  false,
	}
}

// Todo model
type Todo struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Done  bool   `json:"done,omitempty"`
}

// Convert the current to-do to json
func (t *Todo) ToJSON() string {
	bytes, err := json.Marshal(t)
	check.Err(err)
	return string(bytes)
}
