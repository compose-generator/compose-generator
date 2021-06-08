package util

import (
	"os"
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

// IsDockerizedEnvironment checks if Compose Generator runs within a dockerized environment
func IsDockerizedEnvironment() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

// IsDevEnvironment checks if Compose Generator runs in a dev environment
func IsDevEnvironment() bool {
	return os.Getenv("COMPOSE_GENERATOR_DEV") == "1"
}

// GetCustomTemplatesPath returns the path to the custom templates directory
func GetCustomTemplatesPath() string {
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
