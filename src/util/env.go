/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"context"
	"os"
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
func GetDockerVersion() (string, error) {
	// Initialize Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		ErrorLogger.Println("Could not intanciate Docker client: " + err.Error())
		logError("Could not intanciate Docker client. Please check your Docker installation", true)
		return "", err
	}
	serverVersion, err := client.ServerVersion(context.Background())
	if err != nil {
		ErrorLogger.Println("Could not obtain Docker engine version: " + err.Error())
		logError("Could not obtain Docker engine version. Please check your Docker installation", true)
		return "", err
	}
	return serverVersion.Version, nil
}

// GetCustomTemplatesPath returns the path to the custom templates directory
func GetCustomTemplatesPath() string {
	if fileExists("/usr/bin/compose-generator") {
		return "/usr/lib/compose-generator/templates" // Linux
	}
	filename, err := executable()
	if err != nil {
		ErrorLogger.Println("Cannot retrieve path of executable: " + err.Error())
		logError("Cannot retrieve path of executable", true)
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
	if fileExists("/usr/bin/compose-generator") {
		return "/usr/lib/compose-generator/predefined-services" // Linux
	}
	filename, err := executable()
	if err != nil {
		ErrorLogger.Println("Cannot retrieve path of executable: " + err.Error())
		logError("Cannot retrieve path of executable", true)
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
		ErrorLogger.Println("Docker client initialization failed: " + err.Error())
		logError("Could not intanciate Docker client. Please check your Docker installation", true)
		return false
	}
	images, err := imageList(client, context.Background(), types.ImageListOptions{})
	if err != nil {
		ErrorLogger.Println("Could not load Docker images: " + err.Error())
		logError("Could not load Docker images", true)
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

func getLogfilesPath() string {
	if fileExists("/usr/bin/compose-generator") {
		return "/usr/lib/compose-generator/log" // Linux
	}
	filename, err := executable()
	if err != nil {
		logError("Cannot retrieve path of executable", true)
	}
	filename = filepath.ToSlash(filename)
	filename = filename[:strings.LastIndex(filename, "/")]
	if fileExists(filename + "/log") {
		return filename + "/log" // Windows + Docker
	}
	return "../log" // Dev
}

func getOuterVolumePathOnDockerizedEnvironment() string {
	// Obtain Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		ErrorLogger.Println("Docker client initialization failed: " + err.Error())
		logError("Could not intanciate Docker client. Please check your Docker installation", true)
		return ""
	}
	// Get hostname as it is the container id
	hostname, err := os.Hostname()
	if err != nil {
		ErrorLogger.Println("Could not obtain hostname: " + err.Error())
		logError("Could not obtain the hostname of the container", true)
		return ""
	}
	// Get container details
	container, err := client.ContainerInspect(context.Background(), hostname)
	if err != nil {
		ErrorLogger.Println("Could not obtain container details: " + err.Error())
		logError("Could not inspect the container", true)
		return ""
	}
	// Search for volume which is mounted to /cg/out
	for _, mount := range container.Mounts {
		if mount.Type == spec.VolumeTypeBind && mount.Destination == "/cg/out" {
			return mount.Source
		}
	}
	// Volume not found => error
	ErrorLogger.Println("Could not find volume on host")
	logError("Could not find a volume that is mounted to /cg/out", true)
	return ""
}
