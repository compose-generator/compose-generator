package pass

import (
	"compose-generator/model"
	"testing"

	"github.com/briandowns/spinner"
	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------- GenerateReplaceVarsInConfigFiles -------------------------------------------------------

func TestGenerateReplaceVarsInConfigFiles(t *testing.T) {
	// Test data
	project := &model.CGProject{
		Composition: &types.Project{
			WorkingDir: "./work-dir",
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
	replaceVarsInFileCallCount := 0
	replaceVarsInFileMockable = func(filePath string, vars map[string]string) {
		replaceVarsInFileCallCount++
		if replaceVarsInFileCallCount == 1 {
			assert.Equal(t, "./work-dir/Dockerfile", filePath)
		} else {
			assert.Equal(t, "./work-dir/test/another-config-file.conf", filePath)
		}
		assert.EqualValues(t, map[string]string{
			"NODE_VERSION": "3.14.1",
			"NODE_PORT":    "3000",
		}, vars)
	}
	// Execute test
	GenerateReplaceVarsInConfigFiles(project, selectedTemplates)
	// Assert
	assert.Equal(t, 2, startProcessCallCount)
	assert.Equal(t, 2, replaceVarsInFileCallCount)
}

// --------------------------------------------------------------- ReplaceVarsInFile ---------------------------------------------------------------

func TestReplaceVarsInFile(t *testing.T) {
	// Test data

	// Mock functions

	// Execute test

	// Assert
}
