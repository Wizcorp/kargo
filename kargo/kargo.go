package main

import (
	"github.com/codegangsta/cli"
	_ "github.com/wizcorp/kargo/utils"
	"os"
)

var app = setupApp()

func setupApp() *cli.App {
	app := cli.NewApp()

	app.Name = "kargo"
	app.Usage = "using docker should be easy"
	app.Commands = []cli.Command{}

	setupGlobalFlags(app)

	return app
}

func setupGlobalFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "machine, m",
			Value: "default",
			Usage: "name of the docker machine to use",
		},
		cli.StringFlag{
			Name:  "project-name, p",
			Usage: "name of the project",
		},
	}
}

func main() {
	app.Run(os.Args)
}
