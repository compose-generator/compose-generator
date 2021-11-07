package cmd

import (
	"compose-generator/util"

	"github.com/urfave/cli/v2"
)

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
	dirName := c.Args().Get(0)
	flagForce := c.Bool("force")

	sourceDir := getCustomTemplatesPath() + "/" + dirName
	if dirName == "" {
		// Let the user choose a template
		dirName = askForTemplate()
		sourceDir += dirName
	} else {
		// Check if the stated template exists
		if !util.FileExists(sourceDir) {
			errorLogger.Println("Template directory '" + sourceDir + "' could not be found. Aborting")
			logError("Could not find template '"+dirName+"'", true)
		}
	}

	if !flagForce && !yesNoQuestion("Do you really wanto to delete this template?", false) {
		return nil
	}
	spinner := startProcess("Delete project ...")
	if err := removeAll(sourceDir); err != nil {
		errorLogger.Println("Could not delete template '" + sourceDir + "'")
		logError("Could not delete template", true)
	}
	stopProcess(spinner)

	return nil
}
