/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"errors"
	"os"
	"testing"

	"github.com/briandowns/spinner"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------- GenerateCopyVolumes ----------------------------------------

func TestGenerateCopyVolumes1(t *testing.T) {
	// Test data
	templatesPath := "../predefined-templates/path"
	project := &model.CGProject{
		Composition: &spec.Project{
			WorkingDir: "./",
		},
		Vars: model.Vars{
			"ANGULAR_BUILD_CONTEXT_DIR": "./angular",
			"ANGULAR_DATA":              "./volumes/angular-data",
			"POSTGRES_CONFIG":           "./vol/postgres-configuration",
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name: "angular",
				Type: "frontend",
				Volumes: []model.Volume{
					{
						DefaultValue: "./frontend-angular",
						Variable:     "ANGULAR_BUILD_CONTEXT_DIR",
					},
					{
						DefaultValue: "./volumes/angular-data",
						Variable:     "ANGULAR_DATA",
					},
				},
			},
		},
		DatabaseServices: []model.PredefinedTemplateConfig{
			{
				Name: "postgres",
				Type: "database",
				Volumes: []model.Volume{
					{
						DefaultValue: "./database-postgres",
						Variable:     "POSTGRES_CONFIG",
					},
				},
			},
		},
	}
	// Mock functions
	startProcess = func(text string) (s *spinner.Spinner) {
		assert.Equal(t, "Copying volumes ...", text)
		return nil
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
		assert.Nil(t, s)
	}
	copyVolumeCallCount := 0
	copyVolumeMockable = func(srcPath, dstPath string) {
		copyVolumeCallCount++
		switch copyVolumeCallCount {
		case 1:
			assert.Equal(t, "../predefined-templates/path/frontend/angular/frontend-angular", srcPath)
			assert.Equal(t, "angular", dstPath)
		case 2:
			assert.Equal(t, "../predefined-templates/path/frontend/angular/volumes/angular-data", srcPath)
			assert.Equal(t, "volumes/angular-data", dstPath)
		case 3:
			assert.Equal(t, "../predefined-templates/path/database/postgres/database-postgres", srcPath)
			assert.Equal(t, "vol/postgres-configuration", dstPath)
		}
	}
	getPredefinedServicesPath = func() string {
		return templatesPath
	}
	// Execute test
	GenerateCopyVolumes(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, stopProcessCallCount)
	assert.Equal(t, 3, copyVolumeCallCount)
}

// -------------------------------------------- copyVolume ---------------------------------------------

func TestCopyVolume1(t *testing.T) {
	// Test data
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return nil
	}
	logError = func(message string, exit bool) {
		assert.Fail(t, "Unexpected call of logError")
	}
	// Execute test
	copyVolume(srcPath, dstPath)
}

func TestCopyVolume2(t *testing.T) {
	// Test data
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return false
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	logError = func(message string, exit bool) {
		assert.Equal(t, "Could not create volume dir", message)
	}
	// Execute test
	copyVolume(srcPath, dstPath)
}

func TestCopyVolume3(t *testing.T) {
	// Test data
	srcPath := "../predefined-templates/path/type/template/volumes/volume1"
	dstPath := "volumes/volume1"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, srcPath, path)
		return true
	}
	mkdirAll = func(path string, perm os.FileMode) error {
		assert.Equal(t, dstPath, path)
		assert.Equal(t, os.FileMode(0777), perm)
		return errors.New("MkdirAll error")
	}
	copyFile = func(src, dest string, opt ...copy.Options) error {
		assert.Equal(t, srcPath, src)
		assert.Equal(t, dstPath, dest)
		return nil
	}
	// Execute test
	copyVolume(srcPath, dstPath)
}
