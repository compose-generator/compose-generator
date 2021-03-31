package commands

import (
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

	if runtime.GOOS == "windows" { // Running on windows
		installForWindows()
	} else if runtime.GOOS == "linux" { // Running on linux
		installForLinux(flagOnlyCompose, flagOnlyDocker)
	} else { // Unknown OS
		utils.Error("Compose Generator does not support your host system yet, sorry for that.", true)
	}

	// Check if installation was successful
	if utils.CommandExists("docker") {
		utils.Pel()
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

func installForWindows() {
	const WindowsInstallerURL = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

	// Download Docker installer
	utils.P("Downloading Docker installer ... ")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := utils.DownloadFile(WindowsInstallerURL, filePath)
	if err != nil {
		utils.Error("Download of Docker installer failed", true)
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
}

func installForLinux(flagOnlyCompose bool, flagOnlyDocker bool) {
	const LinuxDockerInstallScriptURL = "https://get.docker.com"

	if utils.IsPrivileged() {
		// Install lsb_release if not installed already
		utils.ExecuteAndWait("apt-get", "update")

		// Install Docker
		if !flagOnlyCompose {
			utils.P("Installing Docker ... ")
			filePath := os.TempDir() + "/install-docker.sh"
			err := utils.DownloadFile(LinuxDockerInstallScriptURL, filePath)
			if err != nil {
				utils.Error("Download of Docker install script failed", true)
			}
			utils.ExecuteAndWait("chmod", "+x", filePath)
			utils.ExecuteAndWait("sh", filePath)
			utils.Done()
		}

		// Install Docker Compose
		if !flagOnlyDocker {
			utils.P("Installing Docker Compose ... ")
			cmd := exec.Command("uname", "-s")
			output, _ := cmd.CombinedOutput()
			unameS := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("uname", "-m")
			output, _ = cmd.CombinedOutput()
			unameM := strings.TrimRight(string(output), "\r\n")

			err := utils.DownloadFile("https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+unameS+"-"+unameM, "/usr/local/bin/docker-compose")
			if err != nil {
				utils.Error("Download of Docker Compose failed", true)
			}
			utils.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
			utils.Done()
		}
	} else {
		color.Red("Please execute this command with root privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
		return
	}
}
