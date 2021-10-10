/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package util

import (
	"context"
	"errors"
	"os/exec"
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

// ---------------------------------------------------------------- IsDockerRunning ----------------------------------------------------------------

func TestIsDockerRunning1(t *testing.T) {
	// Test data
	commandOutput := "Client:\nContext:    default\nDebug Mode: false\nPlugins:\nbuildx: Build with BuildKit (Docker Inc., 0.6.3+azure)\ncompose: Docker Compose (Docker Inc., 2.0.0)\n\n\nServer:\nContainers: 2\nRunning: 0\nPaused: 0\nStopped: 2\nImages: 1\nServer Version: 20.10.8+azure\nStorage Driver: overlay2\nBacking Filesystem: extfs\nSupports d_type: true\nNative Overlay Diff: false\nuserxattr: false\nLogging Driver: json-file\nCgroup Driver: cgroupfs\nCgroup Version: 1\nPlugins:\nVolume: local\nNetwork: bridge host ipvlan macvlan null overlay\nLog: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog\nSwarm: inactive\nRuntimes: io.containerd.runtime.v1.linux runc io.containerd.runc.v2\nDefault Runtime: runc\nInit Binary: docker-init\ncontainerd version: e25210fe30a0a703442421b0f60afac609f950a3\nrunc version: 4144b63817ebcc5b358fc2c8ef95f7cddd709aa7\ninit version:\nSecurity Options:\napparmor\nseccomp\nProfile: default\nKernel Version: 5.4.0-1059-azure\nOperating System: Ubuntu 20.04.3 LTS (containerized)\nOSType: linux\nArchitecture: x86_64\nCPUs: 4\nTotal Memory: 7.775GiB\nName: codespaces_09bc9d\nID: G4C5:7KMT:LQVT:QEDB:PW4I:DER3:OONZ:YKG5:TOMZ:BYX5:3Z2W:XGSV\nDocker Root Dir: /var/lib/docker\nDebug Mode: false\nUsername: codespacesdev\nRegistry: https://index.docker.io/v1/\nLabels:\nExperimental: false\nInsecure Registries:\n127.0.0.0/8\nLive Restore Enabled: false\n\n\nWARNING: No swap limit support"
	// Mock functions
	executeCommand = func(name string, arg ...string) *exec.Cmd {
		assert.Equal(t, "docker", name)
		assert.Equal(t, 1, len(arg))
		return nil
	}
	getCommandOutput = func(cmd *exec.Cmd) ([]byte, error) {
		assert.Nil(t, cmd)
		return []byte(commandOutput), nil
	}
	// Execute test
	result := IsDockerRunning()
	// Assert
	assert.True(t, result)
}

func TestIsDockerRunning2(t *testing.T) {
	// Test data
	commandOutput := "Client:\nContext:    default\nDebug Mode: false\nPlugins:\nbuildx: Build with BuildKit (Docker Inc., 0.6.3+azure)\ncompose: Docker Compose (Docker Inc., 2.0.0)\n\n\nServer:\nContainers: 2\nRunning: 0\nPaused: 0\nStopped: 2\nImages: 1\nServer Version: 20.10.8+azure\nStorage Driver: overlay2\nBacking Filesystem: extfs\nSupports d_type: true\nNative Overlay Diff: false\nuserxattr: false\nLogging Driver: json-file\nCgroup Driver: cgroupfs\nCgroup Version: 1\nPlugins:\nVolume: local\nNetwork: bridge host ipvlan macvlan null overlay\nLog: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog\nSwarm: inactive\nRuntimes: io.containerd.runtime.v1.linux runc io.containerd.runc.v2\nDefault Runtime: runc\nInit Binary: docker-init\ncontainerd version: e25210fe30a0a703442421b0f60afac609f950a3\nrunc version: 4144b63817ebcc5b358fc2c8ef95f7cddd709aa7\ninit version:\nSecurity Options:\napparmor\nseccomp\nProfile: default\nKernel Version: 5.4.0-1059-azure\nOperating System: Ubuntu 20.04.3 LTS (containerized)\nOSType: linux\nArchitecture: x86_64\nCPUs: 4\nTotal Memory: 7.775GiB\nName: codespaces_09bc9d\nID: G4C5:7KMT:LQVT:QEDB:PW4I:DER3:OONZ:YKG5:TOMZ:BYX5:3Z2W:XGSV\nDocker Root Dir: /var/lib/docker\nDebug Mode: false\nUsername: codespacesdev\nRegistry: https://index.docker.io/v1/\nLabels:\nExperimental: false\nInsecure Registries:\n127.0.0.0/8\nLive Restore Enabled: false\n\n\nWARNING: No swap limit support"
	// Mock functions
	executeCommand = func(name string, arg ...string) *exec.Cmd {
		assert.Equal(t, "docker", name)
		assert.Equal(t, 1, len(arg))
		return nil
	}
	getCommandOutput = func(cmd *exec.Cmd) ([]byte, error) {
		assert.Nil(t, cmd)
		return []byte(commandOutput), errors.New("Error message")
	}
	// Execute test
	result := IsDockerRunning()
	// Assert
	assert.False(t, result)
}

func TestIsDockerRunning3(t *testing.T) {
	// Test data
	commandOutput := "Client:\nContext:    default\nDebug Mode: false\nPlugins:\nbuildx: Build with BuildKit (Docker Inc., v0.6.3)\ncompose: Docker Compose (Docker Inc., v2.0.0)\nscan: Docker Scan (Docker Inc., v0.8.0)\n\n\nServer:\nERROR: error during connect: This error may indicate that the docker daemon is not running.: Get \"http://%2F%2F.%2Fpipe%2Fdocker_engine/v1.24/info\": open //./pipe/docker_engine: The system cannot find the file specified.\nerrors pretty printing info"
	// Mock functions
	executeCommand = func(name string, arg ...string) *exec.Cmd {
		assert.Equal(t, "docker", name)
		assert.Equal(t, 1, len(arg))
		return nil
	}
	getCommandOutput = func(cmd *exec.Cmd) ([]byte, error) {
		assert.Nil(t, cmd)
		return []byte(commandOutput), nil
	}
	// Execute test
	result := IsDockerRunning()
	// Assert
	assert.False(t, result)
}
