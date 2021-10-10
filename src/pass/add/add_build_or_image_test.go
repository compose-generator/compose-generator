/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package pass

import (
	"compose-generator/model"
	"errors"
	"testing"

	diu "github.com/compose-generator/diu/model"
	spec "github.com/compose-spec/compose-go/types"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------- TestAddBuildOrImage ----------------------------------------

func TestAddBuildOrImage1(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Build: &spec.BuildConfig{
			Context:    ".",
			Dockerfile: "./Dockerfile",
		},
	}
	// Mock functions
	textQuestionWithDefault = func(question, defaultValue string) (result string) {
		assert.Equal(t, "Where is your Dockerfile located?", question)
		assert.Equal(t, "./Dockerfile", defaultValue)
		return "./Dockerfile"
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Build from source?", question)
		assert.False(t, defaultValue)
		return true
	}
	printError = func(description string, err error, exit bool) {}
	fileExists = func(path string) bool {
		return true
	}
	menuQuestion = func(label string, items []string) (result string) {
		return ""
	}
	// Execute test
	AddBuildOrImage(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}

func TestAddBuildOrImage2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	// Mock functions
	textQuestionWithDefault = func(question, defaultValue string) (result string) {
		assert.Equal(t, "Where is your Dockerfile located?", question)
		assert.Equal(t, "./Dockerfile", defaultValue)
		return "./Dockerfile"
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Build from source?", question)
		assert.False(t, defaultValue)
		return true
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "The Dockerfile could not be found", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	fileExists = func(path string) bool {
		return false
	}
	menuQuestion = func(label string, items []string) (result string) {
		return ""
	}
	// Execute test
	AddBuildOrImage(service, project)
}

func TestAddBuildOrImage3(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	service := &spec.ServiceConfig{}
	expectedService := &spec.ServiceConfig{
		Name:  "backend-spice",
		Image: "ghcr.io/chillibits/spice:0.3.0",
	}
	testManifest := diu.DockerManifest{
		SchemaV2Manifest: diu.SchemaV2Manifest{
			Layers: []diu.Layer{
				{}, {}, {}, {}, {}, {}, {},
			},
		},
	}
	// Mock functions
	textQuestionCallCounter := 0
	textQuestionWithDefault = func(question, defaultValue string) (result string) {
		textQuestionCallCounter++
		if textQuestionCallCounter == 1 {
			assert.Equal(t, "From which registry do you want to pick?", question)
			assert.Equal(t, "docker.io", defaultValue)
			result = "ghcr.io"
		} else {
			assert.Equal(t, "Which Image do you want to use? (e.g. chillibits/ccom:0.8.0)", question)
			assert.Equal(t, "hello-world", defaultValue)
			result = "chillibits/spice:0.3.0"
		}
		return
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Build from source?", question)
		assert.False(t, defaultValue)
		return false
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "The Dockerfile could not be found", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	fileExists = func(path string) bool {
		return false
	}
	menuQuestion = func(label string, items []string) (result string) {
		assert.Equal(t, "Which type is the closest match for this service?", label)
		assert.EqualValues(t, []string{"frontend", "backend", "database", "db-admin"}, items)
		return "backend"
	}
	getImageManifest = func(image string) (diu.DockerManifest, error) {
		assert.Equal(t, "ghcr.io/chillibits/spice:0.3.0", image)
		return testManifest, nil
	}
	pel = func() {}
	p = func(text string) {
		assert.Equal(t, "Searching image ... ", text)
	}
	success = func(text string) {
		assert.Equal(t, " found - 7 layer(s)", text)
	}
	// Execute test
	AddBuildOrImage(service, project)
	// Assert
	assert.Equal(t, expectedService, service)
}

// --------------------------------------- TestSearchRemoteImage ---------------------------------------

func TestSearchRemoteImage1(t *testing.T) {
	// Test data
	testManifest := diu.DockerManifest{
		SchemaV2Manifest: diu.SchemaV2Manifest{
			Layers: []diu.Layer{
				{}, {}, {}, {}, {}, {}, {},
			},
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	p = func(text string) {
		assert.Equal(t, "Searching image ... ", text)
	}
	success = func(text string) {
		assert.Equal(t, " found - 7 layer(s)", text)
	}
	printError = func(description string, err error, exit bool) {}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		return false
	}
	getImageManifest = func(image string) (diu.DockerManifest, error) {
		assert.Equal(t, "ghcr.io/chillibits/compose-generator", image)
		return testManifest, nil
	}
	// Execute test
	result := searchRemoteImage("ghcr.io/", "chillibits/compose-generator")
	// Assert
	assert.False(t, result)
	assert.Equal(t, 2, pelCallCount)
}

func TestSearchRemoteImage2(t *testing.T) {
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	p = func(text string) {
		assert.Equal(t, "Searching image ... ", text)
	}
	success = func(text string) {}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, " not found or no access", description)
		assert.Nil(t, err)
		assert.False(t, exit)
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Choose another image (Y) or proceed anyway (n)?", question)
		assert.True(t, defaultValue)
		return false
	}
	getImageManifest = func(image string) (diu.DockerManifest, error) {
		assert.Equal(t, "chillibits/compose-generator", image)
		return diu.DockerManifest{}, errors.New("Could not parse manifest")
	}
	// Execute test
	result := searchRemoteImage("", "chillibits/compose-generator")
	// Assert
	assert.False(t, result)
	assert.Equal(t, 1, pelCallCount)
}

func TestSearchRemoteImage3(t *testing.T) {
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	p = func(text string) {
		assert.Equal(t, "Searching image ... ", text)
	}
	success = func(text string) {}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, " not found or no access", description)
		assert.Nil(t, err)
		assert.False(t, exit)
	}
	yesNoQuestion = func(question string, defaultValue bool) (result bool) {
		assert.Equal(t, "Choose another image (Y) or proceed anyway (n)?", question)
		assert.True(t, defaultValue)
		return true
	}
	getImageManifest = func(image string) (diu.DockerManifest, error) {
		assert.Equal(t, "chillibits/compose-generator", image)
		return diu.DockerManifest{}, errors.New("Could not parse manifest")
	}
	// Execute test
	result := searchRemoteImage("", "chillibits/compose-generator")
	// Assert
	assert.True(t, result)
	assert.Equal(t, 1, pelCallCount)
}
