package cmd

import (
	"compose-generator/util"
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	dcu "github.com/compose-generator/dcu"
	model "github.com/compose-generator/dcu/model"
	"github.com/compose-generator/diu"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Add a service to an existing compose file
func Add(flagAdvanced bool, flagRun bool, flagDetached bool, flagForce bool) {
	// Check if CCom is installed
	util.CheckIfCComIsInstalled()

	util.ClearScreen()

	// Check for predefined service templates updates
	util.CheckForServiceTemplateUpdate()

	// Ask for custom YAML file
	path := "./docker-compose.yml"
	if flagAdvanced {
		path = util.TextQuestionWithDefault("Which compose file do you want to add the service to?", "./docker-compose.yml")
	}

	util.P("Parsing compose file ... ")
	// Load compose file
	composeFile, err := dcu.DeserializeFromFile(path)
	if err != nil {
		util.Error("Internal error - unable to load compose file", err, true)
	}
	util.Done()
	util.Pel()

	service, serviceName, existingServiceNames := AddService(composeFile.Services, flagAdvanced, flagForce, false)

	// Ask for services that depend on the new service
	for _, existingServiceName := range askForDependant(existingServiceNames) {
		currentService := composeFile.Services[existingServiceName]
		currentService.DependsOn = append(currentService.DependsOn, serviceName)
		composeFile.Services[existingServiceName] = currentService
	}

	// Add service
	util.P("Adding service ... ")
	composeFile.Services[serviceName] = service
	util.Done()

	// Write to file
	util.P("Saving compose file ... ")
	if dcu.SerializeToFile(composeFile, path) != nil {
		util.Error("Could not write yaml to compose file", err, true)
	}
	util.Done()

	// Run if the corresponding flag is set
	if flagRun || flagDetached {
		util.DockerComposeUp(flagDetached)
	}
}

// AddService asks the user for a new service
func AddService(
	existingServices map[string]model.Service,
	flagAdvanced bool,
	flagForce bool,
	modeGenerate bool,
) (service model.Service, serviceName string, existingServiceNames []string) {
	// Initialize Docker client
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		util.Error("Docker is not installed on your system", err, true)
	}

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
		searchRemoteImage(client, registry, imageName)
	}

	// Ask for service name
	serviceName = askForServiceName(existingServices, imageName)

	// Ask for container name
	containerName := serviceName
	if flagAdvanced {
		containerName = askForContainerName(serviceName)
	}

	// Ask for volumes
	volumes := askForVolumes(client, flagAdvanced)

	// Ask for networks
	networks := askForNetworks(client)

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
		dependsServices = askForDependsOn(util.RemoveStringFromSlice(existingServiceNames, serviceName))
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
	build = util.YesNoQuestion("Build from source?", false)
	if build {
		// Ask for build path
		buildPath = util.TextQuestionWithDefault("Where is your Dockerfile located?", ".")
		// Check if Dockerfile exists
		if !util.FileExists(buildPath+"/Dockerfile") && !util.FileExists(buildPath+"Dockerfile") {
			util.Error("Aborting. The Dockerfile cannot be found.", nil, true)
		}
		return
	}
	// Ask for registry
	registry = util.TextQuestionWithDefault("From which registry do you want to pick?", "docker.io")
	if registry == "docker.io" {
		registry = ""
		return
	}
	registry = registry + "/"
	return
}

func askForImage(build bool) string {
	if build {
		return util.TextQuestion("How do you want to call the built image?")
	}
	return util.TextQuestionWithDefault("From which image do you want to build your service?", "hello-world")
}

func searchRemoteImage(client *client.Client, registry string, image string) {
	util.P("\nSearching image ... ")
	manifest, err := diu.GetImageManifest(registry + image)
	if err == nil {
		color.Green(" found - " + strconv.Itoa(len(manifest.SchemaV2Manifest.Layers)) + " layer(s)\n\n")
		return
	}
	color.Red(" not found or no access\n")
	proceed := util.YesNoQuestion("Proceed anyway?", false)
	if !proceed {
		os.Exit(0)
	}
	util.Pel()
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
	name = util.TextQuestionWithDefault("How do you want to call your service (best practice: lower cased):", defaultName)
	if _, exists := existingServices[name]; exists {
		// Service name already exists
		if !util.YesNoQuestion("This service name alreay exists in the compose file. It will be overwritten if you continue. Continue?", false) {
			os.Exit(0)
		}
	}
	return
}

func askForContainerName(serviceName string) (name string) {
	name = util.TextQuestionWithDefault("How do you want to call your container (best practice: lower cased):", serviceName)
	return
}

