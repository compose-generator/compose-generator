package pass

import (
	"compose-generator/model"
	"errors"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestAddNetworks1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	// Mock functions
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	yesNoCallCount := 0
	YesNoQuestion = func(question string, defaultValue bool) bool {
		yesNoCallCount++
		switch yesNoCallCount {
		case 1:
			assert.Equal(t, "Do you want to add networks to your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Do you want to add an external network (y) or create a new one (N)?", question)
			assert.False(t, defaultValue)
		case 3:
			assert.Equal(t, "Add another network?", question)
			assert.True(t, defaultValue)
		case 4:
			assert.Equal(t, "Do you want to add an external network (y) or create a new one (N)?", question)
			assert.False(t, defaultValue)
			return false
		case 5:
			assert.Equal(t, "Add another network?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	askForExternalNetworkMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		if yesNoCallCount != 2 {
			assert.Fail(t, "Unexpected call of askForExternalNetwork")
		}
	}
	askForNewNetworkMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		if yesNoCallCount != 4 {
			assert.Fail(t, "Unexpected call of askForNewNetwork")
		}
	}
	// Execute test
	AddNetworks(service, project, cli)
}

func TestAddNetworks2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{}
	// Mock functions
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to add networks to your service?", question)
		assert.False(t, defaultValue)
		return false
	}
	Pel = func() {
		assert.Fail(t, "Unexpected call of Pel")
	}
	// Execute test
	AddNetworks(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
}

func TestAskForExternalNetwork1(t *testing.T) {
	// Test data
	testNetworkName := "Renamed network"
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedProject := &model.CGProject{
		Composition: &spec.Project{
			Networks: spec.Networks{
				testNetworkName: spec.NetworkConfig{
					Name: testNetworkName,
					External: spec.External{
						Name:     "Existing network 2",
						External: true,
					},
				},
			},
		},
	}
	expectedService := &spec.ServiceConfig{
		Networks: map[string]*spec.ServiceNetworkConfig{
			testNetworkName: nil,
		},
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	ListDockerNetworks = func(client *client.Client) ([]types.NetworkResource, error) {
		return []types.NetworkResource{
			{
				Name: "Existing network 1",
			},
			{
				Name: "Existing network 2",
			},
		}, nil
	}
	MenuQuestionIndex = func(label string, items []string) int {
		assert.Equal(t, "Which one?", label)
		assert.EqualValues(t, []string{"Existing network 1", "Existing network 2"}, items)
		return 1
	}
	TextQuestionWithDefault = func(question, defaultValue string) string {
		assert.Equal(t, "How do you want to call the network internally?", question)
		assert.Equal(t, "Existing network 2", defaultValue)
		return testNetworkName
	}
	Error = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of Error")
	}
	// Execute test
	askForExternalNetwork(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}

func TestAskForExternalNetwork2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedProject := &model.CGProject{}
	expectedService := &spec.ServiceConfig{}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	ListDockerNetworks = func(client *client.Client) ([]types.NetworkResource, error) {
		return []types.NetworkResource{}, nil
	}
	MenuQuestionIndex = func(label string, items []string) int {
		assert.Fail(t, "Unexpected call of MenuQuestionIndex")
		return 0
	}
	Error = func(description string, err error, exit bool) {
		assert.Equal(t, "There is no external network existing", description)
		assert.Nil(t, err)
		assert.False(t, exit)
	}
	// Execute test
	askForExternalNetwork(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}

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
