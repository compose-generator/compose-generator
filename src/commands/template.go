package commands

import (
	"compose-generator/utils"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

func SaveTemplate(name string, flag_stash bool) {
	if name == "" {
		name = utils.TextQuestion("How would you like to call your template: ")
	}
	// Check if templated with that name exists already
	target_dir := utils.GetTemplatesPath() + "/" + name
	if utils.FileExists(target_dir) {
		result := utils.YesNoQuestion("There is already a template called '"+name+"'. Do you want to replace it?", false)
		if !result {
			return
		}
	}
	// Save template
	fmt.Println()
	fmt.Print("Saving template ...")
	os.Mkdir(target_dir, os.ModePerm)
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
	err := copy.Copy(".", target_dir, opt)
	if err != nil {
		utils.Error("Could not copy files. Is the permission granted?", true)
	}
	color.Green(" done")
	// Delete files from source dir if stash flag is set
	if flag_stash {
		fmt.Print("Stashing ...")
		for _, f := range saved_files {
			fmt.Println(f)
			os.RemoveAll(f)
		}
		color.Green(" done")
	}
}

func LoadTemplate(name string) {
	target_dir := utils.GetTemplatesPath() + "/" + name
	// Check if the template exists
	if name != "" && !utils.FileExists(target_dir) {
		utils.Error("Template with the name '"+name+"' could not be found. You can query a list of the templates by executing 'compose-generator template load'.", true)
	} else if name == "" {
		// Show list of saved templates

	}

}
