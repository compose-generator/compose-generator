/*
Copyright © 2021 Compose Generator Contributors
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
		Usage:   "Generates a docker compose configuration",
		Flags:   GenerateCliFlags,
		Action:  Generate,
	},
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Adds a service to an existing compose file",
		Flags:   AddCliFlags,
		Action:  Add,
	},
	{
		Name:    "remove",
		Aliases: []string{"r", "rm"},
		Usage:   "Removes a service from an existing compose file",
		Flags:   RemoveCliFlags,
		Action:  Remove,
	},
	{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "Manages snapshots of your compose configuration for later use",
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
		Usage:   "Installs Docker and Docker Compose with a single command",
		Hidden:  isDockerizedEnvironment(),
		Action:  Install,
	},
	{
		Name:    "predefined-template",
		Aliases: []string{"p", "pd"},
		Usage:   "Manages predefined service templates (this command does only exist in dev environments)",
		Hidden:  !isDevVersion(),
		Subcommands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Creates a blank predefined service template",
				Flags:   PredefinedTemplateNewCliFlags,
				Action:  NewPredefinedTemplate,
			},
		},
	},
}
