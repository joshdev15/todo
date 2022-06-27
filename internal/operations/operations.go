package operations

import (
	"fmt"
	"todo/internal/memory"
	"todo/internal/structs/todo"
)

func ListTodos() {
	memory.GetList()
}

func AddTodo(title, text string) {
	newTodo := todo.New(title, text)
	fmt.Println(newTodo.ToJSON())
	memory.Set(newTodo.Id, newTodo.ToJSON())
}

func ShowTodo(id string) {
	memory.Get(id)
}

func MarkAsDone(id string) {
	memory.SetDone(id)
}

func ClearTodoList() {
	memory.DeleteAllTodos()
}

func RemoveTodo(id string) {
	memory.Remove(id)
}
