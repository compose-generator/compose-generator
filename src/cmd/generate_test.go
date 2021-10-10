/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package cmd

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------- Generate --------------------------------------------------------------------

// ---------------------------------------------------------------- generateProject ----------------------------------------------------------------

func TestGenerateProject1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
	}
	config := &model.GenerateConfig{
		FromFile: false,
	}
	availableTemplates := &model.AvailableTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label: "Angular",
			},
			{
				Label: "Wordpress",
			},
		},
		DatabaseServices: []model.PredefinedTemplateConfig{
			{
				Label: "Redis",
			},
			{
				Label: "OrientDB",
			},
		},
	}
	// Mock functions
	clearScreenCallCount := 0
	clearScreen = func() {
		clearScreenCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading predefined service templates ...", text)
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	getAvailablePredefinedTemplates = func() *model.AvailableTemplates {
		return availableTemplates
	}
	generateChooseFrontendsPassCallCount := 0
	generateChooseFrontendsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseFrontendsPassCallCount++
	}
	generateChooseBackendsPassCallCount := 0
	generateChooseBackendsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseBackendsPassCallCount++
	}
	generateChooseDatabasesPassCallCount := 0
	generateChooseDatabasesPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseDatabasesPassCallCount++
	}
	generateChooseDbAdminsPassCallCount := 0
	generateChooseDbAdminsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseDbAdminsPassCallCount++
	}
	generateChooseProxiesPassCallCount := 0
	generateChooseProxiesPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseProxiesPassCallCount++
	}
	generateChooseTlsHelpersPassCallCount := 0
	generateChooseTlsHelpersPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseTlsHelpersPassCallCount++
	}
	generatePassCallCount := 0
	generatePass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generatePassCallCount++
	}
	generateResolveDependencyGroupsPassCallCount := 0
	generateResolveDependencyGroupsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateResolveDependencyGroupsPassCallCount++
	}
	generateSecretsPassCallCount := 0
	generateSecretsPass = func(project *model.CGProject, selected *model.SelectedTemplates) {
		generateSecretsPassCallCount++
	}
	generateAddProfilesPassCallCount := 0
	generateAddProfilesPass = func(project *model.CGProject) {
		generateAddProfilesPassCallCount++
	}
	generateAddProxyNetworksCallCount := 0
	generateAddProxyNetworks = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateAddProxyNetworksCallCount++
	}
	generateCopyVolumesPassCallCount := 0
	generateCopyVolumesPass = func(project *model.CGProject) {
		generateCopyVolumesPassCallCount++
	}
	generateReplaceVarsInConfigFilesPassCallCount := 0
	generateReplaceVarsInConfigFilesPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateReplaceVarsInConfigFilesPassCallCount++
	}
	generateExecServiceInitCommandsPassCallCount := 0
	generateExecServiceInitCommandsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateExecServiceInitCommandsPassCallCount++
	}
	generateExecDemoAppInitCommandsPassCallCount := 0
	generateExecDemoAppInitCommandsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateExecDemoAppInitCommandsPassCallCount++
	}
	// Execute test
	generateProject(project, config)
	// Assert
	assert.Equal(t, 1, clearScreenCallCount)
	assert.Equal(t, 1, generateChooseFrontendsPassCallCount)
	assert.Equal(t, 1, generateChooseBackendsPassCallCount)
	assert.Equal(t, 1, generateChooseDatabasesPassCallCount)
	assert.Equal(t, 1, generateChooseDbAdminsPassCallCount)
	assert.Equal(t, 1, generateChooseProxiesPassCallCount)
	assert.Equal(t, 1, generateChooseTlsHelpersPassCallCount)
	assert.Equal(t, 1, generatePassCallCount)
	assert.Equal(t, 1, generateResolveDependencyGroupsPassCallCount)
	assert.Equal(t, 1, generateSecretsPassCallCount)
	assert.Equal(t, 1, generateAddProfilesPassCallCount)
	assert.Equal(t, 1, generateAddProxyNetworksCallCount)
	assert.Equal(t, 1, generateCopyVolumesPassCallCount)
	assert.Equal(t, 1, generateReplaceVarsInConfigFilesPassCallCount)
	assert.Equal(t, 1, generateExecServiceInitCommandsPassCallCount)
	assert.Equal(t, 1, generateExecDemoAppInitCommandsPassCallCount)
}

