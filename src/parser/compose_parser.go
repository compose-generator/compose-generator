package parser

import (
	"compose-generator/util"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadProject loads the Docker compose project from the current directory
func LoadProject() *types.Project {
	content, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		util.Error("Unable to parse docker-compose.yml file", err, true)
	}
	dict, err := loader.ParseYAML(content)
	if err != nil {
		util.Error("Unable to parse docker-compose.yml file", err, true)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		util.Error("Unable to retrieve workdir", err, true)
	}

	configs := []types.ConfigFile{
		{
			Filename: "docker-compose.yml",
			Config:   dict,
		},
	}
	config := types.ConfigDetails{
		WorkingDir:  workingDir,
		ConfigFiles: configs,
		Environment: nil,
	}
	project, err := loader.Load(config)
	if err != nil {
		util.Error("Could not load project from the current directory", err, true)
	}
	return project
}

// SaveProject saves the Docker compose project to the current directory
func SaveProject(project *types.Project) {
	saveEnvironmentFiles(project)
	saveComposeFile(project)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func saveComposeFile(project *types.Project) {
	// Save docker compose file
	content, err := yaml.Marshal(project)
	if err != nil {
		util.Error("Could not save docker-compose.yml", err, true)
	}
	err = ioutil.WriteFile("docker-compose.yml", content, 0755)
	if err != nil {
		util.Error("Could not save docker-compose.yml", err, true)
	}
}

func saveEnvironmentFiles(project *types.Project) {
	// Make a list of all env files, which are listed in the project
	envFiles := make(map[string]map[string]*string)
	for _, service := range project.AllServices() {
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
