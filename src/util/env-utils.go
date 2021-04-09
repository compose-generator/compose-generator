package util

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kardianos/osext"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// PrintSafetyWarning checks if commonly used files are already existing and warns the user about it
func PrintSafetyWarning(existingCount int) {
	Pel()
	Warning(strconv.Itoa(existingCount) + " output files already exist. By continuing, those files will be overwritten!")
	result := YesNoQuestion("Do you want to continue?", true)
	if !result {
		os.Exit(0)
	}
	Pel()
}

// IsDockerized checks if Compose Generator runs within a dockerized environment
func IsDockerized() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

// GetTemplatesPath returns the path to the custom templates directory
func GetTemplatesPath() string {
	if FileExists("/usr/lib/compose-generator/templates") {
		return "/usr/lib/compose-generator/templates" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/templates") {
		return filename + "/templates" // Windows + Docker
	}
	return "../templates" // Dev
}

// GetPredefinedServicesPath returns the path to the predefined services directory
func GetPredefinedServicesPath() string {
	if FileExists("/usr/lib/compose-generator/predefined-services") {
		return "/usr/lib/compose-generator/predefined-services" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/predefined-services") {
		return filename + "/predefined-services" // Windows + Docker
	}
	return "../predefined-services" // Dev
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getToolboxPath() string {
	if FileExists("/usr/lib/compose-generator/toolbox") {
		return "/usr/lib/compose-generator/toolbox" // Linux
	}
	filename, _ := osext.Executable()
	filename = strings.ReplaceAll(filename, "\\", "/")
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/toolbox") {
		return filename + "/toolbox" // Windows + Docker
	}
	return "../toolbox" // Dev
}

func ensureToolbox() {
	// Check if toolbox image exists locally
	imageInspect := ExecuteAndWaitWithOutput("docker", "image", "inspect", "compose-generator-toolbox")
	if !strings.HasPrefix(imageInspect, "[]") { // Image exists locally -> return
		return
	}
	if FileExists(filepath.Join(getToolboxPath(), "toolbox.img")) { // Image exists as file
		// Load iamge
		ExecuteAndWait("docker", "load", "-i", filepath.Join(getToolboxPath(), "toolbox.img"))
	} else { // Image has to be built
		// Build docker image
		ExecuteAndWait("docker", "build", "-t", "compose-generator-toolbox", getToolboxPath())
		// Save docker image as file
		ExecuteAndWait("docker", "save", "-o", filepath.Join(getToolboxPath(), "toolbox.img"), "compose-generator-toolbox")
	}
}
