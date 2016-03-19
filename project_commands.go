package stackcli

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func ProjectCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "Chef project tasks",
			Subcommands: []cli.Command{
				{
					Name:  "generate",
					Usage: "stack project generate",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "directory-name",
							Usage: "Directory Name to create",
						},
						cli.StringFlag{
							Name:  "project-name",
							Usage: "Project Name (Your company name in lowercase eg: google)",
						},
					},

					Action: func(c *cli.Context) {
						dirName := c.String("directory-name")
						projectName := c.String("project-name")

						if len(projectName) <= 0 || len(dirName) <= 0 {
							fmt.Println("You need to pass project-name and directory-name flags for the command to work")
							return
						}

						project := NewProject(projectName, dirName)
						project.Create()
					},
				},
			},
		},
	}
}
