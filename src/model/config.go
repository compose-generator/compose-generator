package model

type Config struct {
	Label     string
	Questions []Question
}

type Question struct {
	Question     string
	Type         int // 1 = Yes/No; 2 = Text
	DefaultValue string
	EnvVar       string
}
