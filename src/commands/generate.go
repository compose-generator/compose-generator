package commands

import (
	"fmt"
	"strconv"

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
				default_value, _ := strconv.ParseBool(q.DefaultValue)
				env_map[q.EnvVar] = strconv.FormatBool(utils.YesNoQuestion(q.Question, default_value))
			case 2: // Text
				env_map[q.EnvVar] = utils.TextQuestionWithDefault(q.Question, q.DefaultValue)
			}
		}
	} else {
		// Create custom stack
		utils.Heading("Let's create a custom stack for you!")
	}

	// Generate files based on the answers
	process()
}

func process() {

}
