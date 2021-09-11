package util

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/cli/safeexec"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// CommandExists checks if the stated command exists on the host system
func CommandExists(cmd string) bool {
	_, err := safeexec.LookPath(cmd)
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
	Pl("Running docker compose ... ")
	Pel()

	cmd := exec.Command("docker", "compose", "up", "--remove-orphans")
	if detached {
		cmd = exec.Command("docker", "compose", "up", "-d", "--remove-orphans")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}

// ExecuteWithOutput runs a command and prints the output to the console immediately
func ExecuteWithOutput(c string) {
	cmd := exec.Command(c)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// ExecuteAndWait executes a command and wait until the execution is complete
func ExecuteAndWait(c ...string) {
	// #nosec G201 G202 G203
	cmd := exec.Command(c[0], c[1:]...)
	err := cmd.Start()
	if err != nil {
		Error("Could not execute command", err, true)
	}
	err = cmd.Wait()
	if err != nil {
		Error("Could not wait for command", err, true)
	}
}

// ExecuteOnToolbox runs a command in an isolated Linux environment
func ExecuteOnToolbox(c string) {
	imageVersion := getToolboxImageVersion()
	workingDir, err := os.Getwd()
	if err != nil {
		Error("Could not find current working directory", err, true)
	}
	// Start docker container
	err = exec.Command("docker", "run", "-i", "-v", workingDir+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c).Run()
	if err != nil {
		Error("Could not start toolbox", err, true)
	}
}

// ExecuteOnToolboxCustomVolume runs a command in an isolated Linux environment with a custom volume mount
func ExecuteOnToolboxCustomVolume(c string, volumePath string) {
	imageVersion := getToolboxImageVersion()
	// Start docker container
	cmd := exec.Command("docker", "run", "-i", "-v", volumePath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
	err := cmd.Start()
	if err != nil {
		Error("Could not start docker", err, true)
	}
	err = cmd.Wait()
	if err != nil {
		Error("Could not wait for docker", err, true)
	}
}

// ClearScreen errases the console contents
func ClearScreen() {
	cmd := getClearScreenCommand()
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getClearScreenCommand() *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/c", "cls")
	}
	return exec.Command("clear")
}

func getToolboxImageVersion() string {
	if IsDevVersion() || IsPreRelease() {
		return "dev"
	}
	return Version
}
