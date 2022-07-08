// Printer, customized logs package.
package printer

import (
	"fmt"
	"io/ioutil"
	"todo/internal/check"
	"todo/internal/structs/todo"

	"github.com/alexeyco/simpletable"
)

var (
	// Header of the to-do list table
	header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Text"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "ID"},
		},
	}
)

// Print table in the terminal to display the task list.
func TodoTable(todoList []todo.Todo) {
	table := simpletable.New()
	table.Header = header

	for _, v := range todoList {
		isDone := fmt.Sprintf("%v", v.Done)
		contain := fmt.Sprintf("\"%v\"", v.Text)

		body := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: v.Title},
			{Align: simpletable.AlignCenter, Text: contain},
			{Align: simpletable.AlignCenter, Text: isDone},
			{Align: simpletable.AlignCenter, Text: v.Id},
		}

		table.Body.Cells = append(table.Body.Cells, body)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Printf("\n%v\n\n", table.String())
}

// Print message notifying the user that the operation has been successfully
// completed
func Success(functionality string) {
	msg := "was successfully completed"
	fmt.Printf("%v, %v\n", functionality, msg)
}

// System help printout
func Help() {
	file, err := ioutil.ReadFile("docs/help.txt")
	check.Err(err)

	fmt.Printf("%v", string(file))
}
