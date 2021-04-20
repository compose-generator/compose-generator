package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"compose-generator/cmd"
	"compose-generator/util"
)

func main() {
	const VERSION = "0.7.0"

	// Version flag
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Prints the version of compose-generator",
	}

	// Main cli configuration
	app := &cli.App{
		Name:    "compose-generator",
		Version: VERSION,
		Authors: []*cli.Author{
			{Name: "Marc Auberer", Email: "marc.auberer@chillibits.com"},
		},
		Copyright: "© 2021 Marc Auberer",
		Usage:     "Generate and manage docker compose configuration files for your projects.",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode"},
			&cli.PathFlag{Name: "config", Aliases: []string{"c"}, Usage: "Pass a configuration file with predefined answers. Works good for CI"},
			&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file"},
			&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks"},
			&cli.BoolFlag{Name: "with-instructions", Aliases: []string{"i"}, Usage: "Generates a README.md file with instructions to use the template"},
			&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file"},
			&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Generates the Dockerfile for your project"},
		},
		Action: func(c *cli.Context) error {
			util.CheckForServiceTemplateUpdate(VERSION)
			cmd.Generate(c.Path("config"), c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"), c.Bool("with-instructions"), c.Bool("with-dockerfile"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generates a docker compose configuration",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode"},
					&cli.PathFlag{Name: "config", Aliases: []string{"c"}, Usage: "Pass a configuration file with predefined answers. Works good for CI"},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file"},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks"},
					&cli.BoolFlag{Name: "with-instructions", Aliases: []string{"i"}, Usage: "Generates a README.md file with instructions to use the template"},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file"},
					&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Generates the Dockerfile for your project"},
				},
				Action: func(c *cli.Context) error {
					util.CheckForServiceTemplateUpdate(VERSION)
					cmd.Generate(c.Path("config"), c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"), c.Bool("with-instructions"), c.Bool("with-dockerfile"))
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds a service to an existing compose file",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode"},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file"},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file"},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks"},
				},
				Action: func(c *cli.Context) error {
					util.CheckForServiceTemplateUpdate(VERSION)
					cmd.Add(c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"))
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rm"},
				Usage:   "Removes a service from an existing compose file",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Show questions for advanced customization"},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks"},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file"},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file"},
					&cli.BoolFlag{Name: "with-volumes", Aliases: []string{"v"}, Usage: "Remove associated volumes"},
				},
				Action: func(c *cli.Context) error {
					serviceNames := c.Args().Slice()
					cmd.Remove(serviceNames, c.Bool("run"), c.Bool("detached"), c.Bool("with-volumes"), c.Bool("force"), c.Bool("advanced"))
					return nil
				},
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "stash", Aliases: []string{"s"}, Usage: "Move the regarding files, instead of copying them"},
							&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks"},
							&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Also save the Dockerfile in the template"},
						},
						Action: func(c *cli.Context) error {
							name := c.Args().Get(0)
							cmd.SaveTemplate(name, c.Bool("stash"), c.Bool("force"), c.Bool("with-dockerfile"))
							return nil
						},
					},
					{
						Name:    "load",
						Aliases: []string{"l"},
						Usage:   "Load a custom template.",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks"},
							&cli.BoolFlag{Name: "show", Aliases: []string{"s"}, Usage: "Do not load a template. Instead only list all templates and terminate"},
							&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Also load the Dockerfile from the template (if existing)"},
						},
						Action: func(c *cli.Context) error {
							name := c.Args().Get(0)
							cmd.LoadTemplate(name, c.Bool("force"), c.Bool("show"), c.Bool("with-dockerfile"))
							return nil
						},
					},
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Installs Docker and Docker Compose with a single command",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "only-compose", Aliases: []string{"c"}, Usage: "Only install Docker Compose"},
					&cli.BoolFlag{Name: "only-docker", Aliases: []string{"d"}, Usage: "Only install Docker"},
				},
				Action: func(c *cli.Context) error {
					cmd.Install(c.Bool("only-compose"), c.Bool("only-docker"))
					return nil
				},
			},
		},
		UseShortOptionHandling: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
