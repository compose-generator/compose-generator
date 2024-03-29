package model

const (
	TemplateTypeFrontend  = "frontend"
	TemplateTypeBackend   = "backend"
	TemplateTypeDatabase  = "database"
	TemplateTypeDbAdmin   = "db-admin"
	TemplateTypeProxy     = "proxy"
	TemplateTypeTlsHelper = "tls-helper"
)

const (
	QuestionTypeYesNo = 1
	QuestionTypeText  = 2
	QuestionTypeMenu  = 3
)

// SelectedTemplate represents predefined service templates, which were selected by the user to add to the stack
type SelectedTemplates struct {
	FrontendServices []PredefinedTemplateConfig `json:"frontend,omitempty"`
	BackendServices  []PredefinedTemplateConfig `json:"backend,omitempty"`
	DatabaseServices []PredefinedTemplateConfig `json:"database,omitempty"`
	DbAdminServices  []PredefinedTemplateConfig `json:"dbadmin,omitempty"`
	ProxyService     []PredefinedTemplateConfig `json:"proxy,omitempty"`
	TlsHelperService []PredefinedTemplateConfig `json:"tlshelper,omitempty"`
}

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
	DemoAppInitCmd []string   `json:"demoAppInitCmd,omitempty"`
	ServiceInitCmd []string   `json:"serviceInitCmd,omitempty"`
	Files          []File     `json:"files,omitempty"`
	Questions      []Question `json:"questions,omitempty"`
	Volumes        []Volume   `json:"volumes,omitempty"`
	Secrets        []Secret   `json:"secrets,omitempty"`
}

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
	Type         int      `json:"type,omitempty"` // 1 = Yes/No; 2 = Text
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
