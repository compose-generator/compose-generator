package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "generate",
		Usage: "Generate a docker compose configuration file",
		Action: func(c *cli.Context) error {
			fmt.Print("What is the name of your project: ")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
