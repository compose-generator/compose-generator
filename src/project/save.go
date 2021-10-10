/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/

package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"sort"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// SaveProject saves the Docker compose project to the current directory
func SaveProject(project *model.CGProject, options ...SaveOption) {
	// Apply options
	opt := applySaveOptions(options...)

	// Save all components
	saveCGFile(project, opt)
	saveGitignore(project, opt)
	saveReadme(project, opt)
	saveEnvFiles(project, opt)
	saveComposeFile(project, opt)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func saveComposeFile(project *model.CGProject, opt SaveOptions) {
	// Minify compose file
	project.Composition.WithoutUnnecessaryResources()
	// Remove unsupported options
	for serviceIndex, service := range project.Composition.Services {
		for volumeIndex := range service.Volumes {
			project.Composition.Services[serviceIndex].Volumes[volumeIndex].Bind = nil
		}
	}
	// Save docker compose file
	content, err := yaml.Marshal(project.Composition)
	if err != nil {
		util.Error("Could not save "+opt.ComposeFileName, err, true)
	}
	err = ioutil.WriteFile(opt.WorkingDir+opt.ComposeFileName, content, 0600)
	if err != nil {
		util.Error("Could not save "+opt.ComposeFileName, err, true)
	}
}

func saveEnvFiles(project *model.CGProject, opt SaveOptions) {
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
		// Sort env variables alphabetically
		keys := make([]string, 0, len(envVars))
		for key := range envVars {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		// Join env variables of this particular env file together
		content := ""
		for _, key := range keys {
			content += key + "=" + *envVars[key] + "\n"
		}
		// Replace vars
		content = util.ReplaceVarsInString(content, project.Vars)
		// Replace secrets
		secretMap := make(map[string]string)
		for _, secret := range project.Secrets {
			secretMap[secret.Variable] = secret.Value
		}
		content = util.ReplaceVarsInString(content, secretMap)
		// Write to disk
		if err := ioutil.WriteFile(opt.WorkingDir+fileName, []byte(content), 0600); err != nil {
			util.Error("Unable to write environment file '"+fileName+"' to the disk", err, true)
		}
	}
}

func saveGitignore(project *model.CGProject, opt SaveOptions) {
	if project.WithGitignore && len(project.GitignorePatterns) > 0 {
		// Create gitignore file with all the paths from the list
		content := ""
		for _, pattern := range project.GitignorePatterns {
			content += pattern + "\n"
		}
		if err := ioutil.WriteFile(opt.WorkingDir+".gitignore", []byte(content), 0600); err != nil {
			util.Error("Could not write .gitignore file", err, true)
		}
	}
}

func saveReadme(project *model.CGProject, opt SaveOptions) {
	if project.WithReadme && len(project.ReadmeChildPaths) > 0 {
		// Create Readme file, which consists of the content of all stated files
		content := ""
		for _, path := range project.ReadmeChildPaths {
			if util.FileExists(path) {
				// #nosec G304
				childContent, err := ioutil.ReadFile(path)
				if err != nil {
					util.Error("Could not load README.md from service template", err, false)
					continue
				}
				content += string(childContent) + "\n\n"
			}
		}
		// Replace vars
		content = util.ReplaceVarsInString(content, project.Vars)
		// Write to output file
		if err := ioutil.WriteFile(opt.WorkingDir+"README.md", []byte(content), 0600); err != nil {
			util.Error("Could not write README file", err, true)
		}
	}
}

func saveCGFile(project *model.CGProject, opt SaveOptions) {
	// Get some information
	project.LastModifiedBy = util.GetUsername()
	project.LastModifiedAt = time.Now().UnixNano()

	// Save config file
	config := viper.New()
	config.Set("project-name", project.Name)
	config.Set("project-container-name", project.ContainerName)
	config.Set("advanced-config", project.AdvancedConfig)
	config.Set("production-ready", project.ProductionReady)
	config.Set("created-by", project.CreatedBy)
	config.Set("created-at", project.CreatedAt)
	config.Set("modified-by", project.LastModifiedBy)
	config.Set("modified-at", project.LastModifiedAt)
	if err := config.WriteConfigAs(opt.WorkingDir + ".cg.yml"); err != nil {
		util.Error("Could not write CG config file", err, true)
	}
}
