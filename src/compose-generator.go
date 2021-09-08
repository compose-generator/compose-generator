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
		Usage:   "Prints the version of compose-generator",
	}

	// Main cli configuration
	app := &cli.App{
		Name:    "compose-generator",
		Version: util.BuildVersion(util.Version, util.Commit, util.Date, util.BuiltBy),
		Authors: []*cli.Author{
			{Name: "Marc Auberer", Email: "marc.auberer@chillibits.com"},
		},
		Copyright: "© 2021 Marc Auberer",
		Usage:     "Generate and manage docker compose configuration files for your projects.",
		Flags:     cmd.GenerateCliFlags,
		Action:    cmd.Generate,
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "Generates a docker compose configuration",
				Flags:   cmd.GenerateCliFlags,
				Action:  cmd.Generate,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds a service to an existing compose file",
				Flags:   cmd.AddCliFlags,
				Action:  cmd.Add,
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rm"},
				Usage:   "Removes a service from an existing compose file",
				Flags:   cmd.RemoveCliFlags,
				Action:  cmd.Remove,
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "Saves / loads snapshots of your compose configuration for later use",
				Subcommands: []*cli.Command{
					{
						Name:    "save",
						Aliases: []string{"s"},
						Usage:   "Save a custom template.",
						Flags:   cmd.TemplateSaveCliFlags,
						Action:  cmd.SaveTemplate,
					},
					{
						Name:    "load",
						Aliases: []string{"l"},
						Usage:   "Load a custom template.",
						Flags:   cmd.TemplateLoadCliFlags,
						Action:  cmd.LoadTemplate,
					},
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i", "in"},
				Usage:   "Installs Docker and Docker Compose with a single command",
				Action:  cmd.Install,
			},
		},
		UseShortOptionHandling: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
