// go:build windows
package pass

import (
	"compose-generator/util"
	"os"
	"os/exec"
)

// InstallDocker installs Docker on the system
func InstallDocker() {
	const downloadUrl = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

	// Download Docker installer
	util.P("Downloading Docker installer ... ")
	filePath := os.TempDir() + "/DockerInstaller.exe"
	err := util.DownloadFile(downloadUrl, filePath)
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
