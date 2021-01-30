package commands

import (
	"compose-generator/utils"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

func SaveTemplate(name string) {
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
	opt := copy.Options{
		Skip: func(src string) (bool, error) {
			return !strings.HasSuffix(src, "docker-compose.yml") && !strings.HasSuffix(src, "environment.env") && !strings.Contains(src, "volumes"), nil
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
}

func LoadTemplate(name string) {

}
