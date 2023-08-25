/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"compose-generator/cmd"
	"compose-generator/util"
)

func main() {
	// Version flag
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Prints the version of Compose Generator",
	}

	// Main cli configuration
	app := &cli.App{
		Name:     "Compose Generator",
		HelpName: "compose-generator",
		Version:  util.BuildVersion(util.Version, util.Commit, util.Date, util.BuiltBy),
		Authors: []*cli.Author{
			{
				Name:  "Marc Auberer",
				Email: "marc.auberer@chillibits.com",
			},
		},
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		Usage:                  "Generate and manage Docker Compose configuration files for your projects.",
		Copyright:              "© 2021-2023 Compose Generator Contributors",
		Flags:                  cmd.GenerateCliFlags,
		Action:                 cmd.Generate,
		Commands:               cmd.CliCommands,
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ErrorLogger.Println("Fatal error initializing cli. Aborting.")
		log.Fatal(err)
	}
}
