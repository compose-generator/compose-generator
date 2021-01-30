package model

type TemplateConfig struct {
	Label     string
	Dir       string
	Questions []Question
	Secrets   []Secret
}

type Question struct {
	Text          string
	Type          int // 1 = Yes/No; 2 = Text
	Default_value string
	Env_var       string
	Advanced      bool
}

type Secret struct {
	Name   string
	Var    string
	Length int
}
