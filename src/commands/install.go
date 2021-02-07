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
	} else { // Running on linux
		if utils.IsPrivileged() {
			// Install lsb_release if not installed already
			utils.ExecuteAndWait("apt-get", "update")
			utils.ExecuteAndWait("apt-get", "install", "lsb-release")

			// Get distribution name
			cmd := exec.Command("lsb_release", "-is")
			output, _ := cmd.CombinedOutput()
			distribution_is := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("lsb_release", "-cs")
			output, _ = cmd.CombinedOutput()
			distribution_cs := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("arch")
			output, _ = cmd.CombinedOutput()
			arch := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("uname", "-s")
			output, _ = cmd.CombinedOutput()
			uname_s := strings.TrimRight(string(output), "\r\n")

			cmd = exec.Command("uname", "-m")
			output, _ = cmd.CombinedOutput()
			uname_m := strings.TrimRight(string(output), "\r\n")

			switch distribution_is {
			case "CentOS":
				// Uninstall old versions
				utils.ExecuteAndWait("yum", "remove", "docker", "docker-client", "docker-client-latest", "docker-common", "docker-latest", "docker-latest-logrotate", "docker-logrotate", "docker-engine")
				// Add repository
				utils.ExecuteAndWait("yum", "install", "-y", "yum-utils")
				utils.ExecuteAndWait("yum-config-manager", "--add-repo", "https://download.docker.com/linux/centos/docker-ce.repo")
				// Install Docker
				utils.ExecuteAndWait("yum", "install", "docker-ce", "docker-ce-cli", "containerd.io")
				// Start Docker demon
				utils.ExecuteAndWait("systemctl", "start", "docker")
			case "Debian":
				// Uninstall old versions
				utils.ExecuteAndWait("apt-get", "remove docker", "docker-engine", "docker.io", "containerd", "runc")
				// Add repository
				utils.ExecuteAndWait("apt-get", "update")
				utils.ExecuteAndWait("apt-get", "install", "apt-transport-https", "ca-certificates", "curl", "gnupg-agent", "software-properties-common", "curl")
				utils.ExecuteAndWait("wget", "https://download.docker.com/linux/ubuntu/gpg")
				utils.ExecuteAndWait("apt-key", "add", "gpg")
				utils.ExecuteAndWait("rm", "gpg")

				switch arch {
				case "x86_64", "amd64":
					utils.ExecuteAndWait("add-apt-repository", "deb [arch=amd64] https://download.docker.com/linux/debian "+distribution_cs+" stable")
				case "armhf":
					utils.ExecuteAndWait("add-apt-repository", "deb [arch=armhf] https://download.docker.com/linux/debian "+distribution_cs+" stable")
				case "arm64":
					utils.ExecuteAndWait("add-apt-repository", "deb [arch=arm64] https://download.docker.com/linux/debian "+distribution_cs+" stable")
				}
				// Install Docker
				utils.ExecuteAndWait("apt-get", "update")
				utils.ExecuteAndWait("apt-get", "install", "docker-ce", "docker-ce-cli", "containerd.io")
			case "Fedora":
				// Uninstall old versions
				utils.ExecuteAndWait("dnf", "remove", "docker", "docker-client", "docker-client-latest", "docker-common", "docker-latest", "docker-latest-logrotate", "docker-logrotate", "docker-selinux", "docker-engine-selinux", "docker-engine")
				// Add repository
				utils.ExecuteAndWait("dnf", "-y", "install", "dnf-plugins-core")
				utils.ExecuteAndWait("dnf", "config-manager", "--add-repo", "https://download.docker.com/linux/fedora/docker-ce.repo")
				// Install Docker
				utils.ExecuteAndWait("dnf", "install", "docker-ce", "docker-ce-cli", "containerd.io")
			case "Ubuntu":
				// Uninstall old versions
				utils.ExecuteAndWaitWithOutput("apt-get", "remove", "docker", "docker-engine", "docker.io", "containerd", "runc", "-y")

				// Add repository
				utils.ExecuteAndWaitWithOutput("apt-get", "update")
				utils.ExecuteAndWaitWithOutput("apt-get", "install", "apt-transport-https", "ca-certificates", "curl", "gnupg-agent", "software-properties-common", "curl", "-y")
				utils.ExecuteAndWaitWithOutput("wget", "https://download.docker.com/linux/ubuntu/gpg")
				utils.ExecuteAndWaitWithOutput("apt-key", "add", "gpg")
				utils.ExecuteAndWaitWithOutput("rm", "gpg")

				switch arch {
				case "x86_64", "amd64":
					utils.ExecuteAndWaitWithOutput("add-apt-repository", "deb [arch=amd64] https://download.docker.com/linux/ubuntu "+distribution_cs+" stable")
				case "armhf":
					utils.ExecuteAndWaitWithOutput("add-apt-repository", "deb [arch=armhf] https://download.docker.com/linux/ubuntu "+distribution_cs+" stable")
				case "arm64":
					utils.ExecuteAndWaitWithOutput("add-apt-repository", "deb [arch=arm64] https://download.docker.com/linux/ubuntu "+distribution_cs+" stable")
				}
				// Install Docker
				utils.ExecuteAndWaitWithOutput("apt-get", "update")
				utils.ExecuteAndWaitWithOutput("apt-get", "install", "docker-ce", "docker-ce-cli", "containerd.io", "-y")
			}

			// Install Docker Compose
			utils.ExecuteAndWaitWithOutput("curl", "-L", "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+uname_s+"-"+uname_m, "-o", "/usr/local/bin/docker-compose")
			utils.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
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
		color.Yellow("Congrats! You have installed " + strings.TrimRight(string(docker_version), "\r\n") + " and " + strings.TrimRight(string(compose_version), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
	} else {
		color.Red("An error occured while installing Docker.")
	}
}
