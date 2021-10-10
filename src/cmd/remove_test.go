/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------- Remove -------------------------------------------------------------------

func TestRemove1(t *testing.T) {
	/*// Test data
	context := &cli.Context{

	}
	// Mock functions

	// Execute test
	Remove(context)
	// Assert
	*/
}

// ------------------------------------------------------------------ removeService ----------------------------------------------------------------

func TestRemoveService1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		ForceConfig: false,
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "service1",
				},
				{
					Name: "service2",
				},
				{
					Name: "service3",
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		ForceConfig: false,
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "service1",
				},
				{
					Name: "service3",
				},
			},
		},
	}
	serviceName := "service2"
	// Mock functions
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to remove service '"+serviceName+"'?", question)
		assert.False(t, defaultValue)
		return true
	}
	removeVolumesPassCallCount := 0
	removeVolumesPass = func(service *spec.ServiceConfig, project *model.CGProject) {
		removeVolumesPassCallCount++
	}
	removeNetworksPassCallCount := 0
	removeNetworksPass = func(service *spec.ServiceConfig, project *model.CGProject) {
		removeNetworksPassCallCount++
	}
	removeDependenciesPassCallCount := 0
	removeDependenciesPass = func(service *spec.ServiceConfig, project *model.CGProject) {
		removeDependenciesPassCallCount++
	}
	// Execute test
	removeService(project, serviceName, false)
	// Assert
	assert.Zero(t, printErrorCallCount)
	assert.Equal(t, 1, removeVolumesPassCallCount)
	assert.Equal(t, 1, removeNetworksPassCallCount)
	assert.Equal(t, 1, removeDependenciesPassCallCount)
	assert.Equal(t, expectedProject, project)
}

func TestRemoveService2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		ForceConfig: false,
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "service1",
				},
				{
					Name: "service2",
				},
				{
					Name: "service3",
				},
			},
		},
	}
	serviceName := "service4"
	// Mock functions
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Service not found", description)
		assert.NotNil(t, err)
		assert.False(t, exit)
	}
	// Execute test
	removeService(project, serviceName, false)
	// Assert
	assert.Equal(t, 1, printErrorCallCount)
}

func TestRemoveService3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		ForceConfig: false,
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "service1",
				},
				{
					Name: "service2",
				},
				{
					Name: "service3",
				},
			},
		},
	}
	serviceName := "service2"
	// Mock functions
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
	}
	yesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you really want to remove service '"+serviceName+"'?", question)
		assert.False(t, defaultValue)
		return false
	}
	removeVolumesPassCallCount := 0
	removeVolumesPass = func(service *spec.ServiceConfig, project *model.CGProject) {
		removeVolumesPassCallCount++
	}
	// Execute test
	removeService(project, serviceName, false)
	// Assert
	assert.Zero(t, printErrorCallCount)
	assert.Zero(t, removeVolumesPassCallCount)
}

// ------------------------------------------------------------ removeServiceFromProject -----------------------------------------------------------

func TestRemoveServiceFromProject(t *testing.T) {
	// Test data
	services := &spec.Services{
		{
			Name: "service1",
		},
		{
			Name: "service2",
		},
		{
			Name: "service3",
		},
		{
			Name: "service4",
		},
	}
	expectedServices := &spec.Services{
		{
			Name: "service1",
		},
		{
			Name: "service2",
		},
		{
			Name: "service4",
		},
	}
	// Execute test
	removeServiceFromProject(services, "service3")
	// Assert
	assert.EqualValues(t, expectedServices, services)
}
