/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"path/filepath"
	"strings"

	"github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
)

var copyVolumeMockable = copyVolume
var copyBuildDirMockable = copyBuildDir

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateCopyVolumes reads the volume paths from the composition and copies them over to the current work dir
func GenerateCopyVolumes(project *model.CGProject) {
	infoLogger.Println("Copying volumes ...")
	spinner := startProcess("Copying volumes ...")
	for serviceIndex, service := range project.Composition.Services {
		// Copy volumes if existing
		for volumeIndex, volume := range service.Volumes {
			srcPath := filepath.ToSlash(volume.Source)
			if strings.Contains(srcPath, getPredefinedServicesPath()) {
				dstPath := srcPath[len(getPredefinedServicesPath()):]
				dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
				infoLogger.Println("Copying volume from '" + srcPath + "' to '" + dstPath + "'")
				copyVolumeMockable(
					&project.Composition.Services[serviceIndex].Volumes[volumeIndex],
					srcPath,
					dstPath,
				)
			}
		}
		// Copy build dir if existing
		if service.Build != nil {
			srcPath := filepath.ToSlash(service.Build.Context)
			if strings.Contains(srcPath, getPredefinedServicesPath()) {
				dstPath := srcPath[len(getPredefinedServicesPath()):]
				dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
				infoLogger.Println("Copying build dir from '" + srcPath + "' to '" + dstPath + "'")
				copyBuildDirMockable(service.Build, filepath.Clean(srcPath), filepath.Clean(dstPath))
			}
		}
	}
	stopProcess(spinner)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func copyVolume(volume *types.ServiceVolumeConfig, srcPath string, dstPath string) {
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
	// Set the volume bind path to the destination
	volume.Source = dstPath
}

func copyBuildDir(build *types.BuildConfig, srcPath string, dstPath string) {
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
	// Set the volume bind path to the destination
	build.Context = dstPath
}
