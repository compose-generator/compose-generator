/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	timeFormat = "Jan-02-06 3:04:05 PM"
)

// TemplateLoadCliFlags are the cli flags for the template load command
var TemplateLoadCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "No safety checks",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "Do not load a template. Instead only list all templates and terminate",
		Value:   false,
	},
}

var getTemplateMetadataListMockable = getTemplateMetadataList
var askForTemplateMockable = askForTemplate

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadTemplate copies a template from the central templates directory to the working directory
func LoadTemplate(c *cli.Context) error {
	infoLogger.Println("LoadTemplate command executed")

	// Extract flags
	dirName := c.Args().Get(0)
	flagForce := c.Bool("force")
	flagShow := c.Bool("show")

	if flagShow {
		showTemplateList()
	} else {
		sourceDir := getCustomTemplatesPath() + "/" + dirName
		if dirName == "" {
			// Let the user choose a template
			dirName = askForTemplate("Which template do you want to load?")
			sourceDir += dirName
		} else {
			// Check if the stated template exists
			if !util.FileExists(sourceDir) {
				errorLogger.Println("Template directory '" + sourceDir + "' could not be found. Aborting")
				logError("Could not find template '"+dirName+"'", true)
			}
		}

		// Load project
		spinner := startProcess("Loading project ...")
		proj := project.LoadProject(
			project.LoadFromDir(sourceDir),
		)
		proj.ForceConfig = flagForce
		stopProcess(spinner)

		// Copy volumes and build contexts over to the new template dir
		spinner = startProcess("Loading volumes and build contexts ...")
		copyVolumesAndBuildContextsFromTemplate(proj, sourceDir)
		stopProcess(spinner)

		// Save the project to the current dir
		spinner = startProcess("Saving project ...")
		project.SaveProject(
			proj,
			project.SaveIntoDir("."),
		)
		stopProcess(spinner)
	}
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askForTemplate(question string) string {
	spinner := startProcess("Loading template list ...")
	templateMetadataList := getTemplateMetadataListMockable()
	stopProcess(spinner)
	pel()

	if len(templateMetadataList) > 0 {
		var items []string
		var keys []string
		for key, metadata := range templateMetadataList {
			creationDate := time.Unix(0, int64(metadata.LastModifiedAt)).Format(timeFormat)
			keys = append(keys, key)
			items = append(items, metadata.Name+" (Saved at: "+creationDate+")")
		}
		index := menuQuestionIndex(question, items)
		return keys[index]
	}
	warningLogger.Println("Template dir is empty")
	logWarning("No templates found. Use \"$ compose-generator save <template-name>\" to save one.")
	return ""
}

func showTemplateList() {
	spinner := startProcess("Loading template list ...")
	templateMetadataList := getTemplateMetadataListMockable()
	stopProcess(spinner)
	pel()

	if len(templateMetadataList) > 0 {
		// Show list of saved templates
		printHeading("List of all templates:")
		for _, metadata := range templateMetadataList {
			creationDate := time.Unix(0, int64(metadata.LastModifiedAt)).Format(timeFormat)
			pl(metadata.Name + " (Saved at: " + creationDate + ")")
		}
		pel()
	} else {
		errorLogger.Println("Template dir is empty")
		logError("No templates found. Use \"$ compose-generator save <template-name>\" to save one.", true)
	}
}

func getTemplateMetadataList() map[string]*model.CGProjectMetadata {
	files, err := readDir(getCustomTemplatesPath())
	if err != nil {
		errorLogger.Println("Could not access '" + getCustomTemplatesPath() + "': " + err.Error())
		logError("Cannot access directory for custom templates", true)
		return nil
	}
	templateMetadata := make(map[string]*model.CGProjectMetadata)
	for _, f := range files {
		if f.IsDir() {
			templatePath := getCustomTemplatesPath() + "/" + f.Name()
			metadata := loadProjectMetadata(
				project.LoadFromDir(templatePath),
			)
			templateMetadata[templatePath] = metadata
		}
	}
	return templateMetadata
}

func copyVolumesAndBuildContextsFromTemplate(proj *model.CGProject, sourceDir string) {
	currentAbs, err := abs(".")
	if err != nil {
		errorLogger.Println("Could not find absolute path of current dir: " + err.Error())
		logError("Could not find absolute path of current dir", true)
		return
	}
	paths := append(proj.GetAllVolumePaths(), proj.GetAllBuildContextPaths()...)
	for _, path := range normalizePaths(paths) {
		pathAbs, err := abs(path)
		if err != nil {
			errorLogger.Println("Could not find absolute path of current dir: " + err.Error())
			logError("Could not find absolute path of volume dir", true)
			return
		}
		pathRel, err := rel(currentAbs, pathAbs)
		if err != nil {
			warningLogger.Println("Could not copy volume '" + path + "': " + err.Error())
			logError("Could not copy volume '"+path+"'", false)
			continue
		}
		if err := copyDir(sourceDir+"/"+pathRel, path); err != nil {
			warningLogger.Println("Could not copy volumes from '" + sourceDir + "/" + pathRel + "' to '" + path + "': " + err.Error())
			logWarning("Could not copy volumes from '" + sourceDir + "/" + pathRel + "' to '" + path + "'")
		}
	}
}
