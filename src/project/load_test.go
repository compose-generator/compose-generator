package project

import (
	"compose-generator/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------ LoadProject ------------------------------------------------------------------

func TestLoadProject(t *testing.T) {
	// Test data
	expectedProject := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
			WithReadme:    false,
		},
		GitignorePatterns: []string{},
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
	}
	options := LoadOptions{
		ComposeFileName: "docker.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, "./context/.gitignore", path)
			return true
		}
		assert.Equal(t, "./context/README.md", path)
		return false
	}
	loadComposeFileCallCount := 0
	loadComposeFileMockable = func(project *model.CGProject, opt LoadOptions) {
		loadComposeFileCallCount++
		assert.Equal(t, options, opt)
	}
	loadGitignoreFileCallCount := 0
	loadGitignoreFileMockable = func(project *model.CGProject, opt LoadOptions) {
		loadGitignoreFileCallCount++
		assert.Equal(t, options, opt)
	}
	loadCGFileCallCount := 0
	loadCGFileMockable = func(metadata *model.CGProjectMetadata, opt LoadOptions) {
		loadCGFileCallCount++
		assert.Equal(t, options, opt)
	}
	// Execute test
	result := LoadProject(LoadFromComposeFile("docker.yml"), LoadFromDir("./context"))
	// Assert
	assert.Equal(t, expectedProject, result)
	assert.Equal(t, 1, loadComposeFileCallCount)
	assert.Equal(t, 1, loadGitignoreFileCallCount)
	assert.Equal(t, 1, loadCGFileCallCount)
	assert.Equal(t, 2, fileExistsCallCount)
}

// -------------------------------------------------------------- LoadProjectMetadata --------------------------------------------------------------
