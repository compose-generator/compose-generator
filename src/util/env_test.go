package util

import (
	"context"
	"errors"
	"os/user"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------- IsDockerizedEnvironment -----------------------------------------------------------

func TestIsDockerizedEnvironment1(t *testing.T) {
	// Mock functions
	getEnv = func(key string) string {
		assert.Equal(t, "COMPOSE_GENERATOR_DOCKERIZED", key)
		return "1"
	}
	// Execute test
	result := IsDockerizedEnvironment()
	// Assert
	assert.True(t, result)
}

func TestIsDockerizedEnvironment2(t *testing.T) {
	// Mock functions
	getEnv = func(key string) string {
		assert.Equal(t, "COMPOSE_GENERATOR_DOCKERIZED", key)
		return ""
	}
	// Execute test
	result := IsDockerizedEnvironment()
	// Assert
	assert.False(t, result)
}

// ------------------------------------------------------------- GetCustomTemplatesPath ------------------------------------------------------------

func TestGetCustomTemplatesPath1(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/templates"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, pathLinux, path)
		return true
	}
	// Execute test
	result := GetCustomTemplatesPath()
	// Assert
	assert.Equal(t, pathLinux, result)
}

func TestGetCustomTemplatesPath2(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/templates"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/templates"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return true
	}
	executable = func() (string, error) {
		return pathExecutable, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := GetCustomTemplatesPath()
	// Assert
	assert.Equal(t, pathWindowsDocker, result)
}

func TestGetCustomTemplatesPath3(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/templates"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/templates"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return true
	}
	executable = func() (string, error) {
		return pathExecutable, errors.New("Test error")
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Cannot retrieve path of executable", description)
		assert.NotNil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := GetCustomTemplatesPath()
	// Assert
	assert.Equal(t, pathWindowsDocker, result)
	assert.Equal(t, 1, printErrorCallCount)
}

func TestGetCustomTemplatesPath4(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/templates"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/templates"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	pathDev := "../templates"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return false
	}
	executable = func() (string, error) {
		return pathExecutable, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := GetCustomTemplatesPath()
	// Assert
	assert.Equal(t, pathDev, result)
}

// ------------------------------------------------------------------ GetUsername ------------------------------------------------------------------

func TestGetUsername1(t *testing.T) {
	// Test data
	username := "Marc"
	// Mock functions
	currentUser = func() (*user.User, error) {
		user := &user.User{
			Username: username,
		}
		return user, nil
	}
	// Execute test
	result := GetUsername()
	// Assert
	assert.Equal(t, username, result)
}

func TestGetUsername2(t *testing.T) {
	// Mock functions
	currentUser = func() (*user.User, error) {
		return nil, errors.New("Error")
	}
	// Execute test
	result := GetUsername()
	// Assert
	assert.Equal(t, "unknown", result)
}

// ----------------------------------------------------------- GetPredefinedServicesPath -----------------------------------------------------------

func TestGetPredefinedServicesPath1(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/predefined-services"
	// Mock functions
	fileExists = func(path string) bool {
		assert.Equal(t, pathLinux, path)
		return true
	}
	// Execute test
	result := GetPredefinedServicesPath()
	// Assert
	assert.Equal(t, pathLinux, result)
}

func TestGetPredefinedServicesPath2(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/predefined-services"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/predefined-services"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return true
	}
	executable = func() (string, error) {
		return pathExecutable, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := GetPredefinedServicesPath()
	// Assert
	assert.Equal(t, pathWindowsDocker, result)
}

func TestGetPredefinedServicesPath3(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/predefined-services"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/predefined-services"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return true
	}
	executable = func() (string, error) {
		return pathExecutable, errors.New("Test error")
	}
	printErrorCallCount := 0
	printError = func(description string, err error, exit bool) {
		printErrorCallCount++
		assert.Equal(t, "Cannot retrieve path of executable", description)
		assert.NotNil(t, err)
		assert.True(t, exit)
	}
	// Execute test
	result := GetPredefinedServicesPath()
	// Assert
	assert.Equal(t, pathWindowsDocker, result)
	assert.Equal(t, 1, printErrorCallCount)
}

