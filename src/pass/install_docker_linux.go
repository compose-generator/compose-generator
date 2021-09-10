// go:build linux
package pass

import (
	"os"
)

const downloadUrl = "https://get.docker.com"

// InstallDocker installs Docker on the system
func InstallDocker() {
	if isPrivileged() {
		spinner := startProcess("Installing Docker ...")
		filePath := os.TempDir() + "/install-docker.sh"
		err := downloadFile(downloadUrl, filePath)
		if err != nil {
			printError("Download of Docker install script failed", err, true)
		}
		executeAndWait("chmod", "+x", filePath)
		executeAndWait("sh", filePath)
		stopProcess(spinner)
	}
}
