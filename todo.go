package main

import (
	"log"
	"os"

	"github.com/Dhruv9449/TermTodo/pkg/commands"
	"github.com/Dhruv9449/TermTodo/pkg/database"
	"github.com/urfave/cli"
)

func main() {
	database.Connect()

	app := cli.NewApp()

	app.Name = "TermTodo"
	app.Usage = "A simple todo app for the terminal"
	app.Version = "0.0.1"
	app.Author = "Dhruv Shah"

	app.Commands = []cli.Command{
		commands.AddCommand,
		commands.ShowCommand,
		commands.DeleteCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
