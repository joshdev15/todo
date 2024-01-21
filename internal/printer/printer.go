// Printer, customized todo-cli logs package.
package printer

import (
	"fmt"
	"todo/internal/check"
	"todo/internal/structs/todo"
	"todo/internal/texts"

	"github.com/alexeyco/simpletable"
)

// Print table in the terminal to display the task list.
func TodoTable(todoList []todo.Todo) {
	table := simpletable.New()

	// Todo List Header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Message"},
			{Align: simpletable.AlignCenter, Text: "Done"},
		},
	}

	// Todo List Body
	for _, v := range todoList {
		isDone := fmt.Sprintf("%v", v.Done)

		body := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: v.Id},
			{Align: simpletable.AlignCenter, Text: v.Title.Elipsis()},
			{Align: simpletable.AlignCenter, Text: v.Message.Elipsis()},
			{Align: simpletable.AlignCenter, Text: isDone},
		}

		table.Body.Cells = append(table.Body.Cells, body)
	}

	// Todo List Style
	table.SetStyle(simpletable.StyleRounded)

	// Print Todo List
	fmt.Printf("\n%v\n\n", table.String())
}

// Print message notifying the user that the operation has been successfully
// completed
func Success(functionality string) {
	message := "was successfully completed"
	fmt.Printf("%v, %v\n", functionality, message)
}

// Print to-do
func Show(currentTodo todo.Todo) {
	fmt.Println(currentTodo)
	contain := fmt.Sprintf("\n%v\n", currentTodo.Message)
	footerContain := fmt.Sprintf("%v | Is done: %v", currentTodo.Id, currentTodo.Done)

	table := simpletable.New()

	// Todo List Header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: currentTodo.Title.ToString()},
		},
	}

	// Todo List Body
	table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: contain},
	})

	// Todo List Footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: footerContain},
		},
	}

	// Todo List Style
	table.SetStyle(simpletable.StyleRounded)

	// Print Todo List
	fmt.Printf("\n%v\n\n", table.String())
}

// System help printout
func Help() {
	for _, line := range texts.HelpText {
		fmt.Printf("%v\n", line)
	}

}

func Version() {
	for _, line := range texts.VersionText {
		fmt.Printf("%v\n", line)
	}
}

// Print message when no valid argument exists
func NoArgs() {
	fmt.Println("The todo id to process is missing")
}

func Modal(action string) bool {
	fmt.Printf("You are sure you want %v [y / n]\n", action)
	answer := "n"
	_, err := fmt.Scanln(&answer)
	check.Err(err)

	finalValue := false

	switch answer {
	case "y":
		fmt.Println("Action approved")
		finalValue = true
	case "n":
		fmt.Println("Action rejected")
		finalValue = false
	default:
		fmt.Println("Invalid answer, rejected action")
		finalValue = false
	}

	return finalValue
}
