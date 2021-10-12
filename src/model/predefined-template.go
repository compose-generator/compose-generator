/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package model

// TemplateTypeFrontend, TemplateTypeBackend, TemplateTypeDatabase, TemplateTypeProxy and TemplateTypeTlsHelper represent the different template types
const (
	TemplateTypeFrontend  = "frontend"
	TemplateTypeBackend   = "backend"
	TemplateTypeDatabase  = "database"
	TemplateTypeDbAdmin   = "db-admin"
	TemplateTypeProxy     = "proxy"
	TemplateTypeTlsHelper = "tls-helper"
)

// FileTypeService, FileTypeEnv, FileTypeDocs and FileTypeConfig represent the different type options for template files
const (
	FileTypeService = "service"
	FileTypeEnv     = "env"
	FileTypeDocs    = "docs"
	FileTypeConfig  = "config"
)

// QuestionTypeYesNo, QuestionTypeText and QuestionTypeMenu represent the different question types for predefined service templates
const (
	QuestionTypeYesNo = 1
	QuestionTypeText  = 2
	QuestionTypeMenu  = 3
)

// SelectedTemplates represents predefined service templates, which were selected by the user to add to the stack
type SelectedTemplates struct {
	FrontendServices []PredefinedTemplateConfig `json:"frontend,omitempty"`
	BackendServices  []PredefinedTemplateConfig `json:"backend,omitempty"`
	DatabaseServices []PredefinedTemplateConfig `json:"database,omitempty"`
	DbAdminServices  []PredefinedTemplateConfig `json:"dbadmin,omitempty"`
	ProxyService     []PredefinedTemplateConfig `json:"proxy,omitempty"`
	TlsHelperService []PredefinedTemplateConfig `json:"tlshelper,omitempty"`
}

// GetAll returns all templates of all types, mixed in a single slice
func (t SelectedTemplates) GetAll() []PredefinedTemplateConfig {
	templates := []PredefinedTemplateConfig{}
	templates = append(templates, t.FrontendServices...)
	templates = append(templates, t.BackendServices...)
	templates = append(templates, t.DatabaseServices...)
	templates = append(templates, t.DbAdminServices...)
	templates = append(templates, t.ProxyService...)
	templates = append(templates, t.TlsHelperService...)
	return templates
}

// GetTotal returns the total number of selected templates
func (t SelectedTemplates) GetTotal() int {
	count := len(t.FrontendServices)
	count += len(t.BackendServices)
	count += len(t.DatabaseServices)
	count += len(t.DbAdminServices)
	count += len(t.ProxyService)
	count += len(t.TlsHelperService)
	return count
}

func (t SelectedTemplates) GetAllProxyQuestions() []Question {
	questions := []Question{}
	for _, template := range t.ProxyService {
		questions = append(questions, template.ProxyQuestions...)
	}
	for _, template := range t.TlsHelperService {
		questions = append(questions, template.ProxyQuestions...)
	}
	return questions
}

// AvailableTemplates represents all available predefined service templates
type AvailableTemplates struct {
	FrontendServices []PredefinedTemplateConfig
	BackendServices  []PredefinedTemplateConfig
	DatabaseServices []PredefinedTemplateConfig
	DbAdminServices  []PredefinedTemplateConfig
	ProxyService     []PredefinedTemplateConfig
	TlsHelperService []PredefinedTemplateConfig
}

// PredefinedTemplateConfig represents the JSON structure of predefined template configuration file
type PredefinedTemplateConfig struct {
	Label          string     `json:"label,omitempty"`
	Name           string     `json:"name,omitempty"`
	Dir            string     `json:"dir,omitempty"`
	Type           string     `json:"type,omitempty"`
	Preselected    string     `json:"preselected,omitempty"`
	Proxied        bool       `json:"proxied,omitempty"`
	DemoAppInitCmd []string   `json:"demoAppInitCmd,omitempty"`
	ServiceInitCmd []string   `json:"serviceInitCmd,omitempty"`
	Files          []File     `json:"files,omitempty"`
	Questions      []Question `json:"questions,omitempty"`
	ProxyQuestions []Question `json:"proxy-questions,omitempty"`
	Volumes        []Volume   `json:"volumes,omitempty"`
	Secrets        []Secret   `json:"secrets,omitempty"`
}

// GetFilePathsByType retrieves a slice with the file paths to all files of a template with a certain type
func (t PredefinedTemplateConfig) GetFilePathsByType(fileType string) []string {
	filteredFiles := []string{}
	for _, file := range t.Files {
		if file.Type == fileType {
			filteredFiles = append(filteredFiles, file.Path)
		}
	}
	return filteredFiles
}

// File represents an important file and holds the path and the type of this file
type File struct {
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
}

// Question represents the JSON structure of a question of a predefined template
type Question struct {
	Text         string   `json:"text,omitempty"`
	Type         int      `json:"type,omitempty"`
	DefaultValue string   `json:"defaultValue,omitempty"`
	Options      []string `json:"options,omitempty"`
	Validator    string   `json:"validator,omitempty"`
	Variable     string   `json:"variable,omitempty"`
	Advanced     bool     `json:"advanced,omitempty"`
}

// Volume represents the JSON structure of a volume of a predefined template
type Volume struct {
	Text         string `json:"text,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Variable     string `json:"variable,omitempty"`
	Advanced     bool   `json:"advanced,omitempty"`
}

// Secret represents the JSON structure of a secret of a predefined template
type Secret struct {
	Name     string `json:"name,omitempty"`
	Variable string `json:"variable,omitempty"`
	Length   int    `json:"length,omitempty"`
}
