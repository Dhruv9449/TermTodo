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

	cli.AppHelpTemplate = "  _______               _______        _       \n |__   __|             |__   __|      | |      \n    | | ___ _ __ _ __ ___ | | ___   __| | ___  \n    | |/ _ \\ '__| '_ ` _ \\| |/ _ \\ / _` |/ _ \\ \n    | |  __/ |  | | | | | | | (_) | (_| | (_) |\n    |_|\\___|_|  |_| |_| |_|_|\\___/ \\__,_|\\___/ \n" +
		`
A very simple CLI todo app written in Go!

Developer :
	Dhruv Shah (https://github.com/Dhruv9449)

For more detailed help:
	https://github.com/Dhruv9449/TermTodo/blob/main/README.md
		
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
VERSION:
   {{.Version}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

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
