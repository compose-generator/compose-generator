package project

import (
	"compose-generator/model"
	"errors"
	"testing"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------- DeleteProject -----------------------------------------------------------------

func TestDeleteProject(t *testing.T) {
	// Test data
	project := &model.CGProject{}
	expectedOpt := DeleteOptions{
		ComposeFileName: "docker.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	deleteReadmeCallCount := 0
	deleteReadmeMockable = func(proj *model.CGProject, opt DeleteOptions) {
		deleteReadmeCallCount++
		assert.Equal(t, project, proj)
		assert.Equal(t, expectedOpt, opt)
	}
	deleteEnvFileCallCount := 0
	deleteEnvFilesMockable = func(proj *model.CGProject, opt DeleteOptions) {
		deleteEnvFileCallCount++
		assert.Equal(t, project, proj)
		assert.Equal(t, expectedOpt, opt)
	}
	deleteGitignoreCallCount := 0
	deleteGitignoreMockable = func(proj *model.CGProject, opt DeleteOptions) {
		deleteGitignoreCallCount++
		assert.Equal(t, project, proj)
		assert.Equal(t, expectedOpt, opt)
	}
	deleteVolumesCallCount := 0
	deleteVolumesMockable = func(proj *model.CGProject, opt DeleteOptions) {
		deleteVolumesCallCount++
		assert.Equal(t, project, proj)
		assert.Equal(t, expectedOpt, opt)
	}
	deleteComposeFileCallCount := 0
	deleteComposeFileMockable = func(proj *model.CGProject, opt DeleteOptions) {
		deleteComposeFileCallCount++
		assert.Equal(t, project, proj)
		assert.Equal(t, expectedOpt, opt)
	}
	// Execute test
	DeleteProject(project, DeleteComposeFileName("docker.yml"), DeleteWorkingDir("./context/"))
	// Assert
	assert.Equal(t, 1, deleteReadmeCallCount)
	assert.Equal(t, 1, deleteEnvFileCallCount)
	assert.Equal(t, 1, deleteGitignoreCallCount)
	assert.Equal(t, 1, deleteVolumesCallCount)
	assert.Equal(t, 1, deleteComposeFileCallCount)
}

// ------------------------------------------------------------------ DeleteReadme -----------------------------------------------------------------

func TestDeleteReadme1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/README.md", name)
		return nil
	}
	printWarning = func(description string) {
		assert.Fail(t, "Unexpected call of printWarning")
	}
	// Execute test
	deleteReadme(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

func TestDeleteReadme2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: false,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	remove = func(name string) error {
		assert.Fail(t, "Unexpected call of remove")
		return nil
	}
	// Execute test
	deleteReadme(project, options)
}

func TestDeleteReadme3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/README.md", name)
		return errors.New("Error message")
	}
	printWarning = func(description string) {
		assert.Equal(t, "File 'README.md' could not be deleted", description)
	}
	// Execute test
	deleteReadme(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

// ----------------------------------------------------------------- DeleteEnvFiles ----------------------------------------------------------------

// ---------------------------------------------------------------- DeleteGitignore ----------------------------------------------------------------

func TestDeleteGitignore1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/.gitignore", name)
		return nil
	}
	printWarning = func(description string) {
		assert.Fail(t, "Unexpected call of printWarning")
	}
	// Execute test
	deleteGitignore(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

func TestDeleteGitignore2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: false,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	remove = func(name string) error {
		assert.Fail(t, "Unexpected call of remove")
		return nil
	}
	// Execute test
	deleteGitignore(project, options)
}

func TestDeleteGitignore3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/.gitignore", name)
		return errors.New("Error message")
	}
	printWarning = func(description string) {
		assert.Equal(t, "File '.gitignore' could not be deleted", description)
	}
	// Execute test
	deleteGitignore(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

// ----------------------------------------------------------------- DeleteVolumes -----------------------------------------------------------------

func TestDeleteVolumes1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
		Composition: &spec.Project{
			Services: spec.Services{
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: "./volumes/frontend-wordpress",
						},
					},
				},
				{
					Volumes: []spec.ServiceVolumeConfig{
						{
							Type:   spec.VolumeTypeBind,
							Source: "./volumes/frontend-wordpress/config.yml",
						},
						{
							Type:   spec.VolumeTypeBind,
							Source: "./volumes/backend-gin",
						},
					},
				},
			},
		},
	}
	options := DeleteOptions{
		WorkingDir: "./context/",
	}
	// Mock functions
	normalizePaths = func(paths []string) []string {
		assert.EqualValues(t, []string{
			"./volumes/frontend-wordpress",
			"./volumes/frontend-wordpress/config.yml",
			"./volumes/backend-gin",
		}, paths)
		return []string{"./volumes/frontend-wordpress", "./volumes/backend-gin"}
	}
	removeAllCallCount := 0
	removeAll = func(path string) error {
		removeAllCallCount++
		if removeAllCallCount == 1 {
			assert.Equal(t, "./volumes/frontend-wordpress", path)
			return nil
		}
		assert.Equal(t, "./volumes/backend-gin", path)
		return errors.New("Error message")
	}
	printWarning = func(description string) {
		if removeAllCallCount == 1 {
			assert.Fail(t, "Unexpected call of printWarning")
		}
		assert.Equal(t, "Volume './volumes/backend-gin' could not be deleted", description)
	}
	// Execute test
	deleteVolumes(project, options)
	// Assert
	assert.Equal(t, 2, removeAllCallCount)
}

// --------------------------------------------------------------- DeleteComposeFiles --------------------------------------------------------------

func TestDeleteComposeFile1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
	}
	options := DeleteOptions{
		ComposeFileName: "compose.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/compose.yml", name)
		return nil
	}
	printWarning = func(description string) {
		assert.Fail(t, "Unexpected call of printWarning")
	}
	// Execute test
	deleteComposeFile(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

func TestDeleteComposeFile2(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
	}
	options := DeleteOptions{
		ComposeFileName: "compose.yml",
		WorkingDir:      "./context/",
	}
	// Mock functions
	removeCallCount := 0
	remove = func(name string) error {
		removeCallCount++
		assert.Equal(t, "./context/compose.yml", name)
		return errors.New("Error message")
	}
	printWarningCallCount := 0
	printWarning = func(description string) {
		printWarningCallCount++
		assert.Equal(t, "File 'compose.yml' could not be deleted", description)
	}
	// Execute test
	deleteComposeFile(project, options)
	// Assert
	assert.Equal(t, 1, removeCallCount)
	assert.Equal(t, 1, printWarningCallCount)
}
