package cmd

import (
	"github.com/urfave/cli/v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(_ *cli.Context) error {
	// Check if Compose Generator runs in dockerized environment
	if isDockerizedEnvironment() {
		printError("You are currently using the dockerized version of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", nil, true)
	}

	// Execute passes
	installDockerPass()

	// Check if installation was successful
	if commandExists("docker") {
		pel()
		dockerVersion := getDockerVersion()
		printSuccessMessage("Congrats! You have installed " + dockerVersion + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
	} else {
		printError("An error occurred while installing Docker", nil, true)
	}
	return nil
}
