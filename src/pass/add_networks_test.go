package pass

import (
	"compose-generator/model"
	"errors"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestAskForNewNetwork1(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	testNetworkName := "Test network"
	expectedService := &spec.ServiceConfig{
		Networks: map[string]*spec.ServiceNetworkConfig{
			testNetworkName: {},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				testNetworkName: spec.NetworkConfig{
					Name: testNetworkName,
					External: spec.External{
						External: true,
						Name:     testNetworkName,
					},
				},
			},
		},
	}
	// Mock functions
	CreateDockerNetwork = func(_ *client.Client, networkName string) error {
		assert.Equal(t, testNetworkName, networkName)
		return nil
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "How do you want to call the new network?", question)
		return testNetworkName
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to create it as an external network and link it in?", question)
		assert.False(t, defaultValue)
		return true
	}
	// Execute test
	askForNewNetwork(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}

func TestAskForNewNetwork2(t *testing.T) {
	// Test data
	service := &spec.ServiceConfig{}
	project := &model.CGProject{}
	cli := &client.Client{}
	testNetworkName := "Test network"
	expectedService := &spec.ServiceConfig{
		Networks: map[string]*spec.ServiceNetworkConfig{
			testNetworkName: {},
		},
	}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				testNetworkName: spec.NetworkConfig{
					Name:     testNetworkName,
					External: spec.External{},
				},
			},
		},
	}
	// Mock functions
	CreateDockerNetwork = func(_ *client.Client, networkName string) error {
		assert.Fail(t, "Unexpected call of CreateDockerNetwork")
		return errors.New("Error")
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "How do you want to call the new network?", question)
		return testNetworkName
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to create it as an external network and link it in?", question)
		assert.False(t, defaultValue)
		return false
	}
	Error = func(description string, err error, exit bool) {}
	// Execute test
	askForNewNetwork(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}
