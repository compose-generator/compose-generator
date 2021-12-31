/*
Copyright © 2021-2022 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/util"

	"github.com/urfave/cli/v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(_ *cli.Context) error {
	util.InfoLogger.Println("Install command executed")

	// Check if Compose Generator runs in dockerized environment
	if isDockerizedEnvironment() {
		util.ErrorLogger.Println("Dockerized environment detected")
		logError("You are currently using the dockerized version of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", true)
	}

	// Execute passes
	installDockerPass()

	// Check if installation was successful
	if commandExists("docker") {
		if dockerVersion, err := getDockerVersion(); err == nil {
			pel()
			printSuccess("Congrats! You have installed Docker " + dockerVersion + ". You now can start by executing '$ compose-generator generate' to generate your compose file.")
			util.InfoLogger.Println("Installation successful")
			return nil
		}
	}
	util.ErrorLogger.Println("Installation failed")
	logError("An error occurred while installing Docker", true)
	return nil
}
