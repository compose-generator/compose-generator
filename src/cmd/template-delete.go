/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

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
	dirName := c.Args().Get(0)
	flagForce := c.Bool("force")

	return delete(dirName, flagForce)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func delete(dirName string, flagForce bool) error {
	// Determine template dir
	sourceDir := getCustomTemplatesPath() + "/" + dirName
	if dirName == "" {
		// Let the user choose a template
		dirName = askForTemplateMockable("Which template do you want to delete?")
		sourceDir += dirName
	} else {
		// Check if the stated template exists
		if !fileExists(sourceDir) {
			errorLogger.Println("Template directory '" + sourceDir + "' could not be found. Aborting")
			logError("Could not find template '"+dirName+"'", true)
			return nil
		}
	}

	// Safety check
	if !flagForce && !yesNoQuestion("Do you really want to delete this template?", false) {
		return nil
	}

	// Delete the stated directory recursively
	spinner := startProcess("Deleting project ...")
	if err := removeAll(sourceDir); err != nil {
		errorLogger.Println("Could not delete template '" + sourceDir + "'")
		logError("Could not delete template", true)
		return err
	}
	stopProcess(spinner)
	return nil
}
