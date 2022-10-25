package todo

import (
	"fmt"
	"testing"
)

func TestTodo(t *testing.T) {
	todo := New("My todo", "This is my todo")
	todoID := todo.Id
	todoJSON := fmt.Sprintf("{\"id\":\"%s\",\"title\":\"My todo\",\"message\":\"This is my todo\",\"done\":false}", todoID)
	errTodoJSON := fmt.Sprintf("\"iden\":\"%s\",\"titulo\":\"My todo\",\"msg\":\"This is my todo\",\"dona\":false", todoID)
	fromJSONTodo, err := JSONToTODO(todoJSON)
	_, badJSON := JSONToTODO(errTodoJSON)

	if err != nil {
		t.Error("Test JSON invalid")
	}

	if badJSON == nil {
		t.Error("Bad JSON isn´t nil")
	}

	if todo.Id == "" && len(todo.Id) != 32 {
		t.Error("Id should be a 32 character string")
	}

	if todo.Title.ToString() != "My todo" {
		t.Error("Title should be 'My todo'")
	}

	if todo.Message.ToString() != "This is my todo" {
		t.Error("Message should be 'This is my todo'")
	}

	if todo.Done != false {
		t.Error("Done should be false")
	}

	if todo.ToJSON() != todoJSON {
		t.Error("todo in JSON should be equal to test JSON")
	}

	if fromJSONTodo.Id != todo.Id && fromJSONTodo.Title.ToString() != todo.Title.ToString() {
		t.Error("todo in JSON should be equal to test JSON")
	}
}
