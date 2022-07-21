// Operations package that executes the action related to the operations
// required by the user.
package operations

import (
	"todo/internal/check"
	"todo/internal/memory"
	"todo/internal/printer"
	"todo/internal/structs/todo"
)

// Execution of listing the to-do list
func ListTodos() {
	list, err := memory.GetAll()
	check.Err(err)
	todoList := []todo.Todo{}

	for _, v := range list {
		currentTodo, err := todo.JSONToTODO(v)
		check.Err(err)
		todoList = append(todoList, currentTodo)
	}

	printer.TodoTable(todoList)
}

// Execution of addition of a new to-do
func AddTodo(title, message string) {
	newTodo := todo.New(title, message)
	err := memory.Set(newTodo.Id, newTodo.ToJSON())
	check.Err(err)
	printer.Success("Add to-do")
}

// Execution of displaying the information contained in a task
func Show(key string) {
	result, err := memory.Get(key)
	check.Err(err)

	currentTodo, err := todo.JSONToTODO(result)
	check.Err(err)

	printer.Show(currentTodo)
}

// Execution of eliminating a to-do
func Remove(key string) {
	memory.Remove(key)
}

// Execution of cleaning the to-do list
func RemoveAll() {
	memory.DeleteAllTodos()
}

// Execution of marking a to-do as done
func MarkAsDone(key string) {
	result, err := memory.Get(key)
	check.Err(err)

	currentTodo, err := todo.JSONToTODO(result)
	check.Err(err)

	currentTodo.Done = !currentTodo.Done
	modifiedTodo := currentTodo.ToJSON()

	err = memory.Set(key, modifiedTodo)
	check.Err(err)
}
