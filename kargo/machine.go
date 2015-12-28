package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wizcorp/kargo/machine"
	"os"
)

func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "machine",
		Usage: "start the docker machine if required",
		Action: func(c *cli.Context) {
			_, environmentPresent := os.LookupEnv("DOCKER_HOST")
			status, err := machine.StartIfRequired(c.GlobalString("machine"), !environmentPresent)
			if err != nil {
				panic(err)
			}
			switch status {
			case machine.NOT_USING:
				fmt.Println("You are not using Docker Machine.")
			case machine.ALREADY_RUNNING:
				fmt.Println("Nothing to do. The Machine is already running.")
			case machine.STARTED:
				fmt.Println("The Machine has been started.")
			}
		},
	})
}
