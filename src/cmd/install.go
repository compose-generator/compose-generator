package cmd

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"compose-generator/util"

	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(flagOnlyCompose bool, flagOnlyDocker bool) {
	// Check if Compose Generator runs in dockerized environment
	if util.IsDockerized() {
		util.Error("You are currently using the Docker container of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", nil, true)
	}

	if runtime.GOOS == "windows" { // Running on windows
		installForWindows()
	} else if runtime.GOOS == "linux" { // Running on linux
		installForLinux(flagOnlyCompose, flagOnlyDocker)
	} else { // Unknown OS
		util.Error("Compose Generator does not support your host system yet, sorry for that.", nil, true)
	}

	// Check if installation was successful
	if util.CommandExists("docker") {
		util.Pel()
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
	util.P("Downloading Docker installer ... ")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := util.DownloadFile(WindowsInstallerURL, filePath)
	if err != nil {
		util.Error("Download of Docker installer failed", err, true)
	}
	util.Done()
	// Run Docker installer
	util.Pl("Starting installation ... ")
	util.Pel()
	cmd := exec.Command(filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	util.Pel()
}

func installForLinux(flagOnlyCompose bool, flagOnlyDocker bool) {
	const LinuxDockerInstallScriptURL = "https://get.docker.com"

	if util.IsPrivileged() {
		// Install lsb_release if not installed already
		util.ExecuteAndWait("apt-get", "update")

		// Install Docker
		if !flagOnlyCompose {
			util.P("Installing Docker ... ")
			filePath := os.TempDir() + "/install-docker.sh"
			err := util.DownloadFile(LinuxDockerInstallScriptURL, filePath)
			if err != nil {
				util.Error("Download of Docker install script failed", err, true)
			}
			util.ExecuteAndWait("chmod", "+x", filePath)
			util.ExecuteAndWait("sh", filePath)
			util.Done()
		}

		// Install Docker Compose
		if !flagOnlyDocker {
			util.P("Installing Docker Compose ... ")
			cmd := exec.Command("uname", "-s")
			output, _ := cmd.CombinedOutput()
			unameS := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("uname", "-m")
			output, _ = cmd.CombinedOutput()
			unameM := strings.TrimRight(string(output), "\r\n")

			err := util.DownloadFile("https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+unameS+"-"+unameM, "/usr/local/bin/docker-compose")
			if err != nil {
				util.Error("Download of Docker Compose failed", err, true)
			}
			util.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
			util.Done()
		}
	} else {
		color.Red("Please execute this command with root privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
		return
	}
}
