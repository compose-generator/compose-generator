package cmd

import (
	"compose-generator/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	dcu "github.com/compose-generator/dcu"
	model "github.com/compose-generator/dcu/model"
	"github.com/compose-generator/diu"
	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool) {
	utils.ClearScreen()

	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flagAdvanced {
		path = utils.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	utils.P("Parsing compose file ... ")
	// Load compose file
	composeFile, err := dcu.DeserializeFromFile(path)
	if err != nil {
		utils.Error("Internal error - unable to load compose file", err, true)
	}
	utils.Done()
	utils.Pel()

	service, serviceName, existingServiceNames := AddService(composeFile.Services, flagAdvanced, flagForce, false)

	// Ask for services that depend on the new service
	for _, existingServiceName := range askForDependant(existingServiceNames) {
		currentService := composeFile.Services[existingServiceName]
		currentService.DependsOn = append(currentService.DependsOn, serviceName)
		composeFile.Services[existingServiceName] = currentService
	}

	// Add service
	utils.P("Adding service ... ")
	composeFile.Services[serviceName] = service
	utils.Done()

	// Write to file
	utils.P("Saving compose file ... ")
	if dcu.SerializeToFile(composeFile, path) != nil {
		utils.Error("Could not write yaml to compose file", err, true)
	}
	utils.Done()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		utils.DockerComposeUp(flagDetached)
	}
}

