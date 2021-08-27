package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"context"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/compose-generator/diu"
	"github.com/fatih/color"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	types_filters "github.com/docker/docker/api/types/filters"
	types_volume "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(
	flagAdvanced bool,
	flagRun bool,
	flagDetached bool,
	flagForce bool,
) {
	// Check if CCom is installed
	util.CheckIfCComIsInstalled()

	// Clear the screen for CG output
	util.ClearScreen()

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Ask for custom YAML file
	composeFilePath := "docker-compose.yml"
	if flagAdvanced {
		composeFilePath = util.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	// Load project
	util.P("Loading project ... ")
	options := project.LoadOptions{ComposeFileName: composeFilePath}
	proj := project.LoadProject(options)
	proj.AdvancedConfig = flagAdvanced
	util.Done()
	util.Pel()

	// Add custom service
	AddCustomService(proj)

	// Save project
	util.P("Saving project ... ")
	project.SaveProject(proj)
	util.Done()
	util.Pel()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	}
}

// AddCustomService adds a fully customizable service to the project
func AddCustomService(project *model.CGProject) {
	newService := spec.ServiceConfig{}

	// Initialize Docker client
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		util.Error("Docker is not installed on your system", err, true)
	}

	// Ask questions
	askBuildFromSource(&newService, project)
	askForServiceName(&newService, project)
	askForContainerName(&newService, project)
	askForVolumes(&newService, project, client)
	askForNetworks(&newService, project, client)
	askForPorts(&newService, project)
	askForEnvVariables(&newService, project)
	askForEnvFiles(&newService, project)
	askForRestart(&newService, project)
	askForDependsOn(&newService, project)
	askForDependant(&newService, project)

	// Add the new service to the project
	project.Project.Services = append(project.Project.Services, newService)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askBuildFromSource(service *spec.ServiceConfig, project *model.CGProject) {
	fromSource := util.YesNoQuestion("Build from source?", false)
	if fromSource { // Build from source
		// Ask for build path
		dockerfilePath := util.TextQuestionWithDefault("Where is your Dockerfile located?", "./Dockerfile")
		// Check if Dockerfile exists
		if !util.FileExists(dockerfilePath) {
			util.Error("The Dockerfile could not be found", nil, true)
		}
		// Add build config to service
		service.Build = &spec.BuildConfig{
			Context:    path.Dir(dockerfilePath),
			Dockerfile: dockerfilePath,
		}
	} else { // Load pre-built image
		registry := ""
		image := ""
		chooseAgain := true
		for chooseAgain {
			// Ask for registry
			registry = util.TextQuestionWithDefault("From which registry do you want to pick?", "docker.io")
			if registry == "docker.io" {
				registry = ""
			} else {
				registry += "/"
			}

			// Ask for image
			image = util.TextQuestionWithDefault("Which Image do you want to use? (e.g. chillibits/ccom:0.8.0)", "hello-world")

			chooseAgain = searchRemoteImage(registry, image)
		}

		options := []string{"frontend", "backend", "database", "db-admin"}
		serviceType := util.MenuQuestion("Which type is the closest match for this service?", options)

		imageBaseName := path.Base(image)
		imageBaseName = strings.Split(imageBaseName, ":")[0]

		// Add image config to service
		service.Image = registry + image
		service.Name = serviceType + "-" + imageBaseName
		service.ContainerName = project.ContainerName + "-" + serviceType + "-" + imageBaseName
	}
}

func askForServiceName(service *spec.ServiceConfig, project *model.CGProject) {
	chooseAgain := true
	for chooseAgain {
		name := util.TextQuestionWithDefault("How do you want to call your service (best practice: lower, kebab cased):", service.Name)
		if util.SliceContainsString(project.Project.ServiceNames(), name) {
			util.Error("This service name already exists. Please choose a different one", nil, false)
		} else {
			chooseAgain = false
		}
	}
}

func askForContainerName(service *spec.ServiceConfig, project *model.CGProject) {
	service.ContainerName = util.TextQuestionWithDefault("How do you want to call your container (best practice: lower, kebab cased):", service.ContainerName)
}

