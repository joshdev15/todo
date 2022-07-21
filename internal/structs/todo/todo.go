package todo

import (
	"encoding/json"
	"strings"
	"todo/internal/check"

	"github.com/google/uuid"
)

// New creates a new Todo
// example: newTodo := todo.New("My todo", "This is my todo")
func New(title, message string) *Todo {
	tmpNewId := uuid.NewString()
	id := strings.ReplaceAll(tmpNewId, "-", "")

	return &Todo{
		Id:      id,
		Title:   Text(title),
		Message: Text(message),
		Done:    false,
	}
}

// Todo model
type Todo struct {
	Id      string `json:"id,omitempty"`
	Title   Text   `json:"title,omitempty"`
	Message Text   `json:"message,omitempty"`
	Done    bool   `json:"done,omitempty"`
}

// Convert the current to-do to json
func (t *Todo) ToJSON() string {
	bytes, err := json.Marshal(t)
	check.Err(err)
	return string(bytes)
}

func JSONToTODO(value string) (Todo, error) {
	var todoElement Todo

	err := json.Unmarshal([]byte(value), &todoElement)
	if err != nil {
		return Todo{}, err
	}

	return todoElement, nil
}
