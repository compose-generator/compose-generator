package parser

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/compose-generator/diu/cli"
	"github.com/compose-generator/diu/model"
)

// ParseDockerManifest retrieves the manifest of a remote Docker image and bundles it to an object
func ParseDockerManifest(imageName string) (manifest model.DockerManifest, err error) {
	manifestJSON := cli.ExecuteAndWaitWithOutput("docker", "manifest", "inspect", "-v", imageName)
	if strings.HasPrefix(manifestJSON, "[") {
		var manifestArray []model.DockerManifest
		json.Unmarshal([]byte(manifestJSON), &manifestArray)
		manifest = manifestArray[0]
		return
	} else if strings.HasPrefix(manifestJSON, "{") {
		json.Unmarshal([]byte(manifestJSON), &manifest)
		return
	}
	err = errors.New("Could not parse manifest")
	return
}

// ParseDockerVolumes retrieves all existing volumes of the local docker instance and bundles them to objects
func ParseDockerVolumes() (volumes []model.DockerVolume, err error) {
	// Get volume names
	volumeNameString := cli.ExecuteAndWaitWithOutput("docker", "volume", "ls", "-q")
	volumeNames := strings.Split(volumeNameString, "\n")
	if len(volumeNames) == 0 {
		return
	}
	// Parse JSON
	queryCmd := []string{"docker", "volume", "inspect"}
	queryCmd = append(queryCmd, volumeNames...)
	volumesJSON := cli.ExecuteAndWaitWithOutput(queryCmd...)
	json.Unmarshal([]byte(volumesJSON), &volumes)
	return
}

// ParseDockerNetworks retrieves all existing networks of the local docker instance and bundles them to objects
func ParseDockerNetworks() (networks []model.DockerNetwork, err error) {
	// Get network names
	networkNameString := cli.ExecuteAndWaitWithOutput("docker", "network", "ls", "-q")
	networkNames := strings.Split(networkNameString, "\n")
	if len(networkNames) == 0 {
		return
	}
	// Parse JSON
	queryCmd := []string{"docker", "network", "inspect", "-v"}
	queryCmd = append(queryCmd, networkNames...)
	volumesJSON := cli.ExecuteAndWaitWithOutput(queryCmd...)
	json.Unmarshal([]byte(volumesJSON), &networks)
	return
}
