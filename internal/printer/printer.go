package printer

import (
	"fmt"
	"io/ioutil"
	"todo/internal/check"
	"todo/internal/structs/todo"

	"github.com/alexeyco/simpletable"
)

func TodoTable(todoList []todo.Todo) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Text"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "ID"},
		},
	}

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

func Success(functionality string) {
	msg := "was successfully completed"
	fmt.Printf("%v, %v\n", functionality, msg)
}

func Help() {
	file, err := ioutil.ReadFile("docs/help.txt")
	check.Err(err)

	fmt.Printf("%v", string(file))
}
