// The flags package provides the functionality to be performed by the program
// (read, parse and execute) the arguments and flags that the user enters in
// the command line.
package flags

import (
	"flag"
	"os"
	"todo/internal/operations"
	"todo/internal/printer"
)

var (
	// action is a private package variable that stores at runtime the action to be
	// taken according to the argument used by the user in the program execution.
	action = ""
)

// Function that reads the second argument is os.Args which defines the action
// to be taken by the program.
func ReadFlags() {
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "list":
		listFlag()
	case "add":
		addFlag()
	case "show":
		showFlag()
	case "done":
		doneFlag()
	case "remove":
		removeFlag()
	case "clean":
		cleanFlag()
	case "help":
		helpFlag()
	default:
		helpFlag()
	}
}

// function that reads the "list" flags and executes the operation related to
// the action
func listFlag() {
	listFlag := flag.NewFlagSet("list", flag.ContinueOnError)

	var title string
	listFlag.StringVar(&title, "title", "", "todo title")

	listFlag.Parse(os.Args[2:])

	operations.ListTodos()
}

// function that reads the "add" flags and executes the operation related to
// the action
func addFlag() {
	addFlag := flag.NewFlagSet("add", flag.ContinueOnError)

	var title string
	addFlag.StringVar(&title, "title", "", "todo title")

	var text string
	addFlag.StringVar(&text, "text", "", "todo text")

	addFlag.Parse(os.Args[2:])

	operations.AddTodo(title, text)
}

// function that reads the "show" flags and executes the operation related to
// the action
func showFlag() {
	if len(os.Args) > 2 {
		operations.ShowTodo(os.Args[2])
	}
}

// function that reads the third argument of os.Args and executes the operation
// related to the action "remove".
func removeFlag() {
	if len(os.Args) > 2 {
		operations.RemoveTodo(os.Args[2])
	}
}

// function that reads the third argument of os.Args and executes the operation
// related to the action "done".
func doneFlag() {
	if len(os.Args) > 2 {
		operations.MarkAsDone(os.Args[2])
	}
}

// function that executes the operation related to the action "clean"
func cleanFlag() {
	operations.ClearTodoList()
}

// function that prints help on the terminal
func helpFlag() {
	printer.Help()
}
