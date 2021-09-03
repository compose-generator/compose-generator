// go:build linux
package pass

import (
	"os"
)

const downloadUrl = "https://get.docker.com"

// InstallDocker installs Docker on the system
func InstallDocker() {
	if IsPrivileged() {
		P("Installing Docker ... ")
		filePath := os.TempDir() + "/install-docker.sh"
		err := DownloadFile(downloadUrl, filePath)
		if err != nil {
			Error("Download of Docker install script failed", err, true)
		}
		ExecuteAndWait("chmod", "+x", filePath)
		ExecuteAndWait("sh", filePath)
		Done()
	}
}
