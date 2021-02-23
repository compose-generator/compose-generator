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

const WINDOWS_INSTALLER = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

func Install(flag_only_compose bool, flag_only_docker bool) {
	if runtime.GOOS == "windows" { // Running on windows
		// Download Docker installer
		fmt.Print("Downloading Docker installer ...")
		file_path := os.TempDir() + "/DockerInstaller.exe"
		err := utils.DownloadFile(WINDOWS_INSTALLER, file_path)
		if err != nil {
			panic(err)
		}
		color.Green(" done")
		// Run Docker installer
		fmt.Println("Starting installation ...")
		fmt.Println()
		cmd := exec.Command(file_path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		fmt.Println()
	} else { // Running on linux
		if utils.IsPrivileged() {
			// Install lsb_release if not installed already
			utils.ExecuteAndWait("apt-get", "update")

			// Install Docker
			if !flag_only_compose {
				utils.ExecuteAndWait("curl", "-fsSL", "https://get.docker.com", "|", "sh")
			}

			// Install Docker Compose
			if !flag_only_docker {
				cmd := exec.Command("uname", "-s")
				output, _ := cmd.CombinedOutput()
				uname_s := strings.TrimRight(string(output), "\r\n")

				cmd = exec.Command("uname", "-m")
				output, _ = cmd.CombinedOutput()
				uname_m := strings.TrimRight(string(output), "\r\n")

				utils.ExecuteAndWaitWithOutput("curl", "-L", "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+uname_s+"-"+uname_m, "-o", "/usr/local/bin/docker-compose")
				utils.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
			}
		} else {
			color.Red("Please execute this command with root privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
		}
	}

	// Check if installation was successful
	if utils.CommandExists("docker") {
		cmd := exec.Command("docker", "-v")
		docker_version, _ := cmd.CombinedOutput()
		cmd = exec.Command("docker-compose", "-v")
		compose_version, _ := cmd.CombinedOutput()
		if flag_only_compose {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(compose_version), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else if flag_only_docker {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(docker_version), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(docker_version), "\r\n") + " and " + strings.TrimRight(string(compose_version), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		}
	} else {
		color.Red("An error occured while installing Docker.")
	}
}
