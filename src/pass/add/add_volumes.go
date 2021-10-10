/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package pass

import (
	"compose-generator/model"
	"context"
	"path/filepath"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
	filtertypes "github.com/docker/docker/api/types/filters"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// CreateDockerVolume calls the Docker client to create a new volume
var CreateDockerVolume = func(client *client.Client, volumeName string) error {
	_, err := client.VolumeCreate(context.Background(), volumetypes.VolumeCreateBody{
		Name: volumeName,
	})
	return err
}

// ListDockerVolumes calls the Docker client to list all available volumes
var ListDockerVolumes = func(client *client.Client) (volumetypes.VolumeListOKBody, error) {
	return client.VolumeList(context.Background(), filtertypes.Args{})
}

var askForExternalVolumeMockable = askForExternalVolume
var askForFileVolumeMockable = askForFileVolume
var askForExistingExternalVolumeMockable = askForExistingExternalVolume
var askForNewExternalVolumeMockable = askForNewExternalVolume

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddVolumes ask the user if he/she wants to add volumes to the configuration
func AddVolumes(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if yesNoQuestion("Do you want to add volumes to your service?", false) {
		pel()
		for ok := true; ok; ok = yesNoQuestion("Add another volume?", true) {
			globalVolume := yesNoQuestion("Do you want to add an existing external volume (y) or link a directory / file (N)?", false)
			if globalVolume {
				askForExternalVolumeMockable(service, project, client)
			} else {
				askForFileVolumeMockable(service, project)
			}
		}
	}
}

// ---------------------------------------------------------------- Private functions --------------------------------------------------------------

func askForExternalVolume(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if yesNoQuestion("Do you want to select an existing one (Y) or do you want to create one (n)?", true) {
		askForExistingExternalVolumeMockable(service, project, client)
	} else {
		askForNewExternalVolumeMockable(service, project, client)
	}
}

func askForExistingExternalVolume(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Search for external volumes
	externalVolumes, err := ListDockerVolumes(client)
	if err != nil {
		printError("Error parsing external volumes", err, false)
		return
	}
	if externalVolumes.Volumes == nil || len(externalVolumes.Volumes) == 0 {
		printError("There is no external volume existing", nil, false)
		return
	}
	// Let the user choose one
	menuItems := []string{}
	for _, volume := range externalVolumes.Volumes {
		menuItems = append(menuItems, volume.Name+" | Driver: "+volume.Driver)
	}
	index := menuQuestionIndex("Which one?", menuItems)
	selectedVolume := externalVolumes.Volumes[index]

	// Ask for inner path
	volumeInner := textQuestion("Directory / file inside the container:")
	volumeInner = filepath.ToSlash(volumeInner)

	// Ask for read-only
	readOnly := false
	if project.AdvancedConfig {
		readOnly = yesNoQuestion("Do you want to make the volume read-only?", false)
	}

	// Add the volume to the service
	service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
		Source:   selectedVolume.Name,
		Target:   volumeInner,
		Type:     spec.VolumeTypeVolume,
		ReadOnly: readOnly,
	})
	// Add the volume to the project-wide volume section
	if project.Composition == nil {
		project.Composition = &spec.Project{}
	}
	if project.Composition.Volumes == nil {
		project.Composition.Volumes = spec.Volumes{}
	}
	project.Composition.Volumes[selectedVolume.Name] = spec.VolumeConfig{
		Name: selectedVolume.Name,
		External: spec.External{
			Name:     selectedVolume.Name,
			External: true,
		},
	}
}

func askForNewExternalVolume(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Ask user for volume name
	name := textQuestion("How do you want to call your external volume?")
	// Add external volume
	err := CreateDockerVolume(client, name)
	if err != nil {
		printError("Could not create external volume", err, false)
		return
	}

	// Ask for inner path
	volumeInner := textQuestion("Directory / file inside the container:")

	// Ask for read-only
	readOnly := false
	if project.AdvancedConfig {
		readOnly = yesNoQuestion("Do you want to make the volume read-only?", false)
	}

	// Add the volume to the service
	service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
		Source:   name,
		Target:   volumeInner,
		Type:     spec.VolumeTypeVolume,
		ReadOnly: readOnly,
	})
	// Add the volume to the project-wide volume section
	if project.Composition == nil {
		project.Composition = &spec.Project{}
	}
	if project.Composition.Volumes == nil {
		project.Composition.Volumes = spec.Volumes{}
	}
	project.Composition.Volumes[name] = spec.VolumeConfig{
		Name: name,
		External: spec.External{
			Name:     name,
			External: true,
		},
	}
}

func askForFileVolume(service *spec.ServiceConfig, project *model.CGProject) {
	// Ask for outer path
	volumeOuter := textQuestionWithSuggestions("Directory / file on host machine:", func(toComplete string) []string {
		files, _ := filepath.Glob(toComplete + "*")
		return files
	})
	volumeOuter = strings.TrimSpace(volumeOuter)
	if !strings.HasPrefix(volumeOuter, "./") && !strings.HasPrefix(volumeOuter, "/") {
		volumeOuter = "./" + volumeOuter
	}

	// Ask for inner path
	volumeInner := textQuestion("Directory / file inside the container:")
	volumeInner = strings.TrimSpace(volumeInner)

	// Ask for read-only
	readOnly := false
	if project.AdvancedConfig {
		readOnly = yesNoQuestion("Do you want to make the volume read-only?", false)
	}

	service.Volumes = append(service.Volumes, spec.ServiceVolumeConfig{
		Type:     spec.VolumeTypeBind,
		Source:   volumeOuter,
		Target:   volumeInner,
		ReadOnly: readOnly,
	})
}
