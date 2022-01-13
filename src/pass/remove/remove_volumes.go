/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

var isVolumeUsedByOtherServicesMockable = isVolumeUsedByOtherServices

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// RemoveVolumes removes all volumes of a service
func RemoveVolumes(service *spec.ServiceConfig, project *model.CGProject) {
	if len(service.Volumes) > 0 {
		deleteFromDisk := yesNoQuestion("Do you really want to delete all attached volumes of '"+service.Name+"' on disk?", false)
		infoLogger.Println("Removing volumes ...")
		for i := range service.Volumes {
			volume := &service.Volumes[i]
			// Check if volume exists
			if !fileExists(volume.Source) {
				continue // Volume is either a external volume or was already deleted
			}
			// Check if the volume is used by other services
			if !isVolumeUsedByOtherServicesMockable(volume, service, project) {
				// Delete volume recursively
				if deleteFromDisk {
					if err := removeAll(volume.Source); err != nil {
						util.WarningLogger.Println("Could not remove volume at path '" + volume.Source + "': " + err.Error())
						logWarning("Could not remove volume at path '" + volume.Source + "'")
					}
				}
				// Remove in project-wide volumes section
				delete(project.Composition.Volumes, volume.Source)
			}
		}
		infoLogger.Println("Removing volumes (done)")
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func isVolumeUsedByOtherServices(volume *spec.ServiceVolumeConfig, service *spec.ServiceConfig, project *model.CGProject) bool {
	volumeAbs, err := abs(volume.Source)
	if err != nil {
		return false
	}
	for _, otherService := range project.Composition.Services {
		// Skip the service, we are currently editing
		if otherService.Name == service.Name {
			continue
		}
		// Search through volumes of all other services
		for _, otherVolume := range otherService.Volumes {
			otherVolumeAbs, err := abs(otherVolume.Source)
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