func askForVolumes(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if util.YesNoQuestion("Do you want to add volumes to your service?", false) {
		util.Pel()
		for ok := true; ok; ok = util.YesNoQuestion("Add another volume?", true) {
			globalVolume := util.YesNoQuestion("Do you want to add an existing external volume (y) or link a directory / file (n)?", false)
			if globalVolume {
				askForExternalVolume(service, project, client)
			} else {
				askForFileVolume(service, project)
			}
		}
	}
}

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
		project.Project.Volumes[selectedVolume.Name] = spec.VolumeConfig{
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
		if project.Project.Volumes == nil {
			project.Project.Volumes = make(spec.Volumes)
		}
		project.Project.Volumes[volume.Name] = spec.VolumeConfig{
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

func askForNetworks(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	if util.YesNoQuestion("Do you want to add networks to your service?", false) {
		util.Pel()
		for ok := true; ok; ok = util.YesNoQuestion("Add another network?", true) {
			globalNetwork := util.YesNoQuestion("Do you want to add an external network (y) or create a new one (N)?", false)
			if globalNetwork {
				askForExternalNetwork(service, project, client)
			} else {
				askForNewNetwork(service, project, client)
			}
		}
	}
}

func askForExternalNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Search for external networks
	externalNetworks, err := client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		util.Error("Error parsing external networks", err, false)
		return
	}
	if externalNetworks == nil || len(externalNetworks) == 0 {
		util.Error("There is no external network existing", nil, false)
		return
	}
	// Let the user choose one
	menuItems := []string{}
	for _, network := range externalNetworks {
		menuItems = append(menuItems, network.Name)
	}
	index := util.MenuQuestionIndex("Which one?", menuItems)
	selectedNetwork := externalNetworks[index]

	// Ask for a custom name within the compose file
	customName := util.TextQuestionWithDefault("How do you want to call the network internally?", selectedNetwork.Name)

	// Create maps if not exists
	if service.Networks == nil {
		service.Networks = make(map[string]*spec.ServiceNetworkConfig)
	}
	if project.Project.Networks == nil {
		project.Project.Networks = make(spec.Networks)
	}
	// Add network to the service
	service.Networks[customName] = nil
	// Add network to project-wide network section
	if project.Project.Networks == nil {
		project.Project.Networks = make(map[string]spec.NetworkConfig)
	}
	project.Project.Networks[customName] = spec.NetworkConfig{
		Name: customName,
		External: spec.External{
			Name:     selectedNetwork.Name,
			External: true,
		},
	}
}

func askForNewNetwork(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
	// Ask user to add a new network
	networkName := util.TextQuestion("How do you want to call the new network?")
	external := util.YesNoQuestion("Do you want to create it as an external network and link it in?", false)
	externalConfig := spec.External{}
	if external {
		// Create external network
		_, err := client.NetworkCreate(context.Background(), networkName, types.NetworkCreate{
			Internal: false,
		})
		if err != nil {
			util.Error("External network could not be created", err, false)
			return
		}
		externalConfig = spec.External{
			External: true,
			Name:     networkName,
		}
	}
	// Create maps if not exists
	if service.Networks == nil {
		service.Networks = make(map[string]*spec.ServiceNetworkConfig)
	}
	if project.Project.Networks == nil {
		project.Project.Networks = make(spec.Networks)
	}
	// Add network to the service
	service.Networks[networkName] = &spec.ServiceNetworkConfig{}
	// Add network to project-wide network section
	project.Project.Networks[networkName] = spec.NetworkConfig{
		Name:     networkName,
		External: externalConfig,
	}
}

func askForPorts(service *spec.ServiceConfig, _ *model.CGProject) {
	if util.YesNoQuestion("Do you want to expose ports of your service?", false) {
		util.Pel()
		// Create list if not exists
		if service.Ports == nil {
			service.Ports = []spec.ServicePortConfig{}
		}
		// Question loop
		for another := true; another; another = util.YesNoQuestion("Expose another port?", true) {
			// Ask for inner and outer port
			portInner := util.TextQuestionWithValidator("Which port do you want to expose? (inner port)", util.PortValidator)
			portOuter := util.TextQuestionWithValidator("To which destination port on the host machine?", util.PortValidator)
			portInnerInt, err := strconv.ParseUint(portInner, 10, 32)
			if err != nil {
				util.Error("Port could not be converted to uint32", err, false)
				return
			}
			portOuterInt, err := strconv.ParseUint(portOuter, 10, 32)
			if err != nil {
				util.Error("Port could not be converted to uint32", err, false)
				return
			}

			// Add port to service
			service.Ports = append(service.Ports, spec.ServicePortConfig{
				Mode:      "ingress",
				Protocol:  "tcp",
				Target:    uint32(portInnerInt),
				Published: uint32(portOuterInt),
			})
		}
		util.Pel()
	}
}

