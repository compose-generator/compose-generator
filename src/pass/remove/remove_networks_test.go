/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package pass

import (
	"compose-generator/model"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------- RemoveNetworks ------------------------------------------

func TestRemoveNetworks1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{
		Networks: map[string]*spec.ServiceNetworkConfig{
			"test-net":   nil,
			"second-net": nil,
		},
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				"test-net": {
					Name: "test-net",
					External: spec.External{
						External: false,
					},
				},
				"second-net": {
					Name: "second-net",
					External: spec.External{
						External: true,
					},
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{},
		},
	}
	// Mock functions
	callCount := 0
	getServicesWhichUseNetworkMockable = func(networkName string, service *spec.ServiceConfig, project *model.CGProject) []spec.ServiceConfig {
		callCount++
		assert.Contains(t, []string{"test-net", "second-net"}, networkName)
		return []spec.ServiceConfig{}
	}
	// Execute test
	RemoveNetworks(service, project)
	// Assert
	assert.Equal(t, 2, callCount)
	assert.Equal(t, expectedProject, project)
}

func TestRemoveNetworks2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{
		Networks: map[string]*spec.ServiceNetworkConfig{
			"test-net": nil,
		},
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				"test-net": {
					Name: "test-net",
					External: spec.External{
						External: false,
					},
				},
				"second-net": {
					Name: "second-net",
					External: spec.External{
						External: true,
					},
				},
			},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				"second-net": {
					Name: "second-net",
					External: spec.External{
						External: true,
					},
				},
			},
		},
	}
	// Mock functions
	callCount := 0
	getServicesWhichUseNetworkMockable = func(networkName string, service *spec.ServiceConfig, project *model.CGProject) []spec.ServiceConfig {
		callCount++
		assert.Equal(t, "test-net", networkName)
		return []spec.ServiceConfig{
			{},
		}
	}
	// Execute test
	RemoveNetworks(service, project)
	// Assert
	assert.Equal(t, 1, callCount)
	assert.Equal(t, expectedProject, project)
}

// ------------------------------------- getServicesWhichUseNetwork ------------------------------------

func TestGetServicesWhichUseNetwork(t *testing.T) {
	// Test data
	networkName := "test-net"
	service := &spec.ServiceConfig{
		Name: "current-service",
		Networks: map[string]*spec.ServiceNetworkConfig{
			"test-net": nil,
		},
	}
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "other-service-1",
					Networks: map[string]*spec.ServiceNetworkConfig{
						"test-net": {},
					},
				},
				{
					Name: "other-service-2",
					Networks: map[string]*spec.ServiceNetworkConfig{
						"test-net": {},
					},
				},
				{
					Name: "current-service",
					Networks: map[string]*spec.ServiceNetworkConfig{
						"test-net": {},
					},
				},
			},
			Networks: spec.Networks{
				"test-net": {
					Name: "test-net",
					External: spec.External{
						External: false,
					},
				},
			},
		},
	}
	// Execute test
	result := getServicesWhichUseNetwork(networkName, service, project)
	// Assert
	assert.Equal(t, []spec.ServiceConfig{project.Composition.Services[0], project.Composition.Services[1]}, result)
}
