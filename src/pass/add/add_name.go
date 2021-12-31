/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// AddName asks the user if he/she wants to set a name for a service
func AddName(service *spec.ServiceConfig, project *model.CGProject) {
	// Ask for service name
	chooseAgain := true
	for chooseAgain {
		service.Name = textQuestionWithDefault("How do you want to call your service (best practice: lower, kebab cased):", service.Name)
		if util.SliceContainsString(project.Composition.ServiceNames(), service.Name) {
			errorLogger.Println("Service '" + service.Name + "' already exists")
			logError("Service '"+service.Name+"' already exists. Please choose a different one", false)
		} else {
			chooseAgain = false
		}
		infoLogger.Println("New service name: " + service.Name)
	}
	// Set container name
	service.ContainerName = strings.ReplaceAll(strings.ToLower(service.Name), " ", "-")
}
