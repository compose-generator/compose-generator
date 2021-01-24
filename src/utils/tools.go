package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ExecuteSafetyFileChecks() {
	isNotSafe := true
	if IsDockerized() {
		isNotSafe = fileExists("/compose-generator/out/docker-compose.yml") || fileExists("/compose-generator/out/environment.env")
	} else {
		isNotSafe = fileExists("docker-compose.yml") || fileExists("environment.env")
	}
	if isNotSafe {
		color.Red("Warning: docker-compose.yml or environment.env already exists. By continuing, you might overwrite those files.")
		result := YesNoQuestion("Do you want to continue?", true)
		if !result {
			os.Exit(0)
		}
		fmt.Println()
	}
}

func IsDockerized() bool {
	return os.Getenv("COMPOSE_GENERATOR_DOCKERIZED") == "1"
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
