/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"compose-generator/project"
	"testing"

	"github.com/briandowns/spinner"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// --------------------------------------------- Generate ----------------------------------------------

func TestGenerate1(t *testing.T) {
	// Test data
	predefinedTemplatesPath := "../predefined-templates"
	proj := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Name: "angular",
			},
		},
		ProxyService: []model.PredefinedTemplateConfig{
			{
				Name: "nginx",
			},
		},
	}
	expectedProj := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
		ReadmeChildPaths: []string{
			predefinedTemplatesPath + "/INSTRUCTIONS_HEADER.md",
		},
	}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Generating configuration from 2 template(s) ...", text)
		return nil
	}
	getPredefinedServicesPath = func() string {
		return predefinedTemplatesPath
	}
	generateServiceCallCount := 0
	generateServiceMockable = func(
		proj *model.CGProject,
		selectedTemplates *model.SelectedTemplates,
		template model.PredefinedTemplateConfig,
		templateType, serviceName string,
	) {
		generateServiceCallCount++
		if generateServiceCallCount == 1 {
			assert.Equal(t, model.TemplateTypeFrontend, templateType)
			assert.Equal(t, selectedTemplates.FrontendServices[0], template)
		} else {
			assert.Equal(t, model.TemplateTypeProxy, templateType)
			assert.Equal(t, selectedTemplates.ProxyService[0], template)
		}
	}
	stopProcessCallCount := 0
	stopProcess = func(s *spinner.Spinner) {
		stopProcessCallCount++
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	Generate(proj, selectedTemplates)
	// Assert
	assert.Equal(t, 1, pelCallCount)
	assert.Equal(t, 1, stopProcessCallCount)
	assert.Equal(t, expectedProj, proj)
}

func TestGenerate2(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	selectedTemplates := &model.SelectedTemplates{}
	// Mock functions
	pelCallCount := 0
	pel = func() {
		pelCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Fail(t, "Unexpected error of startProcess")
		return nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "No templates selected. Aborting ...", description)
		assert.Nil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	Generate(project, selectedTemplates)
	// Assert
	assert.Equal(t, 1, pelCallCount)
}

// ------------------------------------------ generateService ------------------------------------------

func TestGenerateService1(t *testing.T) {
	// Test data
	proj := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{},
		},
	}
	selectedTemplates := &model.SelectedTemplates{}
	template := model.PredefinedTemplateConfig{
		Name:  "faunadb",
		Label: "FaunaDB",
		Files: []model.File{
			{
				Path: "./docs.md",
				Type: model.FileTypeDocs,
			},
			{
				Path: "./env.env",
				Type: model.FileTypeEnv,
			},
			{
				Path: "./env.env",
				Type: model.FileTypeEnv,
			},
		},
	}
	templateType := model.TemplateTypeDatabase
	serviceName := "faunadb"
	expectedProj := &model.CGProject{
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Name: "faunadb",
				},
			},
		},
		ReadmeChildPaths: []string{
			"docs.md",
		},
		GitignorePatterns: []string{
			"./env.env",
		},
	}
	// Mock function
	loadTemplateService = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates, templateTypeName, serviceName string, options ...project.LoadOption) *spec.ServiceConfig {
		assert.Equal(t, 2, len(options))
		return &spec.ServiceConfig{
			Name: "faunadb",
		}
	}
	sliceContainsStringCallCount := 0
	sliceContainsString = func(slice []string, i string) bool {
		sliceContainsStringCallCount++
		assert.Equal(t, "./env.env", i)
		if sliceContainsStringCallCount == 1 {
			assert.Nil(t, slice)
			return false
		}
		assert.EqualValues(t, []string{"./env.env"}, slice)
		return true
	}
	// Execute test
	generateService(proj, selectedTemplates, template, templateType, serviceName)
	// Assert
	assert.Equal(t, expectedProj, proj)
}
