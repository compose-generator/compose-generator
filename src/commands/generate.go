package commands

import (
	"fmt"

	question "compose-generator/utils"
)

func Generate() {
	// Welcome Message
	question.Heading("Welcome to Compose Generator!")
	fmt.Println("Please continue by answering a few questions:")
	fmt.Println()

	// Project name
	project_name := question.TextQuestion("What is the name of your project: ")
	if project_name == "" {
		fmt.Print("Error. You must specify a project name!")
		return
	}
	fmt.Println(project_name)

	// Docker Swarm compatability (default: no)
	docker_swarm := question.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	fmt.Println(docker_swarm)

	// Predefined stack (default: yes)
	use_predefined_stack := question.YesNoQuestion("Do you want to use a predefined stack?", true)
	if use_predefined_stack {
		// Predefined stack menu
		items := []string{
			"LAMP (Linux, Apache, MySQL, PHP)",
			"MEAN (MongoDB, Express.js, Angular.js, Node.js)",
		}
		stack := question.MenuQuestion("Predefined software stack", items)
		fmt.Println(stack)
	} else {
		// Create custom stack
		question.Heading("Let's create a custom stack for you!")
	}
}
