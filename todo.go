// ████████╗███████╗██████╗ ███╗   ███╗████████╗ ██████╗ ██████╗  ██████╗
// ╚══██╔══╝██╔════╝██╔══██╗████╗ ████║╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗
//    ██║   █████╗  ██████╔╝██╔████╔██║   ██║   ██║   ██║██║  ██║██║   ██║
//    ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║   ██║   ██║   ██║██║  ██║██║   ██║
//    ██║   ███████╗██║  ██║██║ ╚═╝ ██║   ██║   ╚██████╔╝██████╔╝╚██████╔╝
//    ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝

// A very simple CLI todo app written in Go

// Version - 0.0.1
// License - MIT

// Developed with 🩵 by Dhruv Shah

package main

import (
	"log"
	"os"

	"github.com/Dhruv9449/TermTodo/pkg/commands"
	"github.com/Dhruv9449/TermTodo/pkg/database"
	"github.com/urfave/cli"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new cli app
	app := cli.NewApp()

	// Set the app details
	app.Name = "TermTodo"
	app.Usage = "A simple todo app for the terminal"
	app.Version = "0.0.1"
	app.Author = "Dhruv Shah"

	app.Commands = []cli.Command{
		commands.AddCommand,
		commands.ShowCommand,
		commands.DeleteCommand,
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
