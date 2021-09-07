package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
	"github.com/urfave/cli/v2"
)

// Cli flags for the template save command
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
	&cli.BoolFlag{
		Name:    "with-dockerfile",
		Aliases: []string{"w"},
		Usage:   "Also save the Dockerfile in the template",
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
	//flagWithDockerfile := c.Bool("with-dockerfile")

	// Load project
	util.P("Loading project ... ")
	proj := project.LoadProject()
	proj.ForceConfig = flagForce
	util.Done()

	// Ask for template name
	if name == "" {
		for another := true; another; another = isTemplateExisting(proj.Name) {
			name = util.TextQuestionWithDefault("How would you like to call your template:", proj.Name)
			proj.Name = name
		}
	}

	// Create the new template
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	os.MkdirAll(targetDir, 0755)

	// Copy volumes over to the new template dir
	util.P("Copying volumes ... ")
	copyVolumesToTemplate(proj, targetDir)
	util.Done()

	// Save the project to the templates dir
	util.P("Saving project ... ")
	project.SaveProject(
		proj,
		project.SaveIntoDir(targetDir),
	)
	util.Done()

	// Delete the original project if the stash flag is set
	if flagStash {
		util.P("Stashing project ... ")
		project.DeleteProject(proj)
		util.Done()
	}

	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func isTemplateExisting(name string) bool {
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	if util.FileExists(targetDir) {
		util.Error("Template with the name '"+name+"' already exists", nil, false)
		return true
	}
	return false
}

func copyVolumesToTemplate(proj *model.CGProject, targetDir string) {
	currentAbs, err := filepath.Abs(".")
	if err != nil {
		util.Error("Could not find absolute path of current dir", err, true)
	}
	for _, path := range proj.GetAllVolumePathsNormalized() {
		pathAbs, err := filepath.Abs(path)
		if err != nil {
			util.Error("Could not find absolute path of volume dir", err, true)
		}
		pathRel, err := filepath.Rel(currentAbs, pathAbs)
		if err != nil {
			util.Error("Could not copy volume '"+path+"'", err, false)
			continue
		}
		copy.Copy(path, targetDir+"/"+pathRel)
	}
}
