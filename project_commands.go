package stackcli

import (
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
							Value: "chef",
							Usage: "Directory Name to create",
						},
						cli.StringFlag{
							Name:  "project-name",
							Value: "stack",
							Usage: "Project Name (Your company name in lowercase eg: google)",
						},
					},
					Action: func(c *cli.Context) {
						dirName := c.String("directory-name")
						projectName := c.String("project-name")

						project := NewProject(projectName, dirName)
						project.Create()
					},
				},
			},
		},
	}
}
