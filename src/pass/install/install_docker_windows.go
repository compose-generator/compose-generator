/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

// go:build windows

package pass

import (
	"os"
)

const downloadUrl = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

// InstallDocker installs Docker on the system
func InstallDocker() {
	infoLogger.Println("Executing Install command")
	// Download Docker installer
	infoLogger.Println("Downloading Docker installer ...")
	spinner := startProcess("Downloading Docker installer ...")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := downloadFile(downloadUrl, filePath)
	if err != nil {
		errorLogger.Println("Download of Docker installer failed: " + err.Error())
		logError("Download of Docker installer failed", true)
	}
	stopProcess(spinner)
	infoLogger.Println("Downloading Docker installer (done)")

	// Run Docker installer
	pl("Running installation ... ")
	pel()
	executeWithOutput(filePath)
	pel()
}
