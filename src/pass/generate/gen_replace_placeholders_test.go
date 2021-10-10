/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package pass

import (
	"compose-generator/model"
	"errors"
	"io/fs"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------- GenerateReplaceVarsInConfigFiles -------------------------------------------------------

func TestGenerateReplacePlaceholdersInConfigFiles(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &types.Project{
			WorkingDir: "./work-dir/",
		},
		Vars: map[string]string{
			"NODE_VERSION": "3.14.1",
			"NODE_PORT":    "3000",
		},
	}
	selectedTemplates := &model.SelectedTemplates{
		BackendServices: []model.PredefinedTemplateConfig{
			{
				Label: "Node.js",
				Files: []model.File{
					{
						Path: "Dockerfile",
						Type: model.FileTypeConfig,
					},
					{
						Path: "environment.env",
						Type: model.FileTypeEnv,
					},
					{
						Path: "test/another-config-file.conf",
						Type: model.FileTypeConfig,
					},
					{
						Path: "test/non-existent.conf",
						Type: model.FileTypeConfig,
					},
				},
			},
		},
		DbAdminServices: []model.PredefinedTemplateConfig{
			{
				Label: "PhpMyAdmin",
			},
		},
	}
	// Mock functions
	startProcessCallCount := 0
	startProcess = func(text string) (s *spinner.Spinner) {
		startProcessCallCount++
		if startProcessCallCount == 1 {
			assert.Equal(t, "Applying custom config for Node.js ...", text)
		} else {
			assert.Equal(t, "Applying custom config for PhpMyAdmin ...", text)
		}
		return nil
	}
	stopProcess = func(s *spinner.Spinner) {
		assert.Nil(t, s)
	}
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		switch fileExistsCallCount {
		case 1:
			assert.Equal(t, "work-dir/Dockerfile", path)
		case 2:
			assert.Equal(t, "work-dir/test/another-config-file.conf", path)
		case 3:
			assert.Equal(t, "work-dir/test/non-existent.conf", path)
			return false
		}
		return true
	}
	replaceVarsInFileCallCount := 0
	replaceVarsInFileMockable = func(filePath string, vars map[string]string) {
		replaceVarsInFileCallCount++
		if replaceVarsInFileCallCount == 1 {
			assert.Equal(t, "work-dir/Dockerfile", filePath)
		} else {
			assert.Equal(t, "work-dir/test/another-config-file.conf", filePath)
		}
		assert.EqualValues(t, map[string]string{
			"NODE_VERSION": "3.14.1",
			"NODE_PORT":    "3000",
		}, vars)
	}
	evaluateConditionalSectionsCallCount := 0
	evaluateConditionalSections = func(filePath string, selected *model.SelectedTemplates, vars map[string]string) {
		evaluateConditionalSectionsCallCount++
		if evaluateConditionalSectionsCallCount == 1 {
			assert.Equal(t, "work-dir/Dockerfile", filePath)
		} else {
			assert.Equal(t, "work-dir/test/another-config-file.conf", filePath)
		}
		assert.EqualValues(t, map[string]string{
			"NODE_VERSION": "3.14.1",
			"NODE_PORT":    "3000",
		}, vars)
	}
	// Execute test
	GenerateReplacePlaceholdersInConfigFiles(project, selectedTemplates)
	// Assert
	assert.Equal(t, 2, startProcessCallCount)
	assert.Equal(t, 2, replaceVarsInFileCallCount)
	assert.Equal(t, 2, evaluateConditionalSectionsCallCount)
}

// --------------------------------------------------------------- ReplaceVarsInFile ---------------------------------------------------------------

func TestReplaceVarsInFile1(t *testing.T) {
	// Test data
	filePath := "./work-dir/test/Dockerfile"
	vars := map[string]string{
		"NODE_VERSION": "3.14.1",
		"NODE_PORT":    "3000",
	}
	// Mock functions
	readFile = func(filename string) ([]byte, error) {
		assert.Equal(t, filePath, filename)
		return []byte("Test with ${{NODE_PORT}} and ${{NODE_VERSION}}"), nil
	}
	writeFile = func(filename string, data []byte, perm fs.FileMode) error {
		assert.Equal(t, filePath, filename)
		assert.Equal(t, []byte("Test with 3000 and 3.14.1"), data)
		assert.Equal(t, fs.FileMode(0600), perm)
		return nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	replaceVarsInFile(filePath, vars)
}

func TestReplaceVarsInFile2(t *testing.T) {
	// Test data
	filePath := "./work-dir/test/Dockerfile"
	vars := map[string]string{
		"NODE_VERSION": "3.14.1",
		"NODE_PORT":    "3000",
	}
	// Mock functions
	readFile = func(filename string) ([]byte, error) {
		assert.Equal(t, filePath, filename)
		return nil, errors.New("Error message")
	}
	writeFile = func(filename string, data []byte, perm fs.FileMode) error {
		assert.Fail(t, "Unexpected call of writeFile")
		return nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Unable to read config file './work-dir/test/Dockerfile'", description)
		assert.Equal(t, "Error message", err.Error())
		assert.False(t, exit)
	}
	// Execute test
	replaceVarsInFile(filePath, vars)
}

func TestReplaceVarsInFile3(t *testing.T) {
	// Test data
	filePath := "./work-dir/test/Dockerfile"
	vars := map[string]string{
		"NODE_VERSION": "3.14.1",
		"NODE_PORT":    "3000",
	}
	// Mock functions
	readFile = func(filename string) ([]byte, error) {
		assert.Equal(t, filePath, filename)
		return []byte("Test with ${{NODE_PORT}} and ${{NODE_VERSION}}"), nil
	}
	writeFile = func(filename string, data []byte, perm fs.FileMode) error {
		assert.Equal(t, filePath, filename)
		assert.Equal(t, []byte("Test with 3000 and 3.14.1"), data)
		assert.Equal(t, fs.FileMode(0600), perm)
		return errors.New("Error message")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Unable to write config file './work-dir/test/Dockerfile' back to the disk", description)
		assert.Equal(t, "Error message", err.Error())
		assert.False(t, exit)
	}
	// Execute test
	replaceVarsInFile(filePath, vars)
}
