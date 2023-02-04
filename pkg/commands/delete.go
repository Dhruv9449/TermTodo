package commands

import (
	"fmt"
	"strconv"

	"github.com/Dhruv9449/TermTodo/pkg/database"
	"github.com/Dhruv9449/TermTodo/pkg/models"
	"github.com/urfave/cli"
)

// Function to delete a todo item from the database
func deleteTodo(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.NewExitError("Please provide the id of the todo item", 1)
	}
	id, err := strconv.Atoi(c.Args().Get(0))

	if err != nil {
		return cli.NewExitError("Invalid id", 1)
	}

	if database.DB.Delete(&models.TodoItem{}, id).Error != nil {
		return cli.NewExitError("Invalid id", 1)
	} else {
		fmt.Println("Todo item deleted successfully")
	}

	return nil
}

// Command to delete a todo item
var DeleteCommand = cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "Delete a todo item",
	Action:  deleteTodo,
}
