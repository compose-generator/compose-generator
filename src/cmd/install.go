package cmd

import (
	"os/exec"
	"strings"

	"compose-generator/pass"
	"compose-generator/util"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(c *cli.Context) error {
	// Check if Compose Generator runs in dockerized environment
	if util.IsDockerizedEnvironment() {
		util.Error("You are currently using the dockerized version of Compose Generator. To use this command, please install Compose Generator on your system. Visit https://www.compose-generator.com/install/linux or https://www.compose-generator.com/install/windows for more details.", nil, true)
	}

	pass.InstallDocker()

	// Check if installation was successful
	if util.CommandExists("docker") {
		util.Pel()
		cmd := exec.Command("docker", "-v")
		dockerVersion, _ := cmd.CombinedOutput()
		color.Yellow("Congrats! You have installed " + strings.TrimRight(string(dockerVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
	} else {
		color.Red("An error occurred while installing Docker.")
	}
	return nil
}
