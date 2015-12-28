package docker
import (
    "os/exec"
    "os"
)

func Exec(containerName string, command string) error {
    os.Setenv("TERM", "xterm")
    cmd := exec.Command("/usr/local/bin/docker", "exec", "-t", "-i", containerName, command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    return cmd.Run()
}
