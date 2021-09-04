package pass

import (
	"compose-generator/model"
	"compose-generator/util"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------ AddVolumes ------------------------------------------------------

func TestAddVolumes1(t *testing.T) {
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
			assert.Equal(t, "Do you want to add volumes to your service?", question)
			assert.False(t, defaultValue)
		case 2:
			assert.Equal(t, "Do you want to add an existing external volume (y) or link a directory / file (N)?", question)
			assert.False(t, defaultValue)
		case 3:
			assert.Equal(t, "Add another volume?", question)
			assert.True(t, defaultValue)
		case 4:
			assert.Equal(t, "Do you want to add an existing external volume (y) or link a directory / file (N)?", question)
			assert.False(t, defaultValue)
			return false
		case 5:
			assert.Equal(t, "Add another volume?", question)
			assert.True(t, defaultValue)
			return false
		}
		return true
	}
	askForExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		if yesNoCallCount != 2 {
			assert.Fail(t, "Unexpected call of askForExternalVolume")
		}
	}
	askForFileVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject) {
		if yesNoCallCount != 4 {
			assert.Fail(t, "Unexpected call of askForFileVolume")
		}
	}
	Pel = func() {}
	// Execute test
	AddVolumes(service, project, cli)
}

func TestAddVolumes2(t *testing.T) {
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
		assert.Equal(t, "Do you want to add volumes to your service?", question)
		assert.False(t, defaultValue)
		return false
	}
	Pel = func() {
		assert.Fail(t, "Unexpected call of Pel")
	}
	// Execute test
	AddVolumes(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
}

// ------------------------------------------------- askForExternalVolume -------------------------------------------------

func TestAskForExternalVolume1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to select an existing one (Y) or do you want to create one (n)?", question)
		assert.True(t, defaultValue)
		return true
	}
	callCount := 0
	askForExistingExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		callCount++
	}
	askForNewExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		assert.Fail(t, "Unexpected call of askForNewExternalVolume")
	}
	// Execute test
	askForExternalVolume(service, project, cli)
	// Assert
	assert.Equal(t, 1, callCount)
}

func TestAskForExternalVolume2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to select an existing one (Y) or do you want to create one (n)?", question)
		assert.True(t, defaultValue)
		return false
	}
	askForExistingExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		assert.Fail(t, "Unexpected call of askForExistingExternalVolume")
	}
	callCount := 0
	askForNewExternalVolumeMockable = func(service *spec.ServiceConfig, project *model.CGProject, client *client.Client) {
		callCount++
	}
	// Execute test
	askForExternalVolume(service, project, cli)
	// Assert
	assert.Equal(t, 1, callCount)
}

// --------------------------------------------------- askForFileVolume ---------------------------------------------------

func TestAskForExistingExternalVolume1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: false,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Volumes: []spec.ServiceVolumeConfig{
			{
				Source:   "External volume 2",
				Target:   "./directory/inside/container.conf",
				Type:     spec.VolumeTypeVolume,
				ReadOnly: false,
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: false,
		},
		Composition: &spec.Project{
			Volumes: spec.Volumes{
				"External volume 2": {
					Name: "External volume 2",
					External: spec.External{
						Name:     "External volume 2",
						External: true,
					},
				},
			},
		},
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	ListDockerVolumes = func(client *client.Client) (volume.VolumeListOKBody, error) {
		return volume.VolumeListOKBody{
			Volumes: []*types.Volume{
				{
					Name:   "External volume 1",
					Driver: "local",
				},
				{
					Name:   "External volume 2",
					Driver: "local",
				},
			},
		}, nil
	}
	Error = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of Error")
	}
	MenuQuestionIndex = func(label string, items []string) int {
		assert.Equal(t, "Which one?", label)
		assert.EqualValues(t, []string{"External volume 1 | Driver: local", "External volume 2 | Driver: local"}, items)
		return 1
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "Directory / file inside the container:", question)
		return "./directory/inside/container.conf"
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of YesNoQuestion")
		return false
	}
	// Execute test
	askForExistingExternalVolume(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}

func TestAskForExistingExternalVolume2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: true,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Volumes: []spec.ServiceVolumeConfig{
			{
				Source:   "External volume 1",
				Target:   "./directory/inside/container.conf",
				Type:     spec.VolumeTypeVolume,
				ReadOnly: true,
			},
		},
	}
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: true,
		},
		Composition: &spec.Project{
			Volumes: spec.Volumes{
				"External volume 1": {
					Name: "External volume 1",
					External: spec.External{
						Name:     "External volume 1",
						External: true,
					},
				},
			},
		},
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		assert.Fail(t, "Could not create Docker client for testing")
	}
	// Mock functions
	ListDockerVolumes = func(client *client.Client) (volume.VolumeListOKBody, error) {
		return volume.VolumeListOKBody{
			Volumes: []*types.Volume{
				{
					Name:   "External volume 1",
					Driver: "local",
				},
				{
					Name:   "External volume 2",
					Driver: "local",
				},
			},
		}, nil
	}
	Error = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of Error")
	}
	MenuQuestionIndex = func(label string, items []string) int {
		assert.Equal(t, "Which one?", label)
		assert.EqualValues(t, []string{"External volume 1 | Driver: local", "External volume 2 | Driver: local"}, items)
		return 0
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "Directory / file inside the container:", question)
		return "./directory/inside/container.conf"
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to make the volume read-only?", question)
		assert.False(t, defaultValue)
		return true
	}
	// Execute test
	askForExistingExternalVolume(service, project, cli)
	// Assert
	assert.Equal(t, expectedService, service)
	assert.Equal(t, expectedProject, project)
}

// --------------------------------------------- askForExistingExternalVolume ---------------------------------------------

// ----------------------------------------------- askForNewExternalVolume ------------------------------------------------

func TestAskForFileVolume1(t *testing.T) {
	// Test data
	outerDir := " ./volume-dir/outer"
	innerDir := "/test/destination/dir "
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: false,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Volumes: []spec.ServiceVolumeConfig{
			{
				Type:     spec.VolumeTypeBind,
				Source:   "./volume-dir/outer",
				Target:   "/test/destination/dir",
				ReadOnly: false,
			},
		},
	}
	// Mock functions
	TextQuestionWithSuggestions = func(question string, fn util.Suggest) string {
		assert.Equal(t, "Directory / file on host machine:", question)
		return outerDir
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "Directory / file inside the container:", question)
		return innerDir
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Fail(t, "Unexpected call of YesNoQuestion")
		return false
	}
	// Execute test
	askForFileVolume(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}

func TestAskForFileVolume2(t *testing.T) {
	// Test data
	outerDir := " volume-dir/outer"
	innerDir := "/test/destination/dir "
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			AdvancedConfig: true,
		},
	}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Volumes: []spec.ServiceVolumeConfig{
			{
				Type:     spec.VolumeTypeBind,
				Source:   "./volume-dir/outer",
				Target:   "/test/destination/dir",
				ReadOnly: false,
			},
		},
	}
	// Mock functions
	TextQuestionWithSuggestions = func(question string, fn util.Suggest) string {
		assert.Equal(t, "Directory / file on host machine:", question)
		return outerDir
	}
	TextQuestion = func(question string) string {
		assert.Equal(t, "Directory / file inside the container:", question)
		return innerDir
	}
	YesNoQuestion = func(question string, defaultValue bool) bool {
		assert.Equal(t, "Do you want to make the volume read-only?", question)
		assert.False(t, defaultValue)
		return false
	}
	// Execute test
	askForFileVolume(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}
