/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// go:build windows
package pass

import (
	"os"
)

const downloadUrl = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

// InstallDocker installs Docker on the system
func InstallDocker() {
	// Download Docker installer
	spinner := startProcess("Downloading Docker installer ...")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := downloadFile(downloadUrl, filePath)
	if err != nil {
		printError("Download of Docker installer failed", err, true)
	}
	stopProcess(spinner)

	// Run Docker installer
	pl("Running installation ... ")
	pel()
	executeWithOutput(filePath)
	pel()
}
