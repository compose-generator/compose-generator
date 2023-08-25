/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"path"
	"strconv"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddBuildOrImage asks the user if he/she wants to build from source or add a predefined image to a service
func AddBuildOrImage(service *spec.ServiceConfig, project *model.CGProject, serviceType string) {
	fromSource := yesNoQuestion("Build from source?", false)
	if fromSource { // Build from source
		infoLogger.Println("User chose a source build")
		// Ask for build path
		dockerfilePath := textQuestionWithDefault("Where is your Dockerfile located?", "./Dockerfile")
		// Check if Dockerfile exists
		if !fileExists(dockerfilePath) {
			errorLogger.Println("Dockerfile could not be found")
			logError("Dockerfile could not be found", true)
			return
		}
		// Add build config to service
		service.Build = &spec.BuildConfig{
			Context:    path.Dir(dockerfilePath),
			Dockerfile: dockerfilePath,
		}
	} else { // Load pre-built image
		infoLogger.Println("User chose a pre-built image")
		registry := ""
		image := ""
		chooseAgain := true
		for chooseAgain {
			// Ask for registry
			registry = textQuestionWithDefault("From which registry do you want to pick?", "docker.io")
			if registry == "docker.io" {
				registry = ""
			} else {
				registry += "/"
			}

			// Ask for image
			image = textQuestionWithDefault("Which Image do you want to use? (e.g. chillibits/ccom:0.8.0)", "hello-world")

			chooseAgain = searchRemoteImage(registry, image)
		}

		imageBaseName := path.Base(image)
		imageBaseName = strings.Split(imageBaseName, ":")[0]

		// Add image config to service
		service.Image = registry + image
		service.Name = serviceType + "-" + imageBaseName
	}
}

// ---------------------------------------------------------------- Private functions ---------------------------------------------------------------

func searchRemoteImage(registry string, image string) bool {
	pel()
	p("Searching image ... ")
	manifest, err := getImageManifest(registry + image)
	if err != nil {
		errorLogger.Println("Image '" + registry + image + "' not found or no access")
		logError(" not found or no access", false)
		chooseAgain := yesNoQuestion("Choose another image (Y) or proceed anyway (n)?", true)
		util.Pel()
		return chooseAgain
	}
	success(" found - " + strconv.Itoa(len(manifest.SchemaV2Manifest.Layers)) + " layer(s)")
	pel()
	return false
}
