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
		// Get distribution name
		exec.Command("lsb_release -a")
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
