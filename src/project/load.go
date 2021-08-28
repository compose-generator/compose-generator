package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/viper"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadProject loads the Docker compose project from the current directory
func LoadProject(options ...LoadOption) *model.CGProject {
	opts := applyLoadOptions(options...)

	// Create project instance
	project := &model.CGProject{
		WithGitignore:     util.FileExists(opts.WorkingDir + ".gitignore"),
		GitignorePatterns: []string{},
		WithReadme:        util.FileExists(opts.WorkingDir + "README.md"),
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
	}

	// Load components
	loadComposeFile(project, opts)
	loadGitignoreFile(project)
	loadCGFile(project)

	return project
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func loadComposeFile(project *model.CGProject, opt LoadOptions) {
	if !util.FileExists(opt.WorkingDir + opt.ComposeFileName) {
		util.Error("Compose file not found", nil, true)
	}
	content, err := ioutil.ReadFile(opt.WorkingDir + opt.ComposeFileName)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"'", err, true)
	}
	dict, err := loader.ParseYAML(content)
	if err != nil {
		util.Error("Unable to parse '"+opt.ComposeFileName+"' file", err, true)
	}

	configs := []types.ConfigFile{
		{
			Filename: opt.ComposeFileName,
			Config:   dict,
		},
	}
	config := types.ConfigDetails{
		WorkingDir:  opt.WorkingDir,
		ConfigFiles: configs,
	}
	project.Composition, err = loader.Load(config)
	if err != nil {
		util.Error("Could not load project from the current directory", err, true)
	}
}

func loadGitignoreFile(project *model.CGProject) {
	if project.WithGitignore {
		// Load patterns from .gitignore file
		content, err := ioutil.ReadFile(project.Composition.WorkingDir + ".gitignore")
		if err != nil {
			util.Error("Unable to parse .gitignore file", err, true)
		}
		contentStr := strings.ReplaceAll(string(content), "\r\n", "\n")
		// Save them into the project
		for _, line := range strings.Split(contentStr, "\n") {
			trimmedLine := strings.TrimSpace(line)
			if len(trimmedLine) > 0 && !strings.HasPrefix(trimmedLine, "#") {
				project.GitignorePatterns = append(project.GitignorePatterns, trimmedLine)
			}
		}
	}
}

func loadCGFile(project *model.CGProject) {
	// Get default config values
	workingDir, err := os.Getwd()
	if err != nil {
		util.Error("Unable to retrieve workdir", err, true)
	}
	defaultProjectName := path.Base(strings.ReplaceAll(workingDir, "\\", "/"))
	defaultContainerName := strings.ReplaceAll(strings.ToLower(defaultProjectName), " ", "-")
	defaultAdvancedMode := false

	configFileName := ".gc.yml"
	if util.FileExists(configFileName) {
		// Set default values
		viper.SetDefault("project-name", defaultProjectName)
		viper.SetDefault("project-container-name", defaultContainerName)
		viper.SetDefault("advanced-config", defaultAdvancedMode)
		// Load config file
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(".")
		err = viper.ReadInConfig()
		if err != nil {
			util.Error("Could not read '"+configFileName+"' file", err, true)
		}
		// Assign values
		project.Name = viper.GetString("project-name")
		project.ContainerName = viper.GetString("project-container-name")
		project.AdvancedConfig = viper.GetBool("advanced-config")
	} else {
		project.Name = defaultProjectName
		project.ContainerName = defaultContainerName
		project.AdvancedConfig = defaultAdvancedMode
	}
}
