package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"os"
	"path/filepath"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// RemoveVolumes removes all volumes of a service
func RemoveVolumes(service *spec.ServiceConfig, project *model.CGProject) {
	if util.YesNoQuestion("Do you really want to delete all attached volumes of '"+service.Name+"' on disk?", false) {
		// Delete volumes on disk
		for _, volume := range service.Volumes {
			// Check if volume exists
			if !util.FileExists(volume.Source) {
				continue // Volume is either a external volume or was already deleted
			}
			// Check if the volume is used by other services
			if !isVolumeUsedByOtherServices(&volume, project) {
				// Delete volume recursively
				os.RemoveAll(volume.Source)
			}

		}
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func isVolumeUsedByOtherServices(volume *spec.ServiceVolumeConfig, project *model.CGProject) bool {
	volumeAbs, err := filepath.Abs(volume.Source)
	if err != nil {
		return false
	}
	for _, otherService := range project.Project.Services {
		for _, otherVolume := range otherService.Volumes {
			otherVolumeAbs, err := filepath.Abs(otherVolume.Source)
			if err != nil {
				continue
			}
			if strings.HasPrefix(otherVolumeAbs, volumeAbs) {
				// Another service binds the same directory or a sub-directory of it => not delete
				return true
			}
		}
	}
	return false
}
