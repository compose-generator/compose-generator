/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	"errors"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
)

// --------------------------------------------------------------- isTemplateExisting --------------------------------------------------------------

func TestIsTemplateExisting1(t *testing.T) {
	// Test data
	name := "Template 1"
	customTemplatePath := "../templates"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatePath
	}
	fileExists = func(path string) bool {
		assert.Equal(t, customTemplatePath+"/"+name, path)
		return false
	}
	// Execute test
	result := isTemplateExisting(name)
	// Assert
	assert.False(t, result)
}

func TestIsTemplateExisting2(t *testing.T) {
	// Test data
	name := "Template 1"
	customTemplatePath := "../templates"
	// Mock functions
	getCustomTemplatesPath = func() string {
		return customTemplatePath
	}
	fileExists = func(path string) bool {
		assert.Equal(t, customTemplatePath+"/"+name, path)
		return true
	}
	logWarning = func(message string) {
		assert.Equal(t, "Template with the name 'Template 1' already exists", message)
	}
	// Execute test
	result := isTemplateExisting(name)
	// Assert
	assert.True(t, result)
}

// ------------------------------------------------------ copyVolumesAndBuildContextsToTemplate ----------------------------------------------------

func TestCopyVolumesAndBuildContextsToTemplate1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "./volume1",
							Type:   spec.VolumeTypeBind,
						},
						{
							Source: "./volume1/volume2",
							Type:   spec.VolumeTypeBind,
						},
					},
				},
				{
					Build: &spec.BuildConfig{
						Context: "./frontend-angular/Dockerfile",
					},
					Volumes: []spec.ServiceVolumeConfig{
						{
							Source: "../volume4",
							Type:   spec.VolumeTypeBind,
						},
					},
				},
			},
		},
	}
	targetDir := "../templates"
	// Mock functions
	absCallCount := 0
	abs = func(path string) (string, error) {
		absCallCount++
		switch absCallCount {
		case 1:
			assert.Equal(t, ".", path)
			return "/usr/lib/compose-generator/templates", nil
		case 2:
			assert.Equal(t, "./volume1", path)
			return "/usr/lib/compose-generator/templates/Template 1/volume1", nil
		case 3:
			assert.Equal(t, "../volume4", path)
			return "/usr/lib/compose-generator/templates/Template 2/volume4", nil
		case 4:
			assert.Equal(t, "frontend-angular", path)
			return "", errors.New("Error message")
		}
		return "", nil
	}
	rel = func(basepath, targpath string) (string, error) {
		assert.Equal(t, "/usr/lib/compose-generator/templates", basepath)
		switch targpath {
		case "/usr/lib/compose-generator/templates/Template 1/volume1":
			assert.Equal(t, "/usr/lib/compose-generator/templates/Template 1/volume1", targpath)
			return "./Template 1/volume1", nil
		case "/usr/lib/compose-generator/templates/Template 2/volume4":
			assert.Equal(t, "/usr/lib/compose-generator/templates/Template 2/volume4", targpath)
			return "", errors.New("Error message 1")
		}
		return "", nil
	}
	copyDirCallCount := 0
	copyDir = func(src, dest string, opt ...copy.Options) error {
		copyDirCallCount++
		assert.Zero(t, len(opt))
		if copyDirCallCount == 1 {
			assert.Equal(t, "./volume1", src)
			assert.Equal(t, "../templates/./Template 1/volume1", dest)
			return nil
		}
		assert.Equal(t, "frontend-angular", src)
		assert.Equal(t, "../templates/", dest)
		return errors.New("Error message")
	}
	logErrorCallCount := 0
	logError = func(message string, exit bool) {
		logErrorCallCount++
		if logErrorCallCount == 1 {
			assert.Equal(t, "Could not copy volume / build context '../volume4'", message)
			assert.False(t, exit)
		} else {
			assert.Equal(t, "Could not find absolute path of volume / build context dir", message)
			assert.True(t, exit)
		}
	}
	logWarning = func(message string) {
		assert.Equal(t, "Could not copy volume / build context from 'frontend-angular' to '../templates/'", message)
	}
	// Execute test
	copyVolumesAndBuildContextsToTemplate(project, targetDir)
	// Assert
	assert.Equal(t, 4, absCallCount)
	assert.Equal(t, 2, copyDirCallCount)
	assert.Equal(t, 2, logErrorCallCount)
}

func TestCopyVolumesAndBuildContextsToTemplate2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	targetDir := "../templates"
	// Mock functions
	abs = func(path string) (string, error) {
		assert.Equal(t, ".", path)
		return "", errors.New("Error message")
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Could not find absolute path of current dir", message)
		assert.True(t, exit)
	}
	// Execute test
	copyVolumesAndBuildContextsToTemplate(project, targetDir)
}
