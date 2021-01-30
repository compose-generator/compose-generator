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

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

const (
	timeFormat = "Jan-02-06 3:04:05 PM"
)

func SaveTemplate(name string, flag_stash bool, flag_force bool) {
	if name == "" {
		name = utils.TextQuestion("How would you like to call your template: ")
	}
	// Check if templated with that name exists already
	target_dir := utils.GetTemplatesPath() + "/" + name
	if !flag_force && utils.FileExists(target_dir) {
		result := utils.YesNoQuestion("There is already a template called '"+name+"'. Do you want to replace it?", false)
		if !result {
			return
		}
	}
	fmt.Println()
	// Create metadata
	fmt.Print("Creating metadata file ...")
	os.MkdirAll(target_dir, os.ModePerm)
	var metadata model.TemplateMetadata
	metadata.Label = name
	metadata.CreationTime = time.Now().UnixNano() / int64(time.Millisecond)
	metadata_json, _ := json.MarshalIndent(metadata, "", " ")
	err := ioutil.WriteFile(target_dir+"/metadata.json", metadata_json, 0777)
	if err != nil {
		utils.Error("Could not write metadata.", true)
	}
	color.Green(" done")
	// Save template
	fmt.Print("Saving template ...")
	var saved_files []string
	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			condition_to_skip := !strings.HasSuffix(src, "docker-compose.yml") && !strings.HasSuffix(src, "environment.env") && !strings.Contains(src, "volumes")
			if !condition_to_skip && flag_stash {
				saved_files = append(saved_files, src)
			}
			return condition_to_skip, nil
		},
		OnDirExists: func(src string, dst string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err = copy.Copy(".", target_dir, opt)
	if err != nil {
		utils.Error("Could not copy files. Is the permission granted?", true)
	}
	color.Green(" done")
	// Delete files from source dir if stash flag is set
	if flag_stash {
		fmt.Print("Stashing ...")
		for _, f := range saved_files {
			os.RemoveAll(f)
		}
		color.Green(" done")
	}
}

func LoadTemplate(name string, flag_force bool) {
	// Execute safety checks
	if !flag_force {
		utils.ExecuteSafetyFileChecks()
	}
	// Check if the template exists
	target_dir := utils.GetTemplatesPath() + "/" + name
	if name != "" && !utils.FileExists(target_dir) {
		utils.Error("Template with the name '"+name+"' could not be found. You can query a list of the templates by executing 'compose-generator template load'.", true)
	} else if name == "" {
		// Load stacks from templates
		template_data := parser.ParseTemplates()
		// Show list of saved templates
		var items []string
		for _, t := range template_data {
			creation_date := time.Unix(0, t.CreationTime*int64(time.Millisecond)).Format(timeFormat)
			items = append(items, t.Label+" (Saved at: "+creation_date+")")
		}
		index, _ := utils.MenuQuestion("Saved templates", items)
		target_dir = target_dir + template_data[index].Label
		fmt.Println()
	}
	// Load template
	fmt.Print("Loading template ...")
	src_path := target_dir
	dst_path := "."

	os.Remove(dst_path + "/docker-compose.yml")
	os.Remove(dst_path + "/environment.env")
	os.RemoveAll(dst_path + "/volumes")

	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			return strings.HasSuffix(src, "metadata.json"), nil
		},
		OnDirExists: func(src string, dst string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := copy.Copy(src_path, dst_path, opt)
	if err != nil {
		utils.Error("Could not load template files.", true)
	}
	color.Green(" done")
}
