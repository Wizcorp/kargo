package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wizcorp/kargo/compose"
	"github.com/wizcorp/kargo/machine"
	"github.com/wizcorp/kargo/utils"
	"os"
)

func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "open",
		Aliases: []string{"o"},
		Usage:   "open the project in your browser",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "container, c",
				Value: "web",
				Usage: "name of the docker container to use",
			},
			cli.StringFlag{
				Name:  "suffix, s",
				Usage: "suffix to append to the url",
			},
		},
		Action: func(c *cli.Context) {
			_, err := machine.StartIfRequired(c.GlobalString("machine"), false)
			if err != nil {
				panic(err)
			}

			projectName := c.GlobalString("project-name")
			if len(projectName) > 0 {
				os.Setenv("COMPOSE_PROJECT_NAME", projectName)
			}

			containerName := c.String("container")
			port, err := compose.GetPort(containerName)
			if err != nil {
				fmt.Println(port)
				panic(err)
			}
			ip := machine.GetIP()
			url := "http://" + ip + ":" + port + c.String("suffix")

			fmt.Println("Opening", url, "...")
			err = utils.Open(url)
			if err != nil {
				panic(err)
			}
		},
	})
}
