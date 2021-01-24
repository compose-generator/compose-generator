package commands

import (
	"fmt"
	"os"

	"compose-generator/utils"

	"github.com/fatih/color"
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
		color.Red("Error. You must specify a project name!")
		os.Exit(1)
	}
	fmt.Println(project_name)

	// Docker Swarm compatability (default: no)
	docker_swarm := utils.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	fmt.Println(docker_swarm)

	// Predefined stack (default: yes)
	use_predefined_stack := utils.YesNoQuestion("Do you want to use a predefined stack?", true)
	if use_predefined_stack {
		// Predefined stack menu
		items := []string{
			"LAMP (Linux, Apache, MySQL, PHP)",
			"MEAN (MongoDB, Express.js, Angular.js, Node.js)",
		}
		stack := utils.MenuQuestion("Predefined software stack", items)
		fmt.Println(stack)
	} else {
		// Create custom stack
		utils.Heading("Let's create a custom stack for you!")
	}

	// Generate files based on the answers
	process()
}

func process() {

}
