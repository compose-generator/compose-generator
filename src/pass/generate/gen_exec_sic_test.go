/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

func TestGenerateExecServiceInitCommands1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Vars: map[string]string{
			"SPRING_GRADLE_SOURCE_DIR": "./spring-gradle",
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:          "Angular",
				ServiceInitCmd: []string{"ls .", "mkdir test"},
			},
		},
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Label:          "Spring Gradle",
				ServiceInitCmd: []string{"cd ${{SPRING_GRADLE_SOURCE_DIR}}", "touch env.env"},
			},
		},
	}
	// Mock functions
	executeLinuxCallCount := 0
	executeOnToolbox = func(c string) {
		executeLinuxCallCount++
		if executeLinuxCallCount == 1 {
			assert.Equal(t, "ls . && mkdir test", c)
		} else {
			assert.Equal(t, "cd ./spring-gradle && touch env.env", c)
		}
	}
	startProcessCallCount := 0
	startProcess = func(text string) (s *spinner.Spinner) {
		startProcessCallCount++
		if startProcessCallCount == 1 {
			assert.Equal(t, "Generating configuration for Angular ...", text)
		} else {
			assert.Equal(t, "Generating configuration for Spring Gradle ...", text)
		}
		return nil
	}
	stopProcessCalled := false
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
		stopProcessCalled = true
	}
	// Execute test
	GenerateExecServiceInitCommands(project, selectedTemplates)
	// Assert
	assert.True(t, stopProcessCalled)
	assert.Equal(t, 2, startProcessCallCount)
}

func TestGenerateExecServiceInitCommands2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label:          "Wordpress",
				DemoAppInitCmd: []string{},
			},
		},
	}
	// Mock functions
	doneCalled := false
	stopProcess = func(s *spinner.Spinner) {
		doneCalled = true
	}
	// Execute test
	GenerateExecServiceInitCommands(project, selectedTemplates)
	// Assert
	assert.False(t, doneCalled)
}
