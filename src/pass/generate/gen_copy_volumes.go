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
	spinner := startProcess("Copying volumes ...")
	for serviceIndex, service := range project.Composition.Services {
		// Copy volumes if existing
		for volumeIndex, volume := range service.Volumes {
			srcPath := filepath.ToSlash(volume.Source)
			if strings.Contains(srcPath, getPredefinedServicesPath()) {
				dstPath := srcPath[len(getPredefinedServicesPath()):]
				dstPath = project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/")
				copyVolumeMockable(
					&project.Composition.Services[serviceIndex].Volumes[volumeIndex],
					srcPath,
					dstPath,
				)
			}
		}
		// Copy build dir if existing
		if service.Build != nil {
			srcPath := filepath.Clean(filepath.ToSlash(service.Build.Context))
			dstPath := srcPath[len(getPredefinedServicesPath()):]
			dstPath = filepath.Clean(project.Composition.WorkingDir + strings.Join(strings.Split(dstPath, "/")[3:], "/"))
			copyBuildDirMockable(service.Build, srcPath, dstPath)
		}
	}
	stopProcess(spinner)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func copyVolume(volume *types.ServiceVolumeConfig, srcPath string, dstPath string) {
	if !fileExists(srcPath) {
		// If srcPath does not exist, simply create a directory at dstPath
		// #nosec G301
		if mkdirAll(dstPath, 0777) != nil {
			printWarning("Could not create volume dir")
		}
	} else {
		// Copy volume
		opt := copy.Options{
			AddPermission: 0777,
		}
		if err := copyFile(srcPath, dstPath, opt); err != nil {
			printWarning("Could not copy volume from '" + srcPath + "' to '" + dstPath + "'")
		}
	}
	// Set the volume bind path to the destination
	volume.Source = dstPath
}

func copyBuildDir(build *types.BuildConfig, srcPath string, dstPath string) {
	if !fileExists(srcPath) {
		// If srcPath does not exist, simply create a directory at dstPath
		// #nosec G301
		if mkdirAll(dstPath, 0777) != nil {
			printWarning("Could not create volume dir")
		}
	} else {
		// Copy volume
		opt := copy.Options{
			AddPermission: 0777,
		}
		if err := copyFile(srcPath, dstPath, opt); err != nil {
			printWarning("Could not copy volume from '" + srcPath + "' to '" + dstPath + "'")
		}
	}
	// Set the volume bind path to the destination
	build.Context = dstPath
}
