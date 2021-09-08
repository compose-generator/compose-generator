package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// GenerateCopyVolumes reads the volume paths from the composition and copies them over to the current work dir
func GenerateCopyVolumes(project *model.CGProject) {
	spinner := startProcess("Copying volumes ...")
	for serviceIndex, service := range project.Composition.Services {
		for volumeIndex, volume := range service.Volumes {
			srcPath := filepath.ToSlash(volume.Source)
			dstPath := srcPath[len(util.GetPredefinedServicesPath()):]
			dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
			copyVolume(&project.Composition.Services[serviceIndex].Volumes[volumeIndex], srcPath, dstPath)
		}
		if service.Build != nil {
			srcPath := filepath.ToSlash(service.Build.Context)
			dstPath := srcPath[len(util.GetPredefinedServicesPath()):]
			dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
			copyBuildDir(service.Build, srcPath, dstPath)
		}
	}
	stopProcess(spinner)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func copyVolume(volume *types.ServiceVolumeConfig, sourcePath string, dstPath string) {
	if !fileExists(sourcePath) {
		// If srcPath does not exist, simply create a directory at dstPath
		if os.MkdirAll(dstPath, 0755) != nil {
			util.Warning("Could not create volume dir")
		}
	} else {
		// Copy volume
		copy.Copy(sourcePath, dstPath)
	}
	// Set the volume bind path to the destination
	volume.Source = dstPath
}

func copyBuildDir(build *types.BuildConfig, sourcePath string, dstPath string) {
	if !fileExists(sourcePath) {
		// If srcPath does not exist, simply create a directory at dstPath
		if os.MkdirAll(dstPath, 0755) != nil {
			util.Warning("Could not create volume dir")
		}
	} else {
		// Copy volume
		copy.Copy(sourcePath, dstPath)
	}
	// Set the volume bind path to the destination
	build.Context = dstPath
}
