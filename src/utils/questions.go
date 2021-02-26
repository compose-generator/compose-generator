package utils

import (
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

type suggest func(toComplete string) []string

func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

func TextQuestion(question string) string {
	result := ""
	prompt := &survey.Input{
		Message: question,
	}
	survey.AskOne(prompt, &result)
	return result
}

func TextQuestionWithDefault(question string, default_value string) string {
	result := ""
	prompt := &survey.Input{
		Message: question,
		Default: default_value,
	}
	survey.AskOne(prompt, &result)
	return result
}

func TextQuestionWithSuggestions(question string, default_value string, sf suggest) string {
	result := ""
	prompt := &survey.Input{
		Message: question,
		Default: default_value,
		Suggest: sf,
	}
	survey.AskOne(prompt, &result)
	return result
}

func YesNoQuestion(question string, default_value bool) bool {
	result := default_value
	prompt := &survey.Confirm{
		Message: question,
		Default: default_value,
	}
	survey.AskOne(prompt, &result)
	return result
}

func MenuQuestion(label string, items []string) string {
	result := ""
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return result
}

func MenuQuestionIndex(label string, items []string) int {
	result := 0
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return result
}

func MultiSelectMenuQuestion(label string, items []string) []string {
	result := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return result
}

func Error(description string, exit bool) {
	color.Red("Error: " + description)
	if exit {
		os.Exit(1)
	}
}

// -------------------------- Skip bell sound output on select questions --------------------------

type bellSkipper struct{}

func (bs *bellSkipper) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (bs *bellSkipper) Close() error {
	return os.Stderr.Close()
}
