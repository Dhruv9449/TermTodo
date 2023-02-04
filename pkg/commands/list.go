package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Dhruv9449/TermTodo/pkg/database"
	"github.com/Dhruv9449/TermTodo/pkg/models"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// ListTodo shows all todo items
func showTodo(c *cli.Context) error {
	var todo_item []models.TodoItem
	database.DB.Find(&todo_item)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Desc", "Created on"})

	for _, item := range todo_item {
		table.Append([]string{
			strconv.Itoa(int(item.ID)),
			item.Title,
			item.Description,
			item.CreatedOn.Format("Jan 2"),
		})
	}
	table.SetColumnSeparator(" ")
	table.SetBorder(true)
	table.SetCaption(true, fmt.Sprintf("\nYou have %d todo items", len(todo_item)))
	table.Render()
	return nil
}

var ShowCommand = cli.Command{
	Name:    "show",
	Aliases: []string{"s"},
	Usage:   "Show a todo items",
	Action:  showTodo,
}
