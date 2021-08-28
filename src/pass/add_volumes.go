package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"context"
	"path/filepath"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
	types_filters "github.com/docker/docker/api/types/filters"
	types_volume "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddVolumes ask the user if he/she wants to add volumes to the configuration
func AddVolumes(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if util.YesNoQuestion("Do you want to add volumes to your service?", false) {
		util.Pel()
		for ok := true; ok; ok = util.YesNoQuestion("Add another volume?", true) {
			globalVolume := util.YesNoQuestion("Do you want to add an existing external volume (y) or link a directory / file (N)?", false)
			if globalVolume {
				askForExternalVolume(service, project, client)
			} else {
				askForFileVolume(service, project)
			}
		}
	}
}

// ---------------------------------------------------------------- Private functions --------------------------------------------------------------

func askForExternalVolume(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if util.YesNoQuestion("Do you want to select an existing one (Y) or do you want to create one (n)?", true) {
		// Search for external volumes
		externalVolumes, err := client.VolumeList(context.Background(), types_filters.Args{})
		if err != nil {
			util.Error("Error parsing external volumes", err, false)
			return
		}
		if externalVolumes.Volumes == nil || len(externalVolumes.Volumes) == 0 {
			util.Error("There is no external volume existing", nil, false)
			return
		}
		// Let the user choose one
		menuItems := []string{}
		for _, volume := range externalVolumes.Volumes {
			menuItems = append(menuItems, volume.Name+" | Driver: "+volume.Driver)
		}
		index := util.MenuQuestionIndex("Which one?", menuItems)
		selectedVolume := externalVolumes.Volumes[index]

		// Ask for inner path
		volumeInner := util.TextQuestion("Directory / file inside the container:")
		volumeInner = strings.ReplaceAll(volumeInner, "\\", "/")

		// Ask for read-only
		readOnly := false
		if project.AdvancedConfig {
			readOnly = util.YesNoQuestion("Do you want to make the volume read-only?", false)
		}

		// Add the volume to the service
		service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
			Source:   selectedVolume.Name,
			Target:   volumeInner,
			Type:     spec.VolumeTypeVolume,
			ReadOnly: readOnly,
		})
		// Add the volume to the project-wide volume section
		project.Composition.Volumes[selectedVolume.Name] = spec.VolumeConfig{
			Name: selectedVolume.Name,
			External: spec.External{
				Name:     selectedVolume.Name,
				External: true,
			},
		}
	} else {
		// Ask user for volume name
		name := util.TextQuestion("How do you want to call your external volume?")
		// Add external volume
		volume, err := client.VolumeCreate(context.Background(), types_volume.VolumeCreateBody{
			Name: name,
		})
		if err != nil {
			util.Error("Could not create external volume", err, false)
			return
		}

		// Ask for inner path
		volumeInner := util.TextQuestion("Directory / file inside the container:")

		// Ask for read-only
		readOnly := false
		if project.AdvancedConfig {
			readOnly = util.YesNoQuestion("Do you want to make the volume read-only?", false)
		}

		// Add the volume to the service
		service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
			Source:   volume.Name,
			Target:   volumeInner,
			Type:     spec.VolumeTypeVolume,
			ReadOnly: readOnly,
		})
		// Add the volume to the project-wide volume section
		if project.Composition.Volumes == nil {
			project.Composition.Volumes = make(spec.Volumes)
		}
		project.Composition.Volumes[volume.Name] = spec.VolumeConfig{
			Name: volume.Name,
			External: spec.External{
				Name:     volume.Name,
				External: true,
			},
		}
	}
}

func askForFileVolume(service *spec.ServiceConfig, project *model.CGProject) {
	// Ask for outer path
	volumeOuter := util.TextQuestionWithSuggestions("Directory / file on host machine:", func(toComplete string) (files []string) {
		files, _ = filepath.Glob(toComplete + "*")
		return
	})
	volumeOuter = strings.TrimSpace(volumeOuter)
	if !strings.HasPrefix(volumeOuter, "./") && !strings.HasPrefix(volumeOuter, "/") {
		volumeOuter = "./" + volumeOuter
	}

	// Ask for inner path
	volumeInner := util.TextQuestion("Directory / file inside the container:")

	// Ask for read-only
	readOnly := false
	if project.AdvancedConfig {
		readOnly = util.YesNoQuestion("Do you want to make the volume read-only?", false)
	}

	service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
		Type:     spec.VolumeTypeBind,
		Source:   volumeOuter,
		Target:   volumeInner,
		ReadOnly: readOnly,
	})
}
