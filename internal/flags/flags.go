package flags

import (
	"flag"
	"os"
	"todo/internal/operations"
	"todo/internal/printer"
)

var (
	action = ""
)

func ReadFlags() {
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "list":
		readListFlags()
	case "add":
		readAddFlags()
	case "show":
		readShowFlags()
	case "done":
		readDoneFlags()
	case "remove":
		readRemoveFlags()
	case "clear":
		readClearFlag()
	case "help":
		readHelpFlag()
	default:
		readHelpFlag()
	}
}

func readListFlags() {
	listFlag := flag.NewFlagSet("list", flag.ContinueOnError)

	var title string
	listFlag.StringVar(&title, "title", "", "todo title")

	listFlag.Parse(os.Args[2:])

	operations.ListTodos()
}

func readAddFlags() {
	addFlag := flag.NewFlagSet("add", flag.ContinueOnError)

	var title string
	addFlag.StringVar(&title, "title", "", "todo title")

	var text string
	addFlag.StringVar(&text, "text", "", "todo text")

	addFlag.Parse(os.Args[2:])

	operations.AddTodo(title, text)
}

func readShowFlags() {
	if len(os.Args) > 2 {
		operations.ShowTodo(os.Args[2])
	}
}

func readRemoveFlags() {
	if len(os.Args) > 2 {
		operations.RemoveTodo(os.Args[2])
	}
}

func readDoneFlags() {
	if len(os.Args) > 2 {
		operations.MarkAsDone(os.Args[2])
	}
}

func readClearFlag() {
	operations.ClearTodoList()
}

func readHelpFlag() {
	printer.Help()
}
