package util

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// CommandExists checks if the stated command exists on the host system
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// IsPrivileged checks if the user has root priviledges
func IsPrivileged() bool {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	// 0 = root, 501 = non-root user
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		panic(err)
	}
	return i == 0
}

// DockerComposeUp executes 'docker compose up' in the current directory
func DockerComposeUp(detached bool) {
	Pel()
	Pl("Running docker-compose ... ")
	Pel()

	cmd := exec.Command("docker-compose", "up", "--remove-orphans")
	if detached {
		cmd = exec.Command("docker-compose", "up", "-d", "--remove-orphans")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}

// ExecuteAndWait executes a command and wait till the execution is complete
func ExecuteAndWait(c ...string) {
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Start()
	cmd.Wait()
}

// ExecuteAndWaitWithOutput executes a command and return the command output as string
func ExecuteAndWaitWithOutput(c ...string) string {
	cmd := exec.Command(c[0], c[1:]...)
	output, _ := cmd.CombinedOutput()
	return strings.TrimRight(string(output), "\r\n")
}

// ExecuteOnLinux runs a command in an isolated Linux environment
func ExecuteOnLinux(c string) {
	// Start docker container
	absolutePath, _ := os.Getwd()
	ExecuteAndWait("docker", "run", "-i", "-v", absolutePath+":/toolbox", "chillibits/compose-generator-toolbox", c)
}

// ClearScreen errases the console contents
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
