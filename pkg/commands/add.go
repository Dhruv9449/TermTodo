package commands

import (
	"fmt"
	"time"

	"github.com/Dhruv9449/TermTodo/pkg/database"
	"github.com/Dhruv9449/TermTodo/pkg/models"
	"github.com/urfave/cli"
)

// AddTodo adds a todo item to the database
func addTodo(c *cli.Context) error {
	var todo_item models.TodoItem

	if c.NArg() > 0 {
		todo_item.Title = c.Args().Get(0)
	} else {
		fmt.Println("Enter task:")
		fmt.Scanln(&todo_item.Title)
	}

	todo_item.CreatedOn = time.Now()
	if c.String("desc") != "" {
		todo_item.Description = c.String("desc")
	}

	database.DB.Create(&todo_item)
	return nil
}

var AddCommand = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add a todo item",
	Action:  addTodo,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "desc",
			Usage: "Description of the todo item",
		},
	},
}
