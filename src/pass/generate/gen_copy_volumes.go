/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

var copyVolumeMockable = copyVolume

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateCopyVolumes reads the volume paths from the composition and copies them over to the current work dir
func GenerateCopyVolumes(project *model.CGProject, selected *model.SelectedTemplates) {
	infoLogger.Println("Copying volumes ...")
	spinner := startProcess("Copying volumes ...")
	// Copy volumes and build contexts
	for _, template := range selected.GetAll() {
		for _, volume := range template.Volumes {
			srcPath := filepath.Clean(template.Dir + "/" + volume.DefaultValue)
			dstPath := filepath.Clean(project.Composition.WorkingDir + project.Vars[volume.Variable])
			infoLogger.Println("Copying volume from '" + srcPath + "' to '" + dstPath + "'")
			copyVolumeMockable(srcPath, dstPath)
		}
	}
	// Change volume paths in composition to the new ones
	for serviceIndex, service := range project.Composition.Services {
		for volumeIndex, volume := range service.Volumes {
			if strings.Contains(volume.Source, getPredefinedServicesPath()) {
				dstPath := volume.Source[len(getPredefinedServicesPath()):]
				dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
				project.Composition.Services[serviceIndex].Volumes[volumeIndex].Source = dstPath
			}
		}
	}
	stopProcess(spinner)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func copyVolume(srcPath string, dstPath string) {
	if !fileExists(srcPath) {
		// If srcPath does not exist, simply create a directory at dstPath
		// #nosec G301
		if err := mkdirAll(dstPath, 0777); err != nil {
			warningLogger.Println("Could not create volume dir: " + err.Error())
			logWarning("Could not create volume dir")
		}
	} else {
		// Copy volume
		opt := copy.Options{
			AddPermission: 0777,
		}
		if err := copyFile(srcPath, dstPath, opt); err != nil {
			warningLogger.Println("Could not copy volume from '" + srcPath + "' to '" + dstPath + "': " + err.Error())
			logWarning("Could not copy volume from '" + srcPath + "' to '" + dstPath + "'")
		}
	}
}
