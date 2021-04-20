package cmd

import (
	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/util"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/otiai10/copy"
)

const (
	timeFormat = "Jan-02-06 3:04:05 PM"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// SaveTemplate copies the compose configuration in the current directory to a central templates directory
func SaveTemplate(name string, flagStash bool, flagForce bool, withDockerfile bool) {
	if name == "" {
		name = util.TextQuestion("How would you like to call your template: ")
	}
	// Check if templated with that name exists already
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	if !flagForce && util.FileExists(targetDir) {
		result := util.YesNoQuestion("There is already a template called '"+name+"'. Do you want to replace it?", false)
		if !result {
			return
		}
		util.Pel()
	}
	// Create metadata
	util.P("Creating metadata file ... ")
	os.MkdirAll(targetDir, os.ModePerm)
	var metadata model.TemplateMetadata
	metadata.Label = name
	metadata.CreationTime = time.Now().UnixNano() / int64(time.Millisecond)
	metadataJSON, _ := json.MarshalIndent(metadata, "", " ")
	err := ioutil.WriteFile(targetDir+"/metadata.json", metadataJSON, 0777)
	if err != nil {
		util.Error("Could not write metadata.", err, true)
	}
	util.Done()
	// Save template
	util.P("Saving template ... ")
	var savedFiles []string
	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			conditionToSkip := !strings.HasSuffix(src, "docker-compose.yml") && !strings.HasSuffix(src, "environment.env") && (!withDockerfile || !strings.HasSuffix(src, "Dockerfile")) && !strings.Contains(src, "volumes")
			if !conditionToSkip && flagStash {
				savedFiles = append(savedFiles, src)
			}
			return conditionToSkip, nil
		},
		OnDirExists: func(src string, dst string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err = copy.Copy(".", targetDir, opt)
	if err != nil {
		util.Error("Could not copy files. Is the permission granted?", err, true)
	}
	util.Done()
	// Delete files from source dir if stash flag is set
	if flagStash {
		util.P("Stashing ... ")
		for _, f := range savedFiles {
			os.RemoveAll(f)
		}
		util.Done()
	}
}

// LoadTemplate copies a template from the central templates directory to the working directory
func LoadTemplate(name string, flagForce bool, flagShow bool, withDockerfile bool) {
	// Execute safety checks
	/*if !flagForce {
		util.PrintSafetyWarning(false, withDockerfile)
	}*/
	// Check if the template exists
	targetDir := util.GetCustomTemplatesPath() + "/" + name
	if name != "" && !util.FileExists(targetDir) {
		util.Error("Template with the name '"+name+"' could not be found. You can query a list of the templates by executing 'compose-generator template load'.", nil, true)
	} else {
		// Load stacks from templates
		if templateData := parser.ParseCustomTemplates(); len(templateData) > 0 {
			// Show list of saved templates
			var items []string
			for _, t := range templateData {
				creationDate := time.Unix(0, t.CreationTime*int64(time.Millisecond)).Format(timeFormat)
				items = append(items, t.Label+" (Saved at: "+creationDate+")")
			}
			if flagShow {
				util.Heading("List of all templates:")
				util.Pel()
				for _, item := range items {
					util.Pl(item)
				}
				os.Exit(0)
			} else {
				index := util.MenuQuestionIndex("Saved templates", items)
				targetDir = targetDir + templateData[index].Label
			}
		} else {
			util.Warning("No templates found. Use \"$ compose-generator save <template-name>\" to save one.")
			os.Exit(0)
		}
		util.Pel()
	}
	// Load template
	util.P("Loading template ... ")
	srcPath := targetDir
	dstPath := "."

	os.Remove(dstPath + "/docker-compose.yml")
	os.Remove(dstPath + "/environment.env")
	if withDockerfile {
		os.Remove(dstPath + "/Dockerfile")
	}
	os.RemoveAll(dstPath + "/volumes")

	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			conditionToSkip := strings.HasSuffix(src, "metadata.json") || (!withDockerfile && strings.HasSuffix(src, "Dockerfile"))
			return conditionToSkip, nil
		},
		OnDirExists: func(src string, dst string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := copy.Copy(srcPath, dstPath, opt)
	if err != nil {
		util.Error("Could not load template files.", err, true)
	}
	util.Done()
}
