package util

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

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
	imageVersion := getToolboxImageVersion()
	// Start docker container
	absolutePath, _ := os.Getwd()
	ExecuteAndWaitWithOutput("docker", "run", "-i", "-v", absolutePath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
}

// ExecuteOnLinuxWithCustomVolume runs a command in an isolated Linux environment with a custom volume mount
func ExecuteOnLinuxWithCustomVolume(c string, volumePath string) {
	imageVersion := getToolboxImageVersion()
	// Start docker container
	ExecuteAndWait("docker", "run", "-i", "-v", volumePath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
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
