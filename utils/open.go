package utils

import (
	"errors"
	"os"
	"os/exec"
)

func getOpenProgram() string {
	possiblePaths := []string{
		"/usr/bin/open",
		"/usr/bin/xdg-open",
	}

	for _, path := range possiblePaths {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			return path
		}
	}
	return ""
}

func Open(path string) error {
	program := getOpenProgram()
	if program == "" {
		return errors.New("No program found to open " + path)
	}

	cmd := exec.Command(program, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
