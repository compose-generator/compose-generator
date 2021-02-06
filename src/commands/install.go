package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"compose-generator/utils"

	"github.com/fatih/color"
)

const WINDOWS_INSTALLER = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

func Install() {
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
		color.Green(" done")
		fmt.Println()
	} else { // Running on linux
		if utils.GetProcessOwner() == "root" {
			// Get distribution name
			cmd := exec.Command("lsb_release -is")
			distribution, _ := cmd.CombinedOutput()
			cmd = exec.Command("arch")
			arch, _ := cmd.CombinedOutput()
			switch string(distribution) {
			case "CentOS":
				// Uninstall old versions
				exec.Command("sudo yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine")
				// Add repository
				exec.Command("sudo yum install -y yum-utils")
				exec.Command("sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo")
				// Install Docker
				exec.Command("sudo yum install docker-ce docker-ce-cli containerd.io")
				// Start Docker demon
				exec.Command("sudo systemctl start docker")
			case "Debian":
				// Uninstall old versions
				exec.Command("sudo apt-get remove docker docker-engine docker.io containerd runc")
				// Add repository
				exec.Command("sudo apt-get update")
				exec.Command("sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common")
				exec.Command("curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -")

				switch string(arch) {
				case "x86_64", "amd64":
					exec.Command("sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable\"")
				case "armhf":
					exec.Command("sudo add-apt-repository \"deb [arch=armhf] https://download.docker.com/linux/debian $(lsb_release -cs) stable\"")
				case "arm64":
					exec.Command("sudo add-apt-repository \"deb [arch=arm64] https://download.docker.com/linux/debian $(lsb_release -cs) stable\"")
				}
				// Install Docker
				exec.Command("sudo apt-get update")
				exec.Command("sudo apt-get install docker-ce docker-ce-cli containerd.io")
			case "Fedora":
				// Uninstall old versions
				exec.Command("sudo dnf remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-selinux docker-engine-selinux docker-engine")
				// Add repository
				exec.Command("sudo dnf -y install dnf-plugins-core")
				exec.Command("sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo")
				// Install Docker
				exec.Command("sudo dnf install docker-ce docker-ce-cli containerd.io")
			case "Ubuntu":
				// Uninstall old versions
				exec.Command("sudo apt-get remove docker docker-engine docker.io containerd runc")
				// Add repository
				exec.Command("sudo apt-get update")
				exec.Command("sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common")
				exec.Command("curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -")

				switch string(arch) {
				case "x86_64", "amd64":
					exec.Command("sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"")
				case "armhf":
					exec.Command("sudo add-apt-repository \"deb [arch=armhf] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"")
				case "arm64":
					exec.Command("sudo add-apt-repository \"deb [arch=arm64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"")
				}
				// Install Docker
				exec.Command("sudo apt-get update")
				exec.Command("sudo apt-get install docker-ce docker-ce-cli containerd.io")
			}
		} else {
			color.Red("Please execute this command with sudo privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
		}
	}

	// Check if installation was successful
	if utils.CommandExists("docker") {
		cmd := exec.Command("docker", "-v")
		version, _ := cmd.CombinedOutput()
		color.Yellow("Congrats! You have installed " + string(version) + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
	} else {
		color.Red("An error occured while installing Docker.")
	}
}
