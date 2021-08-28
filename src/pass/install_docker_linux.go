// go:build linux
package pass

import (
	"compose-generator/util"
	"os"
)

// InstallDocker installs Docker on the system
func InstallDocker() {
	if util.IsPrivileged() {
		const downloadUrl = "https://get.docker.com"
		util.P("Installing Docker ... ")
		filePath := os.TempDir() + "/install-docker.sh"
		err := util.DownloadFile(downloadUrl, filePath)
		if err != nil {
			util.Error("Download of Docker install script failed", err, true)
		}
		util.ExecuteAndWait("chmod", "+x", filePath)
		util.ExecuteAndWait("sh", filePath)
		util.Done()
	}
}
