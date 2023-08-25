/*
Copyright © 2021-2023 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"github.com/urfave/cli/v2"
)

// CliCommands is a list of all available cli commands for Compose Generator
var CliCommands = []*cli.Command{
	{
		Name:    "generate",
		Aliases: []string{"g", "gen"},
		Usage:   "Generate a Docker Compose configuration",
		Flags:   GenerateCliFlags,
		Action:  Generate,
	},
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add a service to an existing Compose file",
		Flags:   AddCliFlags,
		Action:  Add,
	},
	{
		Name:    "remove",
		Aliases: []string{"r", "rm"},
		Usage:   "Remove a service from an existing Compose file",
		Flags:   RemoveCliFlags,
		Action:  Remove,
	},
	{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "Manage snapshots of your Compose configuration for later use",
		Subcommands: []*cli.Command{
			{
				Name:    "save",
				Aliases: []string{"s"},
				Usage:   "Save a custom template.",
				Flags:   TemplateSaveCliFlags,
				Action:  SaveTemplate,
			},
			{
				Name:    "load",
				Aliases: []string{"l"},
				Usage:   "Load a custom template.",
				Flags:   TemplateLoadCliFlags,
				Action:  LoadTemplate,
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a custom template.",
				Flags:   TemplateDeleteCliFlags,
				Action:  DeleteTemplate,
			},
		},
	},
	{
		Name:    "install",
		Aliases: []string{"i", "in"},
		Usage:   "Install Docker and Docker Compose with a single command",
		Hidden:  isDockerizedEnvironment(),
		Action:  Install,
	},
	{
		Name:    "predefined-template",
		Aliases: []string{"p", "pd", "pt"},
		Usage:   "Manage predefined service templates [DEV]",
		Hidden:  !isDevVersion(),
		Subcommands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Create a blank predefined service template",
				Flags:   PredefinedTemplateNewCliFlags,
				Action:  NewPredefinedTemplate,
			},
		},
	},
}
