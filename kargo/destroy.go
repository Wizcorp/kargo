package main

import (
	"github.com/codegangsta/cli"
	"github.com/wizcorp/kargo/compose"
	"github.com/wizcorp/kargo/machine"
	"os"
	"os/exec"
	"syscall"
)

func init() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "destroy",
		Usage: "destroy the docker containers",
		Action: func(c *cli.Context) {
			_, err := machine.StartIfRequired(c.GlobalString("machine"), false)
			if err != nil {
				panic(err)
			}

			projectName := c.GlobalString("project-name")
			if len(projectName) > 0 {
				os.Setenv("COMPOSE_PROJECT_NAME", projectName)
			}

			err = compose.Destroy()
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					// The program has exited with an exit code != 0

					// This works on both Unix and Windows. Although package
					// syscall is generally platform dependent, WaitStatus is
					// defined for both Unix and Windows and in both cases has
					// an ExitStatus() method with the same signature.
					if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
						os.Exit(status.ExitStatus())
					}
				}
				panic(err)
			}
		},
	})
}
