package model

// Json structure of predefined template configuration file
type TemplateConfig struct {
	Label     string
	Dir       string
	Questions []Question
	Volumes   []Volume
	Secrets   []Secret
}

// Json structure of a question of a predefined template
type Question struct {
	Text         string
	Type         int // 1 = Yes/No; 2 = Text
	DefaultValue string
	EnvVar       string
	Advanced     bool
}

// Json structure of a volume of a predefined template
type Volume struct {
	Text         string
	DefaultValue string
	EnvVar       string
	Advanced     bool
}

// Json structure of a secret of a predefined template
type Secret struct {
	Name   string
	Var    string
	Length int
}
