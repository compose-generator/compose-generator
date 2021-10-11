/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

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
	if runtime.GOOS == "linux" {
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
	} else if runtime.GOOS == "windows" {
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		return err == nil
	}
	return false
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
	if err := cmd.Start(); err != nil {
		Error("Could not execute docker compose", err, true)
	}
	if err := cmd.Wait(); err != nil {
		Error("Could not wait for docker compose", err, true)
	}
}

// ExecuteWithOutput runs a command and prints the output to the console immediately
func ExecuteWithOutput(c string) {
	// #nosec G204
	cmd := exec.Command(c)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		Error("Could not execute command", err, true)
	}
}

// ExecuteAndWait executes a command and wait until the execution is complete
func ExecuteAndWait(c ...string) {
	// #nosec G204
	cmd := exec.Command(c[0], c[1:]...)
	if err := cmd.Start(); err != nil {
		Error("Could not execute command", err, true)
	}
	if err := cmd.Wait(); err != nil {
		Error("Could not wait for command", err, true)
	}
}

// ExecuteOnToolbox runs a command in an isolated Linux environment
func ExecuteOnToolbox(c string) {
	imageVersion := getToolboxImageVersion()
	toolboxMountPath := getToolboxMountPath()
	// Start docker container
	// #nosec G204
	cmd := exec.Command("docker", "run", "-i", "-v", toolboxMountPath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
	if err := cmd.Run(); err != nil {
		Error("Toolbox terminated with an error", err, true)
	}
}

// ExecuteOnToolboxCustomVolume runs a command in an isolated Linux environment with a custom volume mount
func ExecuteOnToolboxCustomVolume(c string, volumePath string) {
	imageVersion := getToolboxImageVersion()
	// Start docker container
	// #nosec G204
	cmd := exec.Command("docker", "run", "-i", "-v", volumePath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
	if err := cmd.Start(); err != nil {
		Error("Could not start docker", err, true)
	}
	if err := cmd.Wait(); err != nil {
		Error("Could not wait for docker", err, true)
	}
}

// ClearScreen errases the console contents
func ClearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		Warning("Could not clear screen")
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getToolboxImageVersion() string {
	if IsDevVersion() || IsPreRelease() {
		return "dev"
	}
	return Version
}

func getToolboxMountPath() string {
	if isDockerizedEnvironment() { // Get the path which is mounted to /cg/out from outside the container
		return getOuterVolumePathOnDockerizedEnvironmentMockable()
	} else { // Get the current working directory
		workingDir, err := getwd()
		if err != nil {
			printError("Could not find current working directory", err, true)
		}
		return workingDir
	}
}