func askForEnvVariables(service *spec.ServiceConfig, _ *model.CGProject) {
	if util.YesNoQuestion("Do you want to provide environment variables to your service?", false) {
		util.Pel()
		if service.Environment == nil {
			service.Environment = make(map[string]*string)
		}
		for another := true; another; another = util.YesNoQuestion("Expose another environment variable?", true) {
			// Ask for name and value
			variableName := util.TextQuestionWithValidator("Variable name (BEST_PRACTICE_IS_CAPS):", util.EnvVarNameValidator)
			variableValue := util.TextQuestion("Variable value:")
			// Add env var to service
			service.Environment[variableName] = &variableValue
		}
		util.Pel()
	}
}

func askForEnvFiles(service *spec.ServiceConfig, _ *model.CGProject) {
	if util.YesNoQuestion("Do you want to provide an environment file to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Add another environment file?", true) {
			// Ask for env file with auto-suggested test input
			envFile := util.TextQuestionWithDefaultAndSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
				files, _ = filepath.Glob(toComplete + "*.*")
				return
			})
			// Check if the selected file is valid
			if !util.FileExists(envFile) || util.IsDir(envFile) {
				util.Error("File is not valid. Please select another file", nil, false)
				continue
			}
			// Add env file to service
			service.EnvFile = append(service.EnvFile, envFile)
		}
		util.Pel()
	}
}

func askForRestart(service *spec.ServiceConfig, project *model.CGProject) {
	if project.AdvancedConfig {
		util.Pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		service.Restart = util.MenuQuestion("When should the service get restarted?", items)
		util.Pel()
	}
}

func askForDependsOn(service *spec.ServiceConfig, project *model.CGProject) {
	if util.YesNoQuestion("Do you want your service depend on other services?", false) {
		util.Pel()
		// Ask for services
		selectedServices := util.MultiSelectMenuQuestion("Which ones?", project.Project.ServiceNames())
		// Create map if not exists
		if service.DependsOn == nil {
			service.DependsOn = make(spec.DependsOnConfig)
		}
		// Add service dependencies
		for _, name := range selectedServices {
			service.DependsOn[name] = spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			}
		}
		util.Pel()
	}
}

func askForDependant(service *spec.ServiceConfig, project *model.CGProject) {
	if util.YesNoQuestion("Do you want other services depend on the new one?", false) {
		util.Pel()
		selectedServices := util.MultiSelectMenuQuestion("Which ones?", project.Project.ServiceNames())
		// Add service dependencies
		for _, name := range selectedServices {
			otherService, err := project.Project.GetService(name)
			if err != nil {
				util.Error("Selected service '"+name+"' was not found", err, false)
				continue
			}
			// Create map if not exists
			if otherService.DependsOn == nil {
				service.DependsOn = make(spec.DependsOnConfig)
			}
			// Add service dependency
			otherService.DependsOn[service.Name] = spec.ServiceDependency{
				Condition: spec.ServiceConditionStarted,
			}
		}
		util.Pel()
	}
}

// --------------------------------------------------------------- Helper functions ----------------------------------------------------------------

func searchRemoteImage(registry string, image string) bool {
	util.P("\nSearching image ... ")
	manifest, err := diu.GetImageManifest(registry + image)
	if err != nil {
		color.Red(" not found or no access\n")
		chooseAgain := util.YesNoQuestion("Choose another image (Y) or proceed anyway (n)?", true)
		util.Pel()
		return chooseAgain
	}
	color.Green(" found - " + strconv.Itoa(len(manifest.SchemaV2Manifest.Layers)) + " layer(s)\n\n")
	return false
}
