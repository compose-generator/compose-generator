/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"bytes"
	"os"
	"os/exec"
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
	if isLinux() {
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
	} else if isWindows() {
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		return err == nil
	}
	return false
}

// DockerComposeUp executes 'docker compose up' in the current directory
func DockerComposeUp(detached bool, production bool) {
	Pel()
	Pl("Running docker compose ... ")
	Pel()

	var cmd *exec.Cmd
	if production {
		if detached {
			cmd = exec.Command("docker", "compose", "--profile", "prod", "up", "-d", "--remove-orphans")
		} else {
			cmd = exec.Command("docker", "compose", "--profile", "prod", "up", "--remove-orphans")
		}
	} else {
		if detached {
			cmd = exec.Command("docker", "compose", "up", "-d", "--remove-orphans")
		} else {
			cmd = exec.Command("docker", "compose", "up", "--remove-orphans")
		}
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		ErrorLogger.Println("Could not execute Docker Compose: " + err.Error())
		logError("Could not execute Docker Compose", true)
	}
	if err := cmd.Wait(); err != nil {
		ErrorLogger.Println("Could not wait for Docker Compose: " + err.Error())
		logError("Could not wait for Docker Compose", true)
	}
}

// ExecuteWithOutput runs a command and prints the output to the console immediately
func ExecuteWithOutput(c string) {
	// #nosec G204
	cmd := exec.Command(c)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ErrorLogger.Println("Could not execute command '" + c + "': " + err.Error())
		logError("Could not execute command", true)
	}
}

// ExecuteAndWait executes a command and wait until the execution is complete
func ExecuteAndWait(c ...string) {
	// #nosec G204
	cmd := exec.Command(c[0], c[1:]...)
	if err := cmd.Start(); err != nil {
		ErrorLogger.Println("Could not execute command '" + strings.Join(c, " ") + "': " + err.Error())
		logError("Could not execute command", true)
	}
	if err := cmd.Wait(); err != nil {
		ErrorLogger.Println("Could not wait for command '" + strings.Join(c, " ") + "': " + err.Error())
		logError("Could not wait for command", true)
	}
}

// ExecuteOnToolbox runs a command in an isolated Linux environment
func ExecuteOnToolbox(c string) {
	imageVersion := getToolboxImageVersion()
	toolboxMountPath := getToolboxMountPath()
	// Start docker container
	// #nosec G204
	cmd := exec.Command("docker", "run", "-i", "-v", toolboxMountPath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	DebugLogger.Print("Toolbox stdout:\n" + outb.String())
	DebugLogger.Print("Toolbox stderr:\n" + errb.String())
	if err != nil {
		ErrorLogger.Println("Toolbox terminated with an error for command '" + c + "': " + err.Error())
		logError("Toolbox terminated with an error", true)
	}
}

// ExecuteOnToolboxCustomVolume runs a command in an isolated Linux environment with a custom volume mount
func ExecuteOnToolboxCustomVolume(c string, volumePath string) {
	imageVersion := getToolboxImageVersion()
	// Start docker container
	// #nosec G204
	cmd := exec.Command("docker", "run", "-i", "-v", volumePath+":/toolbox", "chillibits/compose-generator-toolbox:"+imageVersion, c)
	if err := cmd.Start(); err != nil {
		ErrorLogger.Println("Could not execute Toolbox command '" + c + "': " + err.Error())
		logError("Could not start Docker", true)
	}
	if err := cmd.Wait(); err != nil {
		ErrorLogger.Println("Could not wait for Toolbox command '" + c + "': " + err.Error())
		logError("Could not wait for Docker", true)
	}
}

// ClearScreen errases the console contents
func ClearScreen() {
	cmd := exec.Command("clear")
	if isWindows() {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		WarningLogger.Println("Could not clear screen: " + err.Error())
		logWarning("Could not clear screen")
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
	}
	// Get the current working directory
	workingDir, err := getwd()
	if err != nil {
		ErrorLogger.Println("Could not find current working directory: " + err.Error())
		logError("Could not find current working directory", true)
	}
	return workingDir
}
