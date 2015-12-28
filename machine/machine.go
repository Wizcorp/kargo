package machine

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var IsUsingMachine bool

func init() {
	_, err := os.Stat("/usr/local/bin/docker-machine")
	IsUsingMachine = !os.IsNotExist(err)
}

func LoadMachineEnvironment(machineName string, verbose bool) error {
	if !IsUsingMachine {
		return nil
	}

	cmd := exec.Command("/usr/local/bin/docker-machine", "env", machineName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(output) > 0 {
		exportRegexp := regexp.MustCompile(`^export ([A-Z_]+)="([^\"]+)"`)
		for _, str := range strings.Split(string(output), "\n") {
			if !strings.HasPrefix(str, "export ") {
				if verbose == true {
					fmt.Println(str)
				}
				continue
			}
			match := exportRegexp.FindStringSubmatch(str)
			os.Setenv(match[1], match[2])
		}
	}
	return nil
}

func GetIP() string {
	if !IsUsingMachine {
		return "127.0.0.1"
	}

	cmd := exec.Command("/usr/local/bin/docker-machine", "ip", os.Getenv("DOCKER_MACHINE_NAME"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}
