package cmd

import (
	"compose-generator/model"
	"compose-generator/project"
	"compose-generator/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/otiai10/copy"
)

const (
	timeFormat = "Jan-02-06 3:04:05 PM"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// SaveTemplate copies the compose configuration from the current directory to a central templates directory
func SaveTemplate(
	name string,
	flagStash bool,
	flagForce bool,
	flagWithDockerfile bool,
) {
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
	util.P("Copying volumes ...")
	copyVolumes(proj, targetDir)
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
		util.P("Stashing project ...")
		project.DeleteProject(proj)
		util.Done()
	}
}

// LoadTemplate copies a template from the central templates directory to the working directory
func LoadTemplate(
	dirName string,
	flagForce bool,
	flagShow bool,
	withDockerfile bool,
) {
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
		copyVolumes(proj, ".")
		util.Done()

		// Save the project to the current dir
		util.P("Saving project ... ")
		project.SaveProject(
			proj,
			project.SaveIntoDir("."),
		)
		util.Done()
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func askForTemplate() string {
	// Load stacks from templates
	util.P("Loading template list ...")
	templateMetadataList := getTemplateMetadataList()
	util.Done()
	if len(templateMetadataList) > 0 {
		var items []string
		var keys []string
		for key, metadata := range templateMetadataList {
			lastModifiedNanos := metadata.LastModifiedAt * int64(time.Millisecond)
			creationDate := time.Unix(0, lastModifiedNanos).Format(timeFormat)
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
	util.P("Loading template list ...")
	templateMetadataList := getTemplateMetadataList()
	util.Done()
	if len(templateMetadataList) > 0 {
		// Show list of saved templates
		util.Heading("List of all templates:")
		util.Pel()
		for _, metadata := range templateMetadataList {
			lastModifiedNanos := metadata.LastModifiedAt * int64(time.Millisecond)
			creationDate := time.Unix(0, lastModifiedNanos).Format(timeFormat)
			util.Pl(metadata.Name + " (Saved at: " + creationDate + ")")
		}
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

func isTemplateExisting(name string) bool {
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	if util.FileExists(targetDir) {
		util.Error("Template with the name '"+name+"' already exists", nil, false)
		return true
	}
	return false
}

func copyVolumes(proj *model.CGProject, targetDir string) {
	currentAbs, err := filepath.Abs(".")
	if err != nil {
		util.Error("Could not find absolute path of current dir", err, true)
	}
	for _, path := range proj.GetAllVolumePathsNormalized() {
		pathRel, err := filepath.Rel(currentAbs, path)
		if err != nil {
			util.Error("Could not copy volume '"+path+"'", err, false)
			continue
		}
		copy.Copy(path, targetDir+"/"+pathRel)
	}
}
