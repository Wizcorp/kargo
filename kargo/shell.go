package main

import (
	"github.com/codegangsta/cli"
	"github.com/wizcorp/kargo/docker"
	"github.com/wizcorp/kargo/utils"
)

func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "shell",
		Aliases: []string{"s"},
		Usage:   "open a shell into a container",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "container, c",
				Value: "web",
				Usage: "name of the docker container to use",
			},
		},
		Action: func(c *cli.Context) {
			projectName, err := utils.GetProjectName()
			if err != nil {
				panic(err)
			}

			err = docker.Exec(projectName+"_"+c.String("container")+"_1", "bash")
			if err != nil {
				panic(err)
			}
		},
	})
}
