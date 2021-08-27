package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"path"
	"strconv"
	"strings"

	"github.com/compose-generator/diu"
	"github.com/fatih/color"

	spec "github.com/compose-spec/compose-go/types"
)

// AddBuildOrImage asks the user if he/she wants to build from source or add a predefined image to a service
func AddBuildOrImage(service *spec.ServiceConfig, project *model.CGProject) {
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
