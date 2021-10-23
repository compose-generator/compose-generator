/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
)

// GenerateAddCustomService adds a custom service configuration to the composition
func GenerateAddCustomService(project *model.CGProject, serviceType string) {
	newService := spec.ServiceConfig{}

	// Initialize Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		errorLogger.Println("Could not intanciate Docker client: " + err.Error())
		logError("Could not intanciate Docker client. Please check your Docker installation", true)
		return
	}

	// Execute passes on the service
	addBuildOrImagePass(&newService, project, serviceType)
	addNamePass(&newService, project)
	addContainerNamePass(&newService, project)
	addVolumesPass(&newService, project, client)
	addNetworksPass(&newService, project, client)
	addPortsPass(&newService, project)
	addEnvVarsPass(&newService, project)
	addEnvFilesPass(&newService, project)
	addRestartPass(&newService, project)
	addDependsPass(&newService, project)
	addDependantsPass(&newService, project)

	// Add the new service to the project
	project.Composition.Services = append(project.Composition.Services, newService)

	pel()
}
