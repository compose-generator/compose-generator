package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"compose-generator/utils"

	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(flagOnlyCompose bool, flagOnlyDocker bool) {
	// Check if Compose Generator runs in dockerized environment
	if utils.IsDockerized() {
		utils.Error("You are currently using the Docker container of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", true)
	}

	const WindowsInstallerURL = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

	if runtime.GOOS == "windows" { // Running on windows
		// Download Docker installer
		fmt.Print("Downloading Docker installer ... ")
		filePath := os.TempDir() + "/DockerInstaller.exe"
		err := utils.DownloadFile(WindowsInstallerURL, filePath)
		if err != nil {
			panic(err)
		}
		utils.Done()
		// Run Docker installer
		utils.Pl("Starting installation ... ")
		utils.Pel()
		cmd := exec.Command(filePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		utils.Pel()
	} else if runtime.GOOS == "linux" { // Running on linux
		if utils.IsPrivileged() {
			// Install lsb_release if not installed already
			utils.ExecuteAndWait("apt-get", "update")

			// Install Docker
			if !flagOnlyCompose {
				utils.ExecuteAndWait("wget", "https://get.docker.com", "-O", "install-docker.sh")
				utils.ExecuteAndWait("chmod", "+x", "install-docker.sh")
				utils.ExecuteAndWait("sh", "install-docker.sh")
				utils.ExecuteAndWait("rm", "install-docker.sh")
			}

			// Install Docker Compose
			if !flagOnlyDocker {
				cmd := exec.Command("uname", "-s")
				output, _ := cmd.CombinedOutput()
				unameS := strings.TrimRight(string(output), "\r\n")

				cmd = exec.Command("uname", "-m")
				output, _ = cmd.CombinedOutput()
				unameM := strings.TrimRight(string(output), "\r\n")

				utils.ExecuteAndWaitWithOutput("wget", "-O", "/usr/local/bin/docker-compose", "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+unameS+"-"+unameM)
				utils.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
			}
		} else {
			color.Red("Please execute this command with root privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
			return
		}
	} else {
		utils.Error("Compose Generator does not support your host system yet, sorry for that.", true)
	}

	// Check if installation was successful
	if utils.CommandExists("docker") {
		cmd := exec.Command("docker", "-v")
		dockerVersion, _ := cmd.CombinedOutput()
		cmd = exec.Command("docker-compose", "-v")
		composeVersion, _ := cmd.CombinedOutput()
		if flagOnlyCompose {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(composeVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else if flagOnlyDocker {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(dockerVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(dockerVersion), "\r\n") + " and " + strings.TrimRight(string(composeVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		}
	} else {
		color.Red("An error occurred while installing Docker.")
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------
