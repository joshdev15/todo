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
	action        = ""
	actionOptions = map[string]func(){
		"list":    listFlag,
		"add":     addFlag,
		"show":    showFlag,
		"done":    doneFlag,
		"remove":  removeFlag,
		"help":    helpFlag,
		"version": versionFlag,
	}
)

// Function that reads the second argument is os.Args which defines the action
// to be taken by the program.
func ReadFlags() {
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	var existAction bool = false
	var tmpAction func()
	for k := range actionOptions {
		if action == k {
			existAction = true
			tmpAction = actionOptions[action]
			break
		}
	}

	if !existAction {
		helpFlag()
		return
	}

	tmpAction()
}

// function that reads the "list" flags and executes the operation related to
// the action
func listFlag() {
	listFlag := flag.NewFlagSet("list", flag.ContinueOnError)

	var title string
	listFlag.StringVar(&title, "title", "", "to-do title")

	listFlag.Parse(os.Args[2:])

	operations.ListTodos()
}

// function that reads the "add" flags and executes the operation related to
// the action
func addFlag() {
	addFlag := flag.NewFlagSet("add", flag.ContinueOnError)

	var title string
	addFlag.StringVar(&title, "t", "", "to-do title")

	var message string
	addFlag.StringVar(&message, "m", "", "to-do message")

	addFlag.Parse(os.Args[2:])

	operations.AddTodo(title, message)
}

// function that reads the "show" flags and executes the operation related to
// the action
func showFlag() {
	if len(os.Args) > 2 {
		operations.Show(os.Args[2])
	}
}

// function that reads the third argument of os.Args and executes the operation
// related to the action "remove".
func removeFlag() {
	removeFlag := flag.NewFlagSet("remove", flag.ContinueOnError)

	var all bool
	removeFlag.BoolVar(&all, "a", false, "remove all to-do list")
	removeFlag.Parse(os.Args[2:])

	switch all {
	case true:
		modalResponse := printer.Modal("Remove all todo list")
		if modalResponse {
			operations.RemoveAll()
		}

		return

	case false:
		if len(os.Args) >= 2 {
			operations.Remove(os.Args[2])
			return
		}
	}

	printer.NoArgs()
}

// function that reads the third argument of os.Args and executes the operation
// related to the action "done".
func doneFlag() {
	if len(os.Args) > 2 {
		operations.MarkAsDone(os.Args[2])
	}
}

// function that prints help on the terminal
func helpFlag() {
	printer.Help()
}

func versionFlag() {
	printer.Version()
}
