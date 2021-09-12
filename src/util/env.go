package util

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// IsDockerizedEnvironment checks if Compose Generator runs within a dockerized environment
func IsDockerizedEnvironment() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

// GetUsername returns the username of the current username. If it is not determinable it returns "unknown"
func GetUsername() string {
	if user, err := user.Current(); err == nil {
		return user.Username
	}
	return "unknown"
}

// GetCustomTemplatesPath returns the path to the custom templates directory
func GetCustomTemplatesPath() string {
	if FileExists("/usr/lib/compose-generator/templates") {
		return "/usr/lib/compose-generator/templates" // Linux
	}
	filename, _ := osext.Executable()
	filename = filepath.ToSlash(filename)
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
	filename = filepath.ToSlash(filename)
	filename = filename[:strings.LastIndex(filename, "/")]
	if FileExists(filename + "/predefined-services") {
		return filename + "/predefined-services" // Windows + Docker
	}
	return "../predefined-services" // Dev
}
