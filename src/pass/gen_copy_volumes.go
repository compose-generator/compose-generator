package pass

import (
	"compose-generator/model"
	"path/filepath"
	"strings"

	"github.com/compose-spec/compose-go/types"
)

var copyVolumeMockable = copyVolume
var copyBuildDirMockable = copyBuildDir

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateCopyVolumes reads the volume paths from the composition and copies them over to the current work dir
func GenerateCopyVolumes(project *model.CGProject) {
	spinner := startProcess("Copying volumes ...")
	for serviceIndex, service := range project.Composition.Services {
		// Copy volumes if existing
		for volumeIndex, volume := range service.Volumes {
			srcPath := filepath.ToSlash(volume.Source)
			dstPath := srcPath[len(getPredefinedServicesPath()):]
			dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
			copyVolumeMockable(
				&project.Composition.Services[serviceIndex].Volumes[volumeIndex],
				srcPath,
				dstPath,
			)
		}
		// Copy build dir if existing
		if service.Build != nil {
			srcPath := filepath.ToSlash(service.Build.Context)
			dstPath := srcPath[len(getPredefinedServicesPath()):]
			dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
			copyBuildDirMockable(service.Build, srcPath, dstPath)
		}
	}
	stopProcess(spinner)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func copyVolume(volume *types.ServiceVolumeConfig, srcPath string, dstPath string) {
	if !fileExists(srcPath) {
		// If srcPath does not exist, simply create a directory at dstPath
		if mkdirAll(dstPath, 0755) != nil {
			printWarning("Could not create volume dir")
		}
	} else {
		// Copy volume
		copyFile(srcPath, dstPath)
	}
	// Set the volume bind path to the destination
	volume.Source = dstPath
}

func copyBuildDir(build *types.BuildConfig, srcPath string, dstPath string) {
	if !fileExists(srcPath) {
		// If srcPath does not exist, simply create a directory at dstPath
		if mkdirAll(dstPath, 0755) != nil {
			printWarning("Could not create volume dir")
		}
	} else {
		// Copy volume
		copyFile(srcPath, dstPath)
	}
	// Set the volume bind path to the destination
	build.Context = dstPath
}
