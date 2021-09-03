// go:build windows
package pass

import (
	"os"
)

const downloadUrl = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

// InstallDocker installs Docker on the system
func InstallDocker() {
	// Download Docker installer
	P("Downloading Docker installer ... ")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := DownloadFile(downloadUrl, filePath)
	if err != nil {
		Error("Download of Docker installer failed", err, true)
	}
	Done()

	// Run Docker installer
	Pl("Running installation ... ")
	Pel()
	ExecuteWithOutput(filePath)
	Pel()
}
