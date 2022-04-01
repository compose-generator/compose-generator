/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
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
	if isDevVersion() { // Dev
		return "../templates"
	}
	if isDockerizedEnvironment() { // Docker
		return "/cg/templates"
	}
	if isLinux() { // Linux
		return "/usr/lib/compose-generator/templates"
	}
	// Windows
	path, err := executable()
	if err != nil {
		ErrorLogger.Println("Cannot retrieve path of executable: " + err.Error())
		logError("Cannot retrieve path of executable", true)
		return ""
	}
	path = filepath.ToSlash(path)
	return path[:strings.LastIndex(path, "/")] + "/templates"
}

// GetPredefinedServicesPath returns the path to the predefined services directory
func GetPredefinedServicesPath() string {
	if isDevVersion() { // Dev
		return "../predefined-services"
	}
	if isDockerizedEnvironment() { // Docker
		return "/cg/predefined-services"
	}
	if isLinux() { // Linux
		return "/usr/lib/compose-generator/predefined-services"
	}
	// Windows
	path, err := executable()
	if err != nil {
		ErrorLogger.Println("Cannot retrieve path of executable: " + err.Error())
		logError("Cannot retrieve path of executable", true)
		return ""
	}
	path = filepath.ToSlash(path)
	return path[:strings.LastIndex(path, "/")] + "/predefined-services"
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
	if isDevVersion() { // Dev
		return "../log"
	}
	if isDockerizedEnvironment() { // Docker
		return "/cg/log"
	}
	if isLinux() { // Linux
		return "/var/log/compose-generator"
	}
	// Windows
	cacheDir, err := userCacheDir()
	if err != nil {
		ErrorLogger.Println("Cannot find Windows cache dir: " + err.Error())
		logError("Cannot find Windows cache dir", true)
		return ""
	}
	cacheDir = filepath.ToSlash(cacheDir)
	return cacheDir + "/ComposeGenerator/log"
}

func getCComCompilerPath() string {
	if isWindows() { // Windows
		return "ccomc"
	}
	return "/usr/lib/ccom/ccomc" // Linux + Docker
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

func getUsedPortsOfRunningServices() []int {
	// Obtain Docker client
	client, err := newClientWithOpts(client.FromEnv)
	if err != nil {
		ErrorLogger.Println("Docker client initialization failed: " + err.Error())
		logError("Could not intanciate Docker client. Please check your Docker installation", true)
		return []int{}
	}
	// Get runnings containers
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		WarningLogger.Println("Could not obtain container list: " + err.Error())
		logWarning("Could not obtain container list. Please check your Docker installation")
		return []int{}
	}
	// Get ports from running services
	ports := []int{}
	for _, container := range containers {
		for _, port := range container.Ports {
			ports = append(ports, int(port.PublicPort))
		}
	}
	return ports
}

// IsLinux checks if the os in Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsWindows checks if the os is Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
