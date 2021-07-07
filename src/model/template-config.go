package model

// ServiceTemplateConfig represents the JSON structure of predefined template configuration file
type ServiceTemplateConfig struct {
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

// File represents an important file and holds the path and the type of this file
type File struct {
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
}

// Question represents the JSON structure of a question of a predefined template
type Question struct {
	Text           string   `json:"text,omitempty"`
	Type           int      `json:"type,omitempty"` // 1 = Yes/No; 2 = Text
	DefaultValue   string   `json:"defaultValue,omitempty"`
	Options        []string `json:"options,omitempty"`
	Validator      string   `json:"validator,omitempty"`
	Variable       string   `json:"variable,omitempty"`
	Advanced       bool     `json:"advanced,omitempty"`
	WithDockerfile bool     `json:"withDockerfile,omitempty"`
}

// Volume represents the JSON structure of a volume of a predefined template
type Volume struct {
	Text           string `json:"text,omitempty"`
	DefaultValue   string `json:"defaultValue,omitempty"`
	Variable       string `json:"variable,omitempty"`
	Advanced       bool   `json:"advanced,omitempty"`
	WithDockerfile bool   `json:"withDockerfile,omitempty"`
}

// Secret represents the JSON structure of a secret of a predefined template
type Secret struct {
	Name     string `json:"name,omitempty"`
	Variable string `json:"variable,omitempty"`
	Length   int    `json:"length,omitempty"`
}
