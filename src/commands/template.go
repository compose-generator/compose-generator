package commands

import (
	"compose-generator/model"
	"compose-generator/parser"
	"compose-generator/utils"
	"encoding/json"
	"fmt"
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
func SaveTemplate(name string, flagStash bool, flagForce bool) {
	utils.ClearScreen()

	if name == "" {
		name = utils.TextQuestion("How would you like to call your template: ")
	}
	// Check if templated with that name exists already
	targetDir := utils.GetTemplatesPath() + "/" + name
	if !flagForce && utils.FileExists(targetDir) {
		result := utils.YesNoQuestion("There is already a template called '"+name+"'. Do you want to replace it?", false)
		if !result {
			return
		}
		utils.Pel()
	}
	// Create metadata
	fmt.Print("Creating metadata file ...")
	os.MkdirAll(targetDir, os.ModePerm)
	var metadata model.TemplateMetadata
	metadata.Label = name
	metadata.CreationTime = time.Now().UnixNano() / int64(time.Millisecond)
	metadataJSON, _ := json.MarshalIndent(metadata, "", " ")
	err := ioutil.WriteFile(targetDir+"/metadata.json", metadataJSON, 0777)
	if err != nil {
		utils.Error("Could not write metadata.", true)
	}
	utils.PrintDone()
	// Save template
	fmt.Print("Saving template ...")
	var savedFiles []string
	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			conditionToSkip := !strings.HasSuffix(src, "docker-compose.yml") && !strings.HasSuffix(src, "environment.env") && !strings.Contains(src, "volumes")
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
		utils.Error("Could not copy files. Is the permission granted?", true)
	}
	utils.PrintDone()
	// Delete files from source dir if stash flag is set
	if flagStash {
		fmt.Print("Stashing ...")
		for _, f := range savedFiles {
			os.RemoveAll(f)
		}
		utils.PrintDone()
	}
}

// LoadTemplate copies a template from the central templates directory to the working directory
func LoadTemplate(name string, flagForce bool) {
	utils.ClearScreen()

	// Execute safety checks
	if !flagForce {
		utils.ExecuteSafetyFileChecks()
	}
	// Check if the template exists
	targetDir := utils.GetTemplatesPath() + "/" + name
	if name != "" && !utils.FileExists(targetDir) {
		utils.Error("Template with the name '"+name+"' could not be found. You can query a list of the templates by executing 'compose-generator template load'.", true)
	} else if name == "" {
		// Load stacks from templates
		templateData := parser.ParseTemplates()
		// Show list of saved templates
		var items []string
		for _, t := range templateData {
			creationDate := time.Unix(0, t.CreationTime*int64(time.Millisecond)).Format(timeFormat)
			items = append(items, t.Label+" (Saved at: "+creationDate+")")
		}
		index := utils.MenuQuestionIndex("Saved templates", items)
		targetDir = targetDir + templateData[index].Label
		utils.Pel()
	}
	// Load template
	fmt.Print("Loading template ...")
	srcPath := targetDir
	dstPath := "."

	os.Remove(dstPath + "/docker-compose.yml")
	os.Remove(dstPath + "/environment.env")
	os.RemoveAll(dstPath + "/volumes")

	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			return strings.HasSuffix(src, "metadata.json"), nil
		},
		OnDirExists: func(src string, dst string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := copy.Copy(srcPath, dstPath, opt)
	if err != nil {
		utils.Error("Could not load template files.", true)
	}
	utils.PrintDone()
}
