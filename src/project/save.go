package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// SaveProject saves the Docker compose project to the current directory
func SaveProject(project *model.CGProject, options ...SaveOption) {
	opt := applySaveOptions(options...)

	saveCGFile(project)
	saveGitignore(project)
	saveReadme(project)
	saveEnvironmentFiles(project)
	saveComposeFile(project, opt)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func saveComposeFile(project *model.CGProject, options SaveOptions) {
	// Save docker compose file
	content, err := yaml.Marshal(project.Composition)
	if err != nil {
		util.Error("Could not save "+options.ComposeFileName, err, true)
	}
	err = ioutil.WriteFile(options.ComposeFileName, content, 0755)
	if err != nil {
		util.Error("Could not save "+options.ComposeFileName, err, true)
	}
}

func saveEnvironmentFiles(project *model.CGProject) {
	// Make a list of all env files, which are listed in the project
	envFiles := make(map[string]map[string]*string)
	for _, service := range project.Composition.AllServices() {
		if len(service.EnvFile) > 0 {
			envFileName := service.EnvFile[0]
			// Initialize env file with empty map
			if _, ok := envFiles[envFileName]; !ok {
				envFiles[envFileName] = make(map[string]*string)
			}

			// Append line for each env var and delete the env var from the project
			for key, value := range service.Environment {
				envFiles[envFileName][key] = value
				delete(service.Environment, key)
			}
		}
	}
	// Write each file to the disk
	for fileName, envVars := range envFiles {
		// Join env variables of this particular env file together
		content := ""
		for key, value := range envVars {
			content += key + "=" + *value + "\n"
		}
		// Write to disk
		if err := ioutil.WriteFile(fileName, []byte(content), 0755); err != nil {
			util.Error("Unable to write environment file '"+fileName+"' to the disk", err, true)
		}
	}
}

func saveVolumes(project *model.CGProject) {
	// Make a list with all volume paths
}

func saveGitignore(project *model.CGProject) {
	if project.WithGitignore {
		// Create gitignore file with all the paths from the list
		content := ""
		for _, pattern := range project.GitignorePatterns {
			content += pattern + "\n"
		}
		ioutil.WriteFile(".gitignore", []byte(content), 0755)
	}
}

func saveReadme(project *model.CGProject) {
	if project.WithReadme {
		// Create Readme file, which consists of the content of all stated files
		content := ""
		for _, path := range project.ReadmeChildPaths {
			if util.FileExists(path) {
				childContent, err := ioutil.ReadFile(path)
				if err != nil {
					util.Error("Could not load README.md from service template", err, false)
					continue
				}
				content += string(childContent) + "\n"
			}
		}
		ioutil.WriteFile("README.md", []byte(content), 0755)
	}
}

func saveCGFile(project *model.CGProject) {
	viper.Set("project-name", project.Name)
	viper.Set("project-container-name", project.ContainerName)
	viper.Set("advanced-config", project.AdvancedConfig)
	viper.WriteConfigAs(".cg.yml")
}
