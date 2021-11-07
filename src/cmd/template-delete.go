package cmd

import "github.com/urfave/cli/v2"

// TemplateLoadCliFlags are the cli flags for the template load command
var TemplateDeleteCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "No safety checks",
		Value:   false,
	},
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// DeleteTemplate removes a template from the central template directory
func DeleteTemplate(c *cli.Context) error {
	infoLogger.Println("DeleteTemplate command executed")

	// Extract flags
	//dirName := c.Args().Get(0)
	//flagForce := c.Bool("force")

	return nil
}
