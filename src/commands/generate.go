package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/otiai10/copy"

	"compose-generator/parser"
	"compose-generator/utils"
)

func Generate() {
	// Execute SafetyFileChecks
	utils.ExecuteSafetyFileChecks()

	// Welcome Message
	utils.Heading("Welcome to Compose Generator!")
	fmt.Println("Please continue by answering a few questions:")
	fmt.Println()

	// Project name
	project_name := utils.TextQuestion("What is the name of your project: ")
	if project_name == "" {
		utils.Error("Error. You must specify a project name!", true)
	}

	// Docker Swarm compatability (default: no)
	docker_swarm := utils.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	fmt.Println(docker_swarm)

	// Predefined stack (default: yes)
	use_predefined_stack := utils.YesNoQuestion("Do you want to use a predefined stack?", true)
	if use_predefined_stack {
		// Load stacks from templates
		template_data := parser.ParseTemplates()
		// Predefined stack menu
		var items []string
		for _, t := range template_data {
			items = append(items, t.Label)
		}
		index, _ := utils.MenuQuestion("Predefined software stack", items)

		// Ask configured questions to the user
		env_map := make(map[string]string)
		env_map["PROJECT_NAME"] = project_name
		for _, q := range template_data[index].Questions {
			switch q.Type {
			case 1: // Yes/No
				default_value, _ := strconv.ParseBool(q.Default_value)
				env_map[q.Env_var] = strconv.FormatBool(utils.YesNoQuestion(q.Text, default_value))
			case 2: // Text
				env_map[q.Env_var] = utils.TextQuestionWithDefault(q.Text, q.Default_value)
			}
		}
		// Copy files and replace variables
		src_path := utils.GetTemplatesPath() + "/" + template_data[index].Dir
		dst_path := "."
		opt := copy.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasSuffix(src, "config.json") || strings.HasSuffix(src, "README.md") || strings.HasSuffix(src, ".gitkeep"), nil
			},
		}
		err := copy.Copy(src_path, dst_path, opt)
		if err != nil {
			utils.Error("Could not copy template files.", true)
		}
	} else {
		// Create custom stack
		utils.Heading("Let's create a custom stack for you!")
	}
}
