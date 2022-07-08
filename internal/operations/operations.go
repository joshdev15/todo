// Operations package that executes the action related to the operations
// required by the user.
package operations

import (
	"todo/internal/memory"
	"todo/internal/structs/todo"
)

// Execution of listing the to-do list
func ListTodos() {
	memory.GetList()
}

// Execution of addition of a new to-do
func AddTodo(title, text string) {
	newTodo := todo.New(title, text)
	memory.Set(newTodo.Id, newTodo.ToJSON())
}

// Execution of displaying the information contained in a task
func ShowTodo(id string) {
	memory.Get(id)
}

// Execution of marking a to-do as done
func MarkAsDone(id string) {
	memory.SetDone(id)
}

// Execution of cleaning the to-do list
func ClearTodoList() {
	memory.DeleteAllTodos()
}

// Execution of eliminating a to-do
func RemoveTodo(id string) {
	memory.Remove(id)
}
