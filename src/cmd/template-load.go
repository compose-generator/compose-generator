package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/otiai10/copy"
	"github.com/urfave/cli/v2"
)

const (
	timeFormat = "Jan-02-06 3:04:05 PM"
)

// Cli flags for the template load command
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
	&cli.BoolFlag{
		Name:    "with-dockerfile",
		Aliases: []string{"w"},
		Usage:   "Also load the Dockerfile from the template (if existing)",
		Value:   false,
	},
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadTemplate copies a template from the central templates directory to the working directory
func LoadTemplate(c *cli.Context) error {
	// Extract flags
	dirName := c.Args().Get(0)
	flagForce := c.Bool("force")
	flagShow := c.Bool("show")
	//withDockerfile := c.Bool("with-dockerfile")

	if flagShow {
		showTemplateList()
	} else {
		sourceDir := util.GetCustomTemplatesPath() + "/" + dirName
		if dirName == "" {
			// Let the user choose a template
			dirName = askForTemplate()
			sourceDir += dirName
		} else {
			// Check if the stated template exists
			if !util.FileExists(sourceDir) {
				util.Error("Could not find template '"+dirName+"'", nil, true)
			}
		}

		// Load project
		util.P("Loading project ... ")
		proj := project.LoadProject(
			project.LoadFromDir(sourceDir),
		)
		proj.ForceConfig = flagForce
		util.Done()

		// Copy volumes over to the new template dir
		util.P("Copying volumes ...")
		copyVolumesFromTemplate(proj, sourceDir)
		util.Done()

		// Save the project to the current dir
		util.P("Saving project ... ")
		project.SaveProject(
			proj,
			project.SaveIntoDir("."),
		)
		util.Done()
	}
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askForTemplate() string {
	util.P("Loading template list ... ")
	templateMetadataList := getTemplateMetadataList()
	util.Done()
	util.Pel()

	if len(templateMetadataList) > 0 {
		var items []string
		var keys []string
		for key, metadata := range templateMetadataList {
			creationDate := time.Unix(0, int64(metadata.LastModifiedAt)).Format(timeFormat)
			keys = append(keys, key)
			items = append(items, metadata.Name+" (Saved at: "+creationDate+")")
		}
		index := util.MenuQuestionIndex("Which template do you want to load?", items)
		return keys[index]
	} else {
		util.Error("No templates found. Use \"$ compose-generator save <template-name>\" to save one.", nil, true)
	}
	return ""
}

func showTemplateList() {
	util.P("Loading template list ... ")
	templateMetadataList := getTemplateMetadataList()
	util.Done()
	util.Pel()

	if len(templateMetadataList) > 0 {
		// Show list of saved templates
		util.Heading("List of all templates:")
		for _, metadata := range templateMetadataList {
			creationDate := time.Unix(0, int64(metadata.LastModifiedAt)).Format(timeFormat)
			util.Pl(metadata.Name + " (Saved at: " + creationDate + ")")
		}
		util.Pel()
	} else {
		util.Error("No templates found. Use \"$ compose-generator save <template-name>\" to save one.", nil, true)
	}
}

func getTemplateMetadataList() map[string]*model.CGProjectMetadata {
	files, err := ioutil.ReadDir(util.GetCustomTemplatesPath())
	if err != nil {
		util.Error("Cannot access directory for custom templates", err, true)
	}
	templateMetadata := make(map[string]*model.CGProjectMetadata)
	for _, f := range files {
		if f.IsDir() {
			templatePath := util.GetCustomTemplatesPath() + "/" + f.Name()
			metadata := project.LoadProjectMetadata(
				project.LoadFromDir(templatePath),
			)
			templateMetadata[templatePath] = metadata
		}
	}
	return templateMetadata
}

func copyVolumesFromTemplate(proj *model.CGProject, sourceDir string) {
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
		copy.Copy(sourceDir+"/"+pathRel, path)
	}
}
