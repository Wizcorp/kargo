package machine

import (
	"os"
	"os/exec"
)

type StartReturnStatus int

const (
	NOT_USING StartReturnStatus = iota
	ALREADY_RUNNING
	NOT_STARTED
	STARTED
)

func StartIfRequired(machineName string, verbose bool) (StartReturnStatus, error) {
	if !IsUsingMachine {
		return NOT_USING, nil
	}

	err := LoadMachineEnvironment(machineName, verbose)
	if err == nil {
		return ALREADY_RUNNING, nil
	}

	err = Start(machineName)
	if err == nil {
		return NOT_STARTED, err
	}
	return STARTED, nil
}

func Start(machineName string) error {
	if !IsUsingMachine {
		return nil
	}

	cmd := exec.Command("/usr/local/bin/docker-machine", "start", machineName)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	LoadMachineEnvironment(machineName, false)

	return nil
}
