package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func Generate() {
	// Welcome Message
	heading("Welcome to Compose Generator!")
	fmt.Println("Please continue by answering a few questions:")
	fmt.Println()

	// Project name
	project_name := textQuestion("What is the name of your project: ")
	if project_name == "" {
		fmt.Print("Error. You must specify a project name!")
		return
	}
	fmt.Println(project_name)

	// Docker Swarm compatability (default: no)
	docker_swarm := yesNoQuestion("Should your compose file be used for Docker Swarm?", false)
	fmt.Println(docker_swarm)

	// Predefined stack (default: yes)
	use_predefined_stack := yesNoQuestion("Do you want to use a predefined stack?", true)
	if use_predefined_stack {
		// Predefined stack menu
		prompt := promptui.Select{
			Label: "Predefined software stack",
			Items: []string{
				"LAMP (Linux, Apache, MySQL, PHP)",
				"MEAN (MongoDB, Express.js, Angular.js, Node.js)",
			},
		}
		_, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	} else {
		// Create custom stack
		heading("Let's create a custom stack for you!")
	}
}

func heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

func textQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	cyan.Print(question)

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	if result_string == "" {
		fmt.Println("Error. This value is required!")
		os.Exit(1)
	}
	return result_string
}

func textQuestionWithDefault(question string, default_value string) string {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	cyan.Print(question)

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	if result_string != "" {
		return result_string
	} else {
		return default_value
	}
}

func yesNoQuestion(question string, default_value bool) bool {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	if default_value {
		cyan.Print(question + " [Y/n]: ")
	} else {
		cyan.Print(question + " [y/N]: ")
	}

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	result := default_value
	if result_string != "" {
		result = strings.ToLower(result_string) == "y"
	}
	return result
}