func TestGenerateProject2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			ProductionReady: true,
		},
	}
	config := &model.GenerateConfig{
		FromFile: true,
	}
	availableTemplates := &model.AvailableTemplates{
		FrontendServices: []model.PredefinedTemplateConfig{
			{
				Label: "Angular",
			},
			{
				Label: "Wordpress",
			},
		},
		DatabaseServices: []model.PredefinedTemplateConfig{
			{
				Label: "Redis",
			},
			{
				Label: "OrientDB",
			},
		},
	}
	// Mock functions
	clearScreenCallCount := 0
	clearScreen = func() {
		clearScreenCallCount++
	}
	startProcess = func(text string) *spinner.Spinner {
		assert.Equal(t, "Loading predefined service templates ...", text)
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	getAvailablePredefinedTemplates = func() *model.AvailableTemplates {
		return availableTemplates
	}
	generateChooseFrontendsPassCallCount := 0
	generateChooseFrontendsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseFrontendsPassCallCount++
	}
	generateChooseBackendsPassCallCount := 0
	generateChooseBackendsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseBackendsPassCallCount++
	}
	generateChooseDatabasesPassCallCount := 0
	generateChooseDatabasesPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseDatabasesPassCallCount++
	}
	generateChooseDbAdminsPassCallCount := 0
	generateChooseDbAdminsPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseDbAdminsPassCallCount++
	}
	generateChooseProxiesPassCallCount := 0
	generateChooseProxiesPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseProxiesPassCallCount++
	}
	generateChooseTlsHelpersPassCallCount := 0
	generateChooseTlsHelpersPass = func(project *model.CGProject, available *model.AvailableTemplates, selected *model.SelectedTemplates, config *model.GenerateConfig) {
		generateChooseTlsHelpersPassCallCount++
	}
	generatePassCallCount := 0
	generatePass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generatePassCallCount++
	}
	generateResolveDependencyGroupsPassCallCount := 0
	generateResolveDependencyGroupsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateResolveDependencyGroupsPassCallCount++
	}
	generateSecretsPassCallCount := 0
	generateSecretsPass = func(project *model.CGProject, selected *model.SelectedTemplates) {
		generateSecretsPassCallCount++
	}
	generateAddProfilesPassCallCount := 0
	generateAddProfilesPass = func(project *model.CGProject) {
		generateAddProfilesPassCallCount++
	}
	generateAddProxyNetworksCallCount := 0
	generateAddProxyNetworks = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateAddProxyNetworksCallCount++
	}
	generateCopyVolumesPassCallCount := 0
	generateCopyVolumesPass = func(project *model.CGProject) {
		generateCopyVolumesPassCallCount++
	}
	generateReplaceVarsInConfigFilesPassCallCount := 0
	generateReplaceVarsInConfigFilesPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateReplaceVarsInConfigFilesPassCallCount++
	}
	generateExecServiceInitCommandsPassCallCount := 0
	generateExecServiceInitCommandsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateExecServiceInitCommandsPassCallCount++
	}
	generateExecDemoAppInitCommandsPassCallCount := 0
	generateExecDemoAppInitCommandsPass = func(project *model.CGProject, selectedTemplates *model.SelectedTemplates) {
		generateExecDemoAppInitCommandsPassCallCount++
	}
	// Execute test
	generateProject(project, config)
	// Assert
	assert.Zero(t, clearScreenCallCount)
	assert.Equal(t, 1, generateChooseFrontendsPassCallCount)
	assert.Equal(t, 1, generateChooseBackendsPassCallCount)
	assert.Equal(t, 1, generateChooseDatabasesPassCallCount)
	assert.Equal(t, 1, generateChooseDbAdminsPassCallCount)
	assert.Equal(t, 1, generateChooseProxiesPassCallCount)
	assert.Equal(t, 1, generateChooseTlsHelpersPassCallCount)
	assert.Equal(t, 1, generatePassCallCount)
	assert.Equal(t, 1, generateResolveDependencyGroupsPassCallCount)
	assert.Equal(t, 1, generateSecretsPassCallCount)
	assert.Equal(t, 1, generateAddProfilesPassCallCount)
	assert.Equal(t, 1, generateAddProxyNetworksCallCount)
	assert.Equal(t, 1, generateCopyVolumesPassCallCount)
	assert.Equal(t, 1, generateReplaceVarsInConfigFilesPassCallCount)
	assert.Equal(t, 1, generateExecServiceInitCommandsPassCallCount)
	assert.Equal(t, 1, generateExecDemoAppInitCommandsPassCallCount)
}