// AddService asks the user for a new service
func AddService(existingServices map[string]model.Service, flagAdvanced bool, flagForce bool, modeGenerate bool) (service model.Service, serviceName string, existingServiceNames []string) {
	// Get names of existing services
	for name := range existingServices {
		existingServiceNames = append(existingServiceNames, name)
	}

	// Ask if the image should be built from source
	build, buildPath, registry := askBuildFromSource()

	// Ask for image
	imageName := askForImage(build)

	// Search for remote image and check manifest
	if !build && !flagForce {
		searchRemoteImage(registry, imageName)
	}

	// Ask for service name
	serviceName = askForServiceName(existingServices, imageName)

	// Ask for container name
	containerName := serviceName
	if flagAdvanced {
		containerName = askForContainerName(serviceName)
	}

	// Ask for volumes
	volumes := askForVolumes(flagAdvanced)

	// Ask for networks
	networks := askForNetworks()

	// Ask for ports
	ports := askForPorts()

	// Ask for env files
	envFiles := askForEnvFiles()

	// Ask for env variables
	envVariables := []string{}
	if len(envFiles) == 0 {
		envVariables = askForEnvVariables()
	}

	// Ask for services, the new one should depend on
	var dependsServices []string
	if !modeGenerate {
		dependsServices = askForDependsOn(utils.RemoveStringFromSlice(existingServiceNames, serviceName))
	}

	// Ask for restart mode
	restartValue := askForRestart(flagAdvanced)

	// Build service object
	service = model.Service{
		Build:         buildPath,
		Image:         registry + imageName,
		ContainerName: containerName,
		Volumes:       volumes,
		Networks:      networks,
		Ports:         ports,
		Restart:       restartValue,
		DependsOn:     dependsServices,
		EnvFile:       envFiles,
		Environment:   envVariables,
	}
	return
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askBuildFromSource() (build bool, buildPath string, registry string) {
	build = utils.YesNoQuestion("Build from source?", false)
	if build {
		// Ask for build path
		buildPath = utils.TextQuestionWithDefault("Where is your Dockerfile located?", ".")
		// Check if Dockerfile exists
		if !utils.FileExists(buildPath+"/Dockerfile") && !utils.FileExists(buildPath+"Dockerfile") {
			utils.Error("Aborting. The Dockerfile cannot be found.", nil, true)
		}
	} else {
		// Ask for registry
		registry = utils.TextQuestionWithDefault("From which registry do you want to pick?", "docker.io")
		if registry == "docker.io" {
			registry = ""
		} else {
			registry = registry + "/"
		}
	}
	return
}

func askForImage(build bool) string {
	if build {
		return utils.TextQuestion("How do you want to call the built image?")
	}
	return utils.TextQuestionWithDefault("From which image do you want to build your service?", "hello-world")
}

func searchRemoteImage(registry string, image string) {
	utils.P("\nSearching image ... ")
	manifest, err := diu.GetImageManifest(registry + image)
	if err == nil {
		color.Green(" found - " + strconv.Itoa(len(manifest.SchemaV2Manifest.Layers)) + " layer(s)\n\n")
	} else {
		color.Red(" not found or no access\n")
		proceed := utils.YesNoQuestion("Proceed anyway?", false)
		if !proceed {
			os.Exit(0)
		}
		utils.Pel()
	}
}

func askForServiceName(existingServices map[string]model.Service, imageName string) (name string) {
	// Set image name as default service name
	defaultName := imageName
	i := strings.Index(defaultName, "/")
	if i > -1 {
		defaultName = defaultName[i+1:]
	}
	i = strings.Index(defaultName, ":")
	if i > -1 {
		defaultName = defaultName[:i]
	}

	// Ask for the service name
	name = utils.TextQuestionWithDefault("How do you want to call your service (best practise: lower cased):", defaultName)
	if _, exists := existingServices[name]; exists {
		// Service name already exists
		if !utils.YesNoQuestion("This service name alreay exists in the compose file. It will be overwritten if you continue. Continue?", false) {
			os.Exit(0)
		}
	}
	return
}

func askForContainerName(serviceName string) (name string) {
	name = utils.TextQuestionWithDefault("How do you want to call your container (best practise: lower cased):", serviceName)
	return
}

func askForVolumes(flagAdvanced bool) (volumes []string) {
	if utils.YesNoQuestion("Do you want to add volumes to your service?", false) {
		utils.Pel()
		for another := true; another; another = utils.YesNoQuestion("Share another volume?", true) {
			// Ask user for volume attachments
			globalVolume := utils.YesNoQuestion("Do you want to add an existing global volume (y) or link a directory / file (n)?", false)
			volumeOuter := ""
			if globalVolume {
				globalVolumes, err := diu.GetExistingVolumes()
				if err == nil {
					menuItems := []string{}
					for _, volume := range globalVolumes {
						menuItems = append(menuItems, volume.Name+" | Driver: "+volume.Driver)
					}
					if len(globalVolumes) >= 1 {
						itemIndex := utils.MenuQuestionIndex("Which global volume?", menuItems)
						volumeOuter = globalVolumes[itemIndex].Name
					} else if utils.YesNoQuestion("No global volumes found. Do you want to create one?", true) {
						volumeOuter = utils.TextQuestion("How do you want to call the new global volume?")
						utils.ExecuteAndWait("docker", "volume", "create", volumeOuter)
					}
				} else {
					utils.Error("Error parsing global volumes.", err, false)
					continue
				}
			} else {
				volumeOuter = utils.TextQuestionWithSuggestions("Directory / file on host machine:", func(toComplete string) (files []string) {
					files, _ = filepath.Glob(toComplete + "*")
					return
				})
				volumeOuter = strings.TrimSpace(volumeOuter)
				if !strings.HasPrefix(volumeOuter, "./") && !strings.HasPrefix(volumeOuter, "/") {
					volumeOuter = "./" + volumeOuter
				}
			}

			// Ask for inner path
			volumeInner := utils.TextQuestion("Directory / file inside the container:")

			// Ask for volume priviledges if advanced more is enabled
			priviledges := "rw"
			if flagAdvanced {
				result := utils.MenuQuestionIndex("Which priviledges does the container has on the volume?", []string{"Read + Write", "Read-only"})
				if result == 1 {
					priviledges = "ro"
				}
			}

			volumes = append(volumes, volumeOuter+":"+volumeInner+":"+priviledges)
		}

		utils.Pel()
	}
	return
}

func askForNetworks() (networks []string) {
	if utils.YesNoQuestion("Do you want to add networks to your service?", false) {
		utils.Pel()
		for another := true; another; another = utils.YesNoQuestion("Assign another network?", true) {
			// Ask user for network assignments
			globalNetwork := utils.YesNoQuestion("Do you want to add an external network (y) or create assign a new one (n)?", false)
			networkName := ""
			if globalNetwork {
				globalNetworks, err := diu.GetExistingNetworks()
				if err == nil {
					menuItems := []string{}
					for _, network := range globalNetworks {
						menuItems = append(menuItems, network.Name+" | Driver: "+network.Driver)
					}
					if len(globalNetworks) >= 1 {
						itemIndex := utils.MenuQuestionIndex("Which external network?", menuItems)
						networkName = globalNetworks[itemIndex].Name
					} else if utils.YesNoQuestion("No external networks found. Do you want to create one?", true) {
						networkName = utils.TextQuestion("How do you want to call the new external network?")
						utils.ExecuteAndWait("docker", "network", "create", networkName)
					}
				} else {
					utils.Error("Error parsing external networks.", err, false)
					continue
				}
			} else {
				networkName = utils.TextQuestion("How do you want to call the new network?")
			}

			networks = append(networks, networkName)
		}
		utils.Pel()
	}
	return
}

func askForPorts() (ports []string) {
	if utils.YesNoQuestion("Do you want to expose ports of your service?", false) {
		utils.Pel()
		for another := true; another; another = utils.YesNoQuestion("Expose another port?", true) {
			portInner := utils.TextQuestionWithValidator("Which port do you want to expose? (inner port)", utils.PortValidator)
			portOuter := utils.TextQuestionWithValidator("To which destination port on the host machine?", utils.PortValidator)
			ports = append(ports, portOuter+":"+portInner)
		}
		utils.Pel()
	}
	return
}

func askForEnvVariables() (envVariables []string) {
	if utils.YesNoQuestion("Do you want to provide environment variables to your service?", false) {
		utils.Pel()
		for another := true; another; another = utils.YesNoQuestion("Expose another environment variable?", true) {
			variableName := utils.TextQuestionWithValidator("Variable name (BEST_PRACTISE_IS_CAPS):", utils.EnvVarNameValidator)
			variableValue := utils.TextQuestion("Variable value:")
			envVariables = append(envVariables, variableName+"="+variableValue)
		}
		utils.Pel()
	}
	return
}

func askForEnvFiles() (envFiles []string) {
	if utils.YesNoQuestion("Do you want to provide an environment file to your service?", false) {
		utils.Pel()
		for another := true; another; another = utils.YesNoQuestion("Add another environment file?", true) {
			// Ask user for env file with auto-suggested text input
			envFile := utils.TextQuestionWithDefaultAndSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
				files, _ = filepath.Glob(toComplete + "*.*")
				return
			})
			// Check if the selected file is valid
			if !utils.FileExists(envFile) || utils.IsDirectory(envFile) {
				utils.Error("File is not valid. Please select another file", nil, false)
				continue
			}
			envFiles = append(envFiles, envFile)
		}
		utils.Pel()
	}
	return
}

func askForDependsOn(serviceNames []string) (dependsServices []string) {
	if utils.YesNoQuestion("Should your service depend on other services?", false) {
		utils.Pel()
		dependsServices = utils.MultiSelectMenuQuestion("On which services should your service depend?", serviceNames)
		utils.Pel()
	}
	return
}

func askForDependant(serviceNames []string) (dependantServices []string) {
	if utils.YesNoQuestion("Should other services depend on your service?", false) {
		utils.Pel()
		dependantServices = utils.MultiSelectMenuQuestion("Which services should depend on your service?", serviceNames)
		utils.Pel()
	}
	return
}

func askForRestart(flagAdvanced bool) (restartValue string) {
	if flagAdvanced {
		utils.Pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		restartValue = utils.MenuQuestion("When should the service get restarted?", items)
		utils.Pel()
	}
	return
}