func TestGetPredefinedServicesPath4(t *testing.T) {
	// Test data
	pathLinux := "/usr/lib/compose-generator/predefined-services"
	pathWindowsDocker := "/usr/bin/compose-generator/test/path/predefined-services"
	pathExecutable := "/usr/bin/compose-generator/test/path/dir"
	pathDev := "../predefined-services"
	// Mock functions
	fileExistsCallCount := 0
	fileExists = func(path string) bool {
		fileExistsCallCount++
		if fileExistsCallCount == 1 {
			assert.Equal(t, pathLinux, path)
			return false
		}
		assert.Equal(t, pathWindowsDocker, path)
		return false
	}
	executable = func() (string, error) {
		return pathExecutable, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := GetPredefinedServicesPath()
	// Assert
	assert.Equal(t, pathDev, result)
}

// --------------------------------------------------------------- IsToolboxPresent ----------------------------------------------------------------

func TestIsToolboxPresent1(t *testing.T) {
	// Test data
	version := "1.0.0-rc2"
	// Mock functions
	getToolboxImageVersionMockable = func() string {
		return version
	}
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		assert.Equal(t, 1, len(ops))
		return nil, nil
	}
	imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
		assert.Nil(t, cli)
		assert.Equal(t, context.Background(), ctx)
		return []types.ImageSummary{
			{
				RepoTags: []string{"hello-world"},
			},
			{
				RepoTags: []string{"chillibits/compose-generator-toolbox:" + version},
			},
		}, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := IsToolboxPresent()
	// Assert
	assert.True(t, result)
}

func TestIsToolboxPresent2(t *testing.T) {
	// Test data
	version := "dev"
	// Mock functions
	getToolboxImageVersionMockable = func() string {
		return version
	}
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		assert.Equal(t, 1, len(ops))
		return nil, errors.New("Error message")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Could not intanciate Docker client. Please check your Docker installation", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	result := IsToolboxPresent()
	// Assert
	assert.False(t, result)
}

func TestIsToolboxPresent3(t *testing.T) {
	// Test data
	version := "0.9.0"
	// Mock functions
	getToolboxImageVersionMockable = func() string {
		return version
	}
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		assert.Equal(t, 1, len(ops))
		return nil, nil
	}
	imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
		assert.Nil(t, cli)
		assert.Equal(t, context.Background(), ctx)
		return []types.ImageSummary{
			{
				RepoTags: []string{"hello-world"},
			},
			{
				RepoTags: []string{"chillibits/compose-generator-toolbox:" + version},
			},
		}, errors.New("Error message")
	}
	printError = func(description string, err error, exit bool) {
		assert.Equal(t, "Could not load Docker images", description)
		assert.Equal(t, "Error message", err.Error())
		assert.True(t, exit)
	}
	// Execute test
	result := IsToolboxPresent()
	// Assert
	assert.False(t, result)
}

func TestIsToolboxPresent4(t *testing.T) {
	// Test data
	version := "0.9.0"
	// Mock functions
	getToolboxImageVersionMockable = func() string {
		return version
	}
	newClientWithOpts = func(ops ...client.Opt) (*client.Client, error) {
		assert.Equal(t, 1, len(ops))
		return nil, nil
	}
	imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
		assert.Nil(t, cli)
		assert.Equal(t, context.Background(), ctx)
		return []types.ImageSummary{
			{
				RepoTags: []string{"hello-world"},
			},
			{
				RepoTags: []string{"chillibits/spice:0.4.0"},
			},
		}, nil
	}
	printError = func(description string, err error, exit bool) {
		assert.Fail(t, "Unexpected call of printError")
	}
	// Execute test
	result := IsToolboxPresent()
	// Assert
	assert.False(t, result)
}
