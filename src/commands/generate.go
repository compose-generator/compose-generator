package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/otiai10/copy"

	"compose-generator/parser"
	"compose-generator/utils"
)

func Generate(flag_advanced bool, flag_run bool, flag_demonized bool, flag_force bool) {
	// Execute SafetyFileChecks
	if !flag_force {
		utils.ExecuteSafetyFileChecks()
	}

	// Welcome Message
	utils.Heading("Welcome to Compose Generator!")
	fmt.Println("Please continue by answering a few questions:")
	fmt.Println()

	// Project name
	project_name := utils.TextQuestion("What is the name of your project: ")
	if project_name == "" {
		utils.Error("Error. You must specify a project name!", true)
	}
	project_name_container := strings.ReplaceAll(strings.ToLower(project_name), " ", "-")

	// Docker Swarm compatability (default: no)
	docker_swarm := utils.YesNoQuestion("Should your compose file be used for distributed deployment with Docker Swarm?", false)
	fmt.Println(docker_swarm)

	// Predefined stack (default: yes)
	use_predefined_stack := utils.YesNoQuestion("Do you want to use a predefined stack?", true)
	if use_predefined_stack {
		// Load stacks from templates
		template_data := parser.ParsePredefinedTemplates()
		// Predefined stack menu
		var items []string
		for _, t := range template_data {
			items = append(items, t.Label)
		}
		index, _ := utils.MenuQuestion("Predefined software stack", items)
		fmt.Println()

		// Ask configured questions to the user
		envMap := make(map[string]string)
		envMap["PROJECT_NAME"] = project_name
		envMap["PROJECT_NAME_CONTAINER"] = project_name_container
		for _, q := range template_data[index].Questions {
			if !q.Advanced || (q.Advanced && flag_advanced) {
				switch q.Type {
				case 1: // Yes/No
					default_value, _ := strconv.ParseBool(q.Default_value)
					envMap[q.Env_var] = strconv.FormatBool(utils.YesNoQuestion(q.Text, default_value))
				case 2: // Text
					envMap[q.Env_var] = utils.TextQuestionWithDefault(q.Text, q.Default_value)
				}
			} else {
				envMap[q.Env_var] = q.Default_value
			}
		}
		fmt.Println()
		// Ask for custom volume paths
		volumesMap := make(map[string]string)
		for _, v := range template_data[index].Volumes {
			if !v.Advanced || (v.Advanced && flag_advanced) {
				envMap[v.Env_var] = utils.TextQuestionWithDefault(v.Text, v.Default_value)
			} else {
				envMap[v.Env_var] = v.Default_value
			}
			volumesMap[v.Default_value] = envMap[v.Env_var]
		}

		// Copy template files
		fmt.Print("Copying template ...")
		src_path := utils.GetPredefinedTemplatesPath() + "/" + template_data[index].Dir
		dst_path := "."

		os.Remove(dst_path + "/docker-compose.yml")
		os.Remove(dst_path + "/environment.env")

		err1 := copy.Copy(src_path+"/docker-compose.yml", dst_path+"/docker-compose.yml")
		err2 := copy.Copy(src_path+"/environment.env", dst_path+"/environment.env")
		if err1 != nil || err2 != nil {
			utils.Error("Could not copy template files.", true)
		}

		color.Green(" done")

		// Create volumes
		fmt.Print("Creating volumes ...")

		for src, dst := range volumesMap {
			os.RemoveAll(dst)
			src = src_path + src[1:]

			opt := copy.Options{
				Skip: func(src string) (bool, error) {
					return strings.HasSuffix(src, ".gitkeep"), nil
				},
				OnDirExists: func(src string, dst string) copy.DirExistsAction {
					return copy.Replace
				},
			}
			err := copy.Copy(src, dst, opt)
			if err != nil {
				utils.Error("Could not copy volume files.", true)
			}
		}

		color.Green(" done")

		// Replace variables
		fmt.Print("Applying customizations ...")
		utils.ReplaceVarsInFile("./docker-compose.yml", envMap)
		utils.ReplaceVarsInFile("./environment.env", envMap)
		color.Green(" done")

		// Generate secrets
		fmt.Print("Generating secrets ...")
		secretsMap := utils.GenerateSecrets("./environment.env", template_data[index].Secrets)
		color.Green(" done")
		// Print secrets to console
		fmt.Println()
		fmt.Println("Following secrets were automatically generated:")
		for key, secret := range secretsMap {
			fmt.Print("   " + key + ": ")
			color.Yellow(secret)
		}
	} else {
		// Create custom stack
		utils.Heading("Let's create a custom stack for you!")
	}

	// Run if the corresponding flag is set
	if flag_run || flag_demonized {
		utils.DockerComposeUp(flag_demonized)
	}
}
