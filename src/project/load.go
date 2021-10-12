/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package project

import (
	"compose-generator/model"
	"compose-generator/util"
	"io/ioutil"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/compose-spec/compose-go/loader"
	spec "github.com/compose-spec/compose-go/types"
	"github.com/spf13/viper"
)

var loadComposeFileMockable = loadComposeFile
var loadGitignoreFileMockable = loadGitignoreFile
var loadCGFileMockable = loadCGFile
var loadComposeFileSingleServiceMockable = loadComposeFileSingleService

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LoadProject loads a project from the disk
func LoadProject(options ...LoadOption) *model.CGProject {
	// Apply options
	opts := applyLoadOptions(options...)

	// Create project instance
	project := &model.CGProject{
		CGProjectMetadata: model.CGProjectMetadata{
			WithGitignore: fileExists(opts.WorkingDir + ".gitignore"),
			WithReadme:    fileExists(opts.WorkingDir + "README.md"),
		},
		GitignorePatterns: []string{},
		ReadmeChildPaths:  []string{"README.md"},
		ForceConfig:       false,
	}

	// Load components
	loadComposeFileMockable(project, opts)
	loadGitignoreFileMockable(project, opts)
	loadCGFileMockable(&project.CGProjectMetadata, opts)

	return project
}

// LoadProjectMetadata loads only the metadata of a project from the disk
func LoadProjectMetadata(options ...LoadOption) *model.CGProjectMetadata {
	opts := applyLoadOptions(options...)

	// Create project metadata instance
	metadata := &model.CGProjectMetadata{
		WithGitignore: util.FileExists(opts.WorkingDir + ".gitignore"),
		WithReadme:    util.FileExists(opts.WorkingDir + "README.md"),
	}

	// Load metadata
	loadCGFileMockable(metadata, opts)

	return metadata
}

// LoadTemplateService loads a project as a single service
func LoadTemplateService(
	project *model.CGProject,
	selectedTemplates *model.SelectedTemplates,
	templateTypeName string,
	serviceName string,
	options ...LoadOption,
) *spec.ServiceConfig {
	opts := applyLoadOptions(options...)
	return loadComposeFileSingleServiceMockable(
		project,
		selectedTemplates,
		templateTypeName,
		serviceName,
		opts,
	)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func loadComposeFile(project *model.CGProject, opt LoadOptions) {
	// Check if file exists
	if !fileExists(opt.WorkingDir + opt.ComposeFileName) {
		printError("Compose file not found", nil, true)
	}
	// Parse compose file
	content, err := readFile(opt.WorkingDir + opt.ComposeFileName)
	if err != nil {
		printError("Unable to parse '"+opt.ComposeFileName+"'", err, true)
	}
	dict, err := parseCompositionYAML(content)
	if err != nil {
		printError("Unable to parse '"+opt.ComposeFileName+"' file", err, true)
	}

	configs := []spec.ConfigFile{
		{
			Filename: opt.ComposeFileName,
			Config:   dict,
		},
	}
	config := spec.ConfigDetails{
		WorkingDir:  opt.WorkingDir,
		ConfigFiles: configs,
	}
	project.Composition, err = loadComposition(config)
	if err != nil {
		printError("Could not load project from the current directory", err, true)
	}
}

func loadComposeFileSingleService(
	project *model.CGProject,
	selectedTemplates *model.SelectedTemplates,
	templateTypeName string,
	serviceName string,
	opt LoadOptions,
) *spec.ServiceConfig {
	if !util.FileExists(opt.WorkingDir + opt.ComposeFileName) {
		util.Error("Compose file not found in template "+templateTypeName+"-"+serviceName, nil, true)
	}
	// Evaluate conditional sections
	evaluated := util.EvaluateConditionalSectionsToString(
		opt.WorkingDir+opt.ComposeFileName,
		selectedTemplates,
		project.Vars,
	)
	// Replace vars
	evaluated = util.ReplaceVarsInString(evaluated, project.Vars)
	// Parse file contents to service
	serviceDict, err := loader.ParseYAML([]byte(evaluated))
	if err != nil {
		util.Error("Unable to unmarshal the evaluated version of '"+templateTypeName+"-"+serviceName+"'", err, true)
	}
	service, err := loader.LoadService(templateTypeName+"-"+serviceName, serviceDict, opt.WorkingDir, nil, true)
	if err != nil {
		util.Error("Unable to load '"+templateTypeName+"-"+serviceName+"'", err, true)
	}
	return service
}

func loadGitignoreFile(project *model.CGProject, opt LoadOptions) {
	if project.WithGitignore {
		// Load patterns from .gitignore file
		content, err := ioutil.ReadFile(opt.WorkingDir + ".gitignore")
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

func loadCGFile(metadata *model.CGProjectMetadata, opt LoadOptions) {
	// Get default config values
	defaultProjectName := path.Base(opt.WorkingDir)
	defaultContainerName := strings.ReplaceAll(strings.ToLower(defaultProjectName), " ", "-")
	defaultAdvancedMode := false
	defaultProductionReady := false
	defaultCreatedBy := "unknown"
	defaultModifiedBy := "unknown"
	if user, err := user.Current(); err == nil {
		defaultCreatedBy = user.Name
		defaultModifiedBy = user.Name
	}
	defaultCreatedAt := time.Now().UnixNano()
	defaultModifiedAt := time.Now().UnixNano()

	configFileName := ".cg.yml"
	if util.FileExists(opt.WorkingDir + configFileName) {
		config := viper.New()
		// Set default values
		config.SetDefault("project-name", defaultProjectName)
		config.SetDefault("project-container-name", defaultContainerName)
		config.SetDefault("advanced-config", defaultAdvancedMode)
		config.SetDefault("production-ready", defaultProductionReady)
		config.SetDefault("created-by", defaultCreatedBy)
		config.SetDefault("created-at", defaultCreatedAt)
		config.SetDefault("modified-by", defaultModifiedBy)
		config.SetDefault("modified-at", defaultModifiedAt)
		// Load config file
		config.SetConfigName(".cg")
		config.SetConfigType("yml")
		config.AddConfigPath(opt.WorkingDir)
		err := config.ReadInConfig()
		if err != nil {
			util.Error("Could not read '"+configFileName+"' file", err, true)
		}
		// Assign values
		metadata.Name = config.GetString("project-name")
		metadata.ContainerName = config.GetString("project-container-name")
		metadata.AdvancedConfig = config.GetBool("advanced-config")
		metadata.ProductionReady = config.GetBool("production-ready")
		metadata.CreatedBy = config.GetString("created-by")
		metadata.CreatedAt = config.GetInt64("created-at")
		metadata.LastModifiedBy = config.GetString("modified-by")
		metadata.LastModifiedAt = config.GetInt64("modified-at")
	} else {
		metadata.Name = defaultProjectName
		metadata.ContainerName = defaultContainerName
		metadata.AdvancedConfig = defaultAdvancedMode
		metadata.CreatedBy = defaultCreatedBy
		metadata.CreatedAt = defaultCreatedAt
		metadata.LastModifiedBy = defaultModifiedBy
		metadata.LastModifiedAt = defaultModifiedAt
	}
}
