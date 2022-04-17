/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

// go:build darwin

package pass

import "os"

const downloadURL = "https://get.docker.com"

// InstallDocker installs Docker on the system
func InstallDocker() {
	infoLogger.Println("Executing Install command")
	if isPrivileged() {
		infoLogger.Println("Installing Docker ...")
		spinner := startProcess("Installing Docker ...")
		filePath := os.TempDir() + "/install-docker.sh"
		err := downloadFile(downloadURL, filePath)
		if err != nil {
			errorLogger.Println("Download of Docker install script failed: " + err.Error())
			logError("Download of Docker install script failed", true)
		}
		executeAndWait("chmod", "+x", filePath)
		executeAndWait("sh", filePath)
		stopProcess(spinner)
		infoLogger.Println("Installing Docker (done)")
	}
}
