/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"os"

	"github.com/urfave/cli/v2"
)

// TemplateSaveCliFlags are the cli flags for the template save command
var TemplateSaveCliFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "stash",
		Aliases: []string{"s"},
		Usage:   "Move the regarding files, instead of copying them",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "No safety checks",
		Value:   false,
	},
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// SaveTemplate copies the compose configuration from the current directory to a central templates directory
func SaveTemplate(c *cli.Context) error {
	// Extract flags
	name := c.Args().Get(0)
	flagStash := c.Bool("stash")
	flagForce := c.Bool("force")

	// Load project
	spinner := util.StartProcess("Loading project ...")
	proj := project.LoadProject()
	proj.ForceConfig = flagForce
	util.StopProcess(spinner)

	// Ask for template name
	if name == "" {
		for another := true; another; another = isTemplateExisting(proj.Name) {
			name = util.TextQuestionWithDefault("How would you like to call your template:", proj.Name)
			proj.Name = name
		}
	}

	// Create the new template
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	if err := os.MkdirAll(targetDir, 0750); err != nil {
		util.Error("Could not create template dir", err, true)
	}

	// Copy volumnes and build contexts over to the new template dir
	spinner = util.StartProcess("Saving volumes and build contexts ...")
	copyVolumesAndBuildContextsToTemplate(proj, targetDir)
	util.StopProcess(spinner)

	// Save the project to the templates dir
	spinner = util.StartProcess("Saving project ...")
	project.SaveProject(
		proj,
		project.SaveIntoDir(targetDir),
	)
	util.StopProcess(spinner)

	// Delete the original project if the stash flag is set
	if flagStash {
		spinner := util.StartProcess("Stashing project ...")
		project.DeleteProject(proj)
		util.StopProcess(spinner)
	}

	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func isTemplateExisting(name string) bool {
	targetDir := getCustomTemplatesPath() + "/" + name
	if fileExists(targetDir) {
		printError("Template with the name '"+name+"' already exists", nil, false)
		return true
	}
	return false
}

func copyVolumesAndBuildContextsToTemplate(proj *model.CGProject, targetDir string) {
	currentAbs, err := abs(".")
	if err != nil {
		printError("Could not find absolute path of current dir", err, true)
	}
	// Copy volumes to template dir
	paths := append(proj.GetAllVolumePaths(), proj.GetAllBuildContextPaths()...)
	for _, path := range normalizePaths(paths) {
		pathAbs, err := abs(path)
		if err != nil {
			printError("Could not find absolute path of volume / build context dir", err, true)
		}
		pathRel, err := rel(currentAbs, pathAbs)
		if err != nil {
			printError("Could not copy volume / build context '"+path+"'", err, false)
			continue
		}
		if copyDir(path, targetDir+"/"+pathRel) != nil {
			printWarning("Could not copy volume / build context from '" + path + "' to '" + targetDir + "/" + pathRel + "'")
		}
	}
}
