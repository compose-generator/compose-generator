package project

import (
	"compose-generator/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------ DeleteReadme -----------------------------------------------------------------

func TestDeleteReadme1(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
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
	deleteReadme(project, DeleteOptions{
		WorkingDir: "./context/",
	})
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
	// Mock functions
	remove = func(name string) error {
		assert.Fail(t, "Unexpected call of remove")
		return nil
	}
	// Execute test
	deleteReadme(project, DeleteOptions{
		WorkingDir: "./context/",
	})
}

func TestDeleteReadme3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithReadme: true,
		},
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
	deleteReadme(project, DeleteOptions{
		WorkingDir: "./context/",
	})
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
	deleteGitignore(project, DeleteOptions{
		WorkingDir: "./context/",
	})
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
	// Mock functions
	remove = func(name string) error {
		assert.Fail(t, "Unexpected call of remove")
		return nil
	}
	// Execute test
	deleteGitignore(project, DeleteOptions{
		WorkingDir: "./context/",
	})
}

func TestDeleteGitignore3(t *testing.T) {
	// Test data
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: true,
		},
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
	deleteGitignore(project, DeleteOptions{
		WorkingDir: "./context/",
	})
	// Assert
	assert.Equal(t, 1, removeCallCount)
}

// ----------------------------------------------------------------- DeleteVolumes -----------------------------------------------------------------
// --------------------------------------------------------------- DeleteComposeFiles --------------------------------------------------------------
