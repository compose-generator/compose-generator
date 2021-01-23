package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const VERSION = "1.0.0"

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
		Version: VERSION,
		Authors: []*cli.Author{
			{Name: "Marc Auberer", Email: "marc.auberer@chillibits.com"},
		},
		Copyright: "© 2021 Marc Auberer",
		Usage:     "Generate docker compose configuration files for your projects.",
		Action: func(c *cli.Context) error {
			generate()
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "advanced", Aliases: []string{"a"}, Usage: "Generate compose file in advanced mode"},
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds a service to an existing compose file",
				Action: func(c *cli.Context) error {
					add()
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rm"},
				Usage:   "Removes a service from an existing compose file",
				Action: func(c *cli.Context) error {
					remove()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func generate() {
	fmt.Print("What is the name of your project: ")
}

func add() {

}

func remove() {

}
