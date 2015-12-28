package compose

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"strings"
)

/**
 * @todo Use docker-compose config when 1.6 is released
 */
func listContainers() ([]string, error) {
	m := make(map[string]interface{})

	data, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, err
	}

	services := []string{}
	for service := range m {
		services = append(services, service)
	}
	return services, nil
}

func GetPort(container string) (string, error) {
	cmd := exec.Command("/usr/local/bin/docker-compose", "port", container, "80")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	port := strings.Split(strings.TrimSpace(string(output)), ":")
	return port[1], nil
}
