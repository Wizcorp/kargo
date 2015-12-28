package compose

import (
	"os"
	"os/exec"
)

func Start() error {
	cmd := exec.Command("/usr/local/bin/docker-compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Stop(services []string) error {
	cmd := exec.Command("/usr/local/bin/docker-compose", append([]string{"stop"}, services...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Remove(services []string) error {
	cmd := exec.Command("/usr/local/bin/docker-compose", append([]string{"rm", "--force"}, services...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Destroy() error {
	services, err := listContainers()
	if err != nil {
		return err
	}

	if err = Stop(services); err != nil {
		return err
	}

	if err = Remove(services); err != nil {
		return err
	}

	return nil
}
