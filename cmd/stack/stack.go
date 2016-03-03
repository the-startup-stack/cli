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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "Chef project tasks",
			Subcommands: []cli.Command{
				{
					Name:  "generate",
					Usage: "stack project generate",
					Action: func(c *cli.Context) {
						project := stackcli.NewProject()
						project.Create()
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
