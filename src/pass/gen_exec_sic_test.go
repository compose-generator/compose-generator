package pass

import (
	"compose-generator/model"
	"testing"

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
	executeOnLinux = func(c string) {
		executeLinuxCallCount++
		if executeLinuxCallCount == 1 {
			assert.Equal(t, "ls . && mkdir test", c)
		} else {
			assert.Equal(t, "cd ./spring-gradle && touch env.env", c)
		}
	}
	doneCalled := false
	done = func() {
		doneCalled = true
	}
	pCallCount := 0
	p = func(text string) {
		pCallCount++
		if pCallCount == 1 {
			assert.Equal(t, "Generating configuration for Angular ... ", text)
		} else {
			assert.Equal(t, "Generating configuration for Spring Gradle ... ", text)
		}
	}
	// Execute test
	GenerateExecServiceInitCommands(project, selectedTemplates)
	// Assert
	assert.True(t, doneCalled)
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
	done = func() {
		doneCalled = true
	}
	// Execute test
	GenerateExecServiceInitCommands(project, selectedTemplates)
	// Assert
	assert.False(t, doneCalled)
}
