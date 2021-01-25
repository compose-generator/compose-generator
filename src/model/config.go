package model

type Config struct {
	Label     string
	Questions []Question
}

type Question struct {
	Question string
	EnvVar   string
}
