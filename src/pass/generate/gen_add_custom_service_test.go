/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"errors"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAddCustomService1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &types.Project{},
	}
	// Mock functions
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		return nil, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	addBuildOrImagePassCallCount := 0
	addBuildOrImagePass = func(service *types.ServiceConfig, project *model.CGProject, serviceType string) {
		addBuildOrImagePassCallCount++
	}
	addNamePassCallCount := 0
	addNamePass = func(service *types.ServiceConfig, project *model.CGProject) {
		addNamePassCallCount++
	}
	addContainerNamePassCallCount := 0
	addContainerNamePass = func(service *types.ServiceConfig, project *model.CGProject) {
		addContainerNamePassCallCount++
	}
	addVolumesPassCallCount := 0
	addVolumesPass = func(service *types.ServiceConfig, project *model.CGProject, client *client.Client) {
		addVolumesPassCallCount++
	}
	addNetworksPassCallCount := 0
	addNetworksPass = func(service *types.ServiceConfig, project *model.CGProject, client *client.Client) {
		addNetworksPassCallCount++
	}
	addPortsPassCallCount := 0
	addPortsPass = func(service *types.ServiceConfig, _ *model.CGProject) {
		addPortsPassCallCount++
	}
	addEnvVarsPassCallCount := 0
	addEnvVarsPass = func(service *types.ServiceConfig, _ *model.CGProject) {
		addEnvVarsPassCallCount++
	}
	addEnvFilesPassCallCount := 0
	addEnvFilesPass = func(service *types.ServiceConfig, _ *model.CGProject) {
		addEnvFilesPassCallCount++
	}
	addRestartPassCallCount := 0
	addRestartPass = func(service *types.ServiceConfig, project *model.CGProject) {
		addRestartPassCallCount++
	}
	addDependsPassCallCount := 0
	addDependsPass = func(service *types.ServiceConfig, project *model.CGProject) {
		addDependsPassCallCount++
	}
	addDependantsPassCallCount := 0
	addDependantsPass = func(service *types.ServiceConfig, project *model.CGProject) {
		addDependantsPassCallCount++
	}
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	// Execute test
	GenerateAddCustomService(project, model.TemplateTypeFrontend)
	// Assert
	assert.Equal(t, 1, addBuildOrImagePassCallCount)
	assert.Equal(t, 1, addNamePassCallCount)
	assert.Equal(t, 1, addContainerNamePassCallCount)
	assert.Equal(t, 1, addVolumesPassCallCount)
	assert.Equal(t, 1, addNetworksPassCallCount)
	assert.Equal(t, 1, addPortsPassCallCount)
	assert.Equal(t, 1, addEnvVarsPassCallCount)
	assert.Equal(t, 1, addEnvFilesPassCallCount)
	assert.Equal(t, 1, addRestartPassCallCount)
	assert.Equal(t, 1, addDependsPassCallCount)
	assert.Equal(t, 1, addDependantsPassCallCount)
	assert.Equal(t, 1, pelCallCount)
}

func TestGenerateAddCustomService2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &types.Project{},
	}
	// Mock functions
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		return nil, errors.New("Error message")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Could not intanciate Docker client. Please check your Docker installation", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	addBuildOrImagePass = func(service *types.ServiceConfig, project *model.CGProject, serviceType string) {
		assert.Fail(t, "Unexpected call of addBuildOrImagePass")
	}
	// Execute test
	GenerateAddCustomService(project, model.TemplateTypeBackend)
}
