package model

type Config struct {
	Label     string
	Dir       string
	Questions []Question
}

type Question struct {
	Question      string
	Type          int // 1 = Yes/No; 2 = Text
	Default_value string
	Env_var       string
}
