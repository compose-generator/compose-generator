package util

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// CheckForServiceTemplateUpdate checks if any updates are available for the predefined service templates
func CheckForServiceTemplateUpdate(version string) {
	// Create predefined templates dir if not exitsts
	predefinedTemplatesDir := GetPredefinedServicesPath()
	if !FileExists(predefinedTemplatesDir) {
		if err := os.MkdirAll(predefinedTemplatesDir, 0777); err != nil {
			Error("Could not create directory for predefined templates", err, true)
		}
	}

	fileUrl := "https://github.com/compose-generator/compose-generator/releases/download/" + version + "/predefined-services.tar.gz"
	outputPath := GetPredefinedServicesPath() + "/predefined-services.tar.gz"
	shouldUpdate := false

	if FileExists(outputPath) { // File exists => version check
		file, err := os.Stat(outputPath)
		if err != nil {
			Error("Could not access existing template archive", err, true)
		}
		lastModifiedLocal := file.ModTime().Unix()

		// Issue HEAD request for services archive
		res, err := http.Head(fileUrl)
		if err != nil {
			Warning("Could not check for template updates")
			return
		}
		lastModified := res.Header["Last-Modified"][0]
		t, err := time.Parse(time.RFC1123, lastModified)
		if err != nil {
			Error("Cannot parse last modified of remote file", err, true)
		}
		if t.Unix() > lastModifiedLocal {
			shouldUpdate = true
		}
	} else { // File does not exist => download directly
		shouldUpdate = true
	}

	// Download update if necessary
	if shouldUpdate {
		P("Downloading predefined services update ... ")
		DownloadFile(fileUrl, outputPath)
		filepath, err := filepath.Abs(predefinedTemplatesDir)
		if err != nil {
			Error("Could not build path", err, true)
		}
		ExecuteOnLinuxWithCustomVolume("tar xfvz predefined-services.tar.gz", filepath)
		Done()
	}
}