func askForVolumes(client *client.Client, flagAdvanced bool) (volumes []string) {
	if util.YesNoQuestion("Do you want to add volumes to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Share another volume?", true) {
			// Ask user for volume attachments
			globalVolume := util.YesNoQuestion("Do you want to add an existing global volume (y) or link a directory / file (n)?", false)
			volumeOuter := ""
			if globalVolume {
				globalVolumes, err := client.VolumeList(context.Background(), filters.Args{})
				if err == nil {
					menuItems := []string{}
					for _, volume := range globalVolumes.Volumes {
						menuItems = append(menuItems, volume.Name+" | Driver: "+volume.Driver)
					}
					if len(globalVolumes.Volumes) >= 1 {
						itemIndex := util.MenuQuestionIndex("Which global volume?", menuItems)
						volumeOuter = globalVolumes.Volumes[itemIndex].Name
					} else if util.YesNoQuestion("No global volumes found. Do you want to create one?", true) {
						volumeOuter = util.TextQuestion("How do you want to call the new global volume?")
						util.ExecuteAndWait("docker", "volume", "create", volumeOuter)
					}
				} else {
					util.Error("Error parsing global volumes.", err, false)
					continue
				}
			} else {
				volumeOuter = util.TextQuestionWithSuggestions("Directory / file on host machine:", func(toComplete string) (files []string) {
					files, _ = filepath.Glob(toComplete + "*")
					return
				})
				volumeOuter = strings.TrimSpace(volumeOuter)
				if !strings.HasPrefix(volumeOuter, "./") && !strings.HasPrefix(volumeOuter, "/") {
					volumeOuter = "./" + volumeOuter
				}
			}

			// Ask for inner path
			volumeInner := util.TextQuestion("Directory / file inside the container:")

			// Ask for volume priviledges if advanced more is enabled
			priviledges := "rw"
			if flagAdvanced {
				result := util.MenuQuestionIndex("Which priviledges does the container has on the volume?", []string{"Read + Write", "Read-only"})
				if result == 1 {
					priviledges = "ro"
				}
			}

			volumes = append(volumes, volumeOuter+":"+volumeInner+":"+priviledges)
		}

		util.Pel()
	}
	return
}

func askForNetworks(client *client.Client) (networks []string) {
	if util.YesNoQuestion("Do you want to add networks to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Assign another network?", true) {
			// Ask user for network assignments
			globalNetwork := util.YesNoQuestion("Do you want to add an external network (y) or create assign a new one (n)?", false)
			networkName := ""
			if globalNetwork {
				globalNetworks, err := client.NetworkList(context.Background(), types.NetworkListOptions{})
				if err == nil {
					menuItems := []string{}
					for _, network := range globalNetworks {
						menuItems = append(menuItems, network.Name+" | Driver: "+network.Driver)
					}
					if len(globalNetworks) >= 1 {
						itemIndex := util.MenuQuestionIndex("Which external network?", menuItems)
						networkName = globalNetworks[itemIndex].Name
					} else if util.YesNoQuestion("No external networks found. Do you want to create one?", true) {
						networkName = util.TextQuestion("How do you want to call the new external network?")
						util.ExecuteAndWait("docker", "network", "create", networkName)
					}
				} else {
					util.Error("Error parsing external networks.", err, false)
					continue
				}
			} else {
				networkName = util.TextQuestion("How do you want to call the new network?")
			}

			networks = append(networks, networkName)
		}
		util.Pel()
	}
	return
}

func askForPorts() (ports []string) {
	if util.YesNoQuestion("Do you want to expose ports of your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Expose another port?", true) {
			portInner := util.TextQuestionWithValidator("Which port do you want to expose? (inner port)", util.PortValidator)
			portOuter := util.TextQuestionWithValidator("To which destination port on the host machine?", util.PortValidator)
			ports = append(ports, portOuter+":"+portInner)
		}
		util.Pel()
	}
	return
}

func askForEnvVariables() (envVariables []string) {
	if util.YesNoQuestion("Do you want to provide environment variables to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Expose another environment variable?", true) {
			variableName := util.TextQuestionWithValidator("Variable name (BEST_PRACTICE_IS_CAPS):", util.EnvVarNameValidator)
			variableValue := util.TextQuestion("Variable value:")
			envVariables = append(envVariables, variableName+"="+variableValue)
		}
		util.Pel()
	}
	return
}

func askForEnvFiles() (envFiles []string) {
	if util.YesNoQuestion("Do you want to provide an environment file to your service?", false) {
		util.Pel()
		for another := true; another; another = util.YesNoQuestion("Add another environment file?", true) {
			// Ask user for env file with auto-suggested text input
			envFile := util.TextQuestionWithDefaultAndSuggestions("Where is your env file located?", "environment.env", func(toComplete string) (files []string) {
				files, _ = filepath.Glob(toComplete + "*.*")
				return
			})
			// Check if the selected file is valid
			if !util.FileExists(envFile) || util.IsDir(envFile) {
				util.Error("File is not valid. Please select another file", nil, false)
				continue
			}
			envFiles = append(envFiles, envFile)
		}
		util.Pel()
	}
	return
}

func askForDependsOn(serviceNames []string) (dependsServices []string) {
	if util.YesNoQuestion("Should your service depend on other services?", false) {
		util.Pel()
		dependsServices = util.MultiSelectMenuQuestion("On which services should your service depend?", serviceNames)
		util.Pel()
	}
	return
}

func askForDependant(serviceNames []string) (dependantServices []string) {
	if util.YesNoQuestion("Should other services depend on your service?", false) {
		util.Pel()
		dependantServices = util.MultiSelectMenuQuestion("Which services should depend on your service?", serviceNames)
		util.Pel()
	}
	return
}

func askForRestart(flagAdvanced bool) (restartValue string) {
	if flagAdvanced {
		util.Pel()
		items := []string{"always", "on-failure", "unless-stopped", "no"}
		restartValue = util.MenuQuestion("When should the service get restarted?", items)
		util.Pel()
	}
	return
}
