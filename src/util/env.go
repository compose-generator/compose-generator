/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	spec "github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var getToolboxImageVersionMockable = getToolboxImageVersion

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// IsDockerizedEnvironment checks if Compose Generator runs within a dockerized environment
func IsDockerizedEnvironment() bool {
	return getEnv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

// IsCIEnvironment checks if Compose Generator runs within a CI environment
func IsCIEnvironment() bool {
	return getEnv("COMPOSE_GENERATOR_CI") == "1"
}

// GetUsername returns the username of the current username. If it is not determinable it returns "unknown"
func GetUsername() string {
	if user, err := currentUser(); err == nil {
		return user.Username
	}
	return "unknown"
}

// GetDockerVersion retrives and returns the version of the installed Docker instance
func GetDockerVersion() string {
	cmd := exec.Command("docker", "-v")
	dockerVersion, err := cmd.CombinedOutput()
	if err != nil {
		printError("Could not read Docker version", err, true)
	}
	return strings.TrimRight(string(dockerVersion), "\r\n")
}

// GetCustomTemplatesPath returns the path to the custom templates directory
func GetCustomTemplatesPath() string {
	if fileExists("/usr/lib/compose-generator/templates") {
		return "/usr/lib/compose-generator/templates" // Linux
	}
	filename, err := executable()
	if err != nil {
		printError("Cannot retrieve path of executable", err, true)
	}
	filename = filepath.ToSlash(filename)
	filename = filename[:strings.LastIndex(filename, "/")]
	if fileExists(filename + "/templates") {
		return filename + "/templates" // Windows + Docker
	}
	return "../templates" // Dev
}

// GetPredefinedServicesPath returns the path to the predefined services directory
func GetPredefinedServicesPath() string {
	if fileExists("/usr/lib/compose-generator/predefined-services") {
		return "/usr/lib/compose-generator/predefined-services" // Linux
	}
	filename, err := executable()
	if err != nil {
		printError("Cannot retrieve path of executable", err, true)
	}
	filename = filepath.ToSlash(filename)
	filename = filename[:strings.LastIndex(filename, "/")]
	if fileExists(filename + "/predefined-services") {
		return filename + "/predefined-services" // Windows + Docker
	}
	return "../predefined-services" // Dev
}

// IsToolboxPresent checks if the Compose Generator toolbox image is present on the Docker host
func IsToolboxPresent() bool {
	// Check if Toolbox is present
	toolboxTag := "chillibits/compose-generator-toolbox:" + getToolboxImageVersionMockable()
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		printError("Could not intanciate Docker client. Please check your Docker installation", err, true)
		return false
	}
	images, err := imageList(client, context.Background(), types.ImageListOptions{})
	if err != nil {
		printError("Could not load Docker images", err, true)
		return false
	}
	for _, image := range images {
		if SliceContainsString(image.RepoTags, toolboxTag) {
			return true
		}
	}
	return false
}

// IsDockerRunning checks if Docker is running
func IsDockerRunning() bool {
	cmd := executeCommand("docker", "info")
	output, err := getCommandOutput(cmd)
	return err == nil && !strings.Contains(string(output), "Server:\nERROR: error during connect")
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getOuterVolumePathOnDockerizedEnvironment() string {
	// Obtain Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		printError("Could not intanciate Docker client. Please check your Docker installation", err, true)
		return ""
	}
	// Get hostname as it is the container id
	hostname, err := os.Hostname()
	if err != nil {
		printError("Could not obtain the hostname of the container", err, true)
		return ""
	}
	// Get container details
	container, err := client.ContainerInspect(context.Background(), hostname)
	if err != nil {
		printError("Could not inspect the container", err, true)
		return ""
	}
	// Search for volume which is mounted to /cg/out
	for _, mount := range container.Mounts {
		if mount.Type == spec.VolumeTypeBind && mount.Destination == "/cg/out" {
			return mount.Source
		}
	}
	// Volume not found => error
	printError("Could not find a volume that is mounted to /cg/out", err, true)
	return ""
}
