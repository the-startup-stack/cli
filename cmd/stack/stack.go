package main

import (
	"github.com/codegangsta/cli"
	"github.com/the-startup-stack/stackcli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "stack"
	app.Usage = "Startup Stack Commands"
	app.Commands = append(app.Commands, stackcli.ProjectCommands()...)

	app.Run(os.Args)
}
