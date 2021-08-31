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
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode", Value: false},
			&cli.PathFlag{Name: "config", Aliases: []string{"c"}, Usage: "Pass a configuration as a `FILE` with predefined answers. Works good for CI"},
			&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file", Value: false},
			&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks", Value: false},
			&cli.BoolFlag{Name: "with-instructions", Aliases: []string{"i"}, Usage: "Generates a README.md file with instructions to use the template", Value: false},
			&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file", Value: false},
		},
		Action: func(c *cli.Context) error {
			cmd.Generate(c.Path("config"), c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"), c.Bool("with-instructions"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "Generates a docker compose configuration",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode", Value: false},
					&cli.PathFlag{Name: "config", Aliases: []string{"c"}, Usage: "Pass a configuration as a `FILE` with predefined answers. Works good for CI"},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file", Value: false},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks", Value: false},
					&cli.BoolFlag{Name: "with-instructions", Aliases: []string{"i"}, Usage: "Generates a README.md file with instructions to use the template", Value: false},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file", Value: false},
				},
				Action: func(c *cli.Context) error {
					cmd.Generate(c.Path("config"), c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"), c.Bool("with-instructions"))
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds a service to an existing compose file",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode", Value: false},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file", Value: false},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file", Value: false},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks", Value: false},
				},
				Action: func(c *cli.Context) error {
					cmd.Add(c.Bool("advanced"), c.Bool("run"), c.Bool("detached"), c.Bool("force"))
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rm", "rem"},
				Usage:   "Removes a service from an existing compose file",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Show questions for advanced customization", Value: false},
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Skip safety checks", Value: false},
					&cli.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Run docker-compose after creating the compose file", Value: false},
					&cli.BoolFlag{Name: "detached", Aliases: []string{"d"}, Usage: "Run docker-compose detached after creating the compose file", Value: false},
					&cli.BoolFlag{Name: "with-volumes", Aliases: []string{"v"}, Usage: "Remove associated volumes", Value: false},
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
							&cli.BoolFlag{Name: "stash", Aliases: []string{"s"}, Usage: "Move the regarding files, instead of copying them", Value: false},
							&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks", Value: false},
							&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Also save the Dockerfile in the template", Value: false},
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
							&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "No safety checks", Value: false},
							&cli.BoolFlag{Name: "show", Aliases: []string{"s"}, Usage: "Do not load a template. Instead only list all templates and terminate", Value: false},
							&cli.BoolFlag{Name: "with-dockerfile", Aliases: []string{"w"}, Usage: "Also load the Dockerfile from the template (if existing)", Value: false},
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
				Aliases: []string{"i", "in"},
				Usage:   "Installs Docker and Docker Compose with a single command",
				Action: func(c *cli.Context) error {
					cmd.Install()
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
