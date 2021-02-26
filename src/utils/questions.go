package utils

import (
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

// Function for settings suggestions to a question for autocompletion
type suggest func(toComplete string) []string

// Print heading to console
func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

// Print simple text question
func TextQuestion(question string) (result string) {
	prompt := &survey.Input{
		Message: question,
	}
	survey.AskOne(prompt, &result)
	return
}

// Print simple text question with default value
func TextQuestionWithDefault(question string, default_value string) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: default_value,
	}
	survey.AskOne(prompt, &result)
	return
}

// Print simple text question with default value and a suggestion function
func TextQuestionWithSuggestions(question string, default_value string, sf suggest) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: default_value,
		Suggest: sf,
	}
	survey.AskOne(prompt, &result)
	return
}

// Print simple yes/no question with default value
func YesNoQuestion(question string, default_value bool) (result bool) {
	prompt := &survey.Confirm{
		Message: question,
		Default: default_value,
	}
	survey.AskOne(prompt, &result)
	return
}

// Prints a selection of predefined items
func MenuQuestion(label string, items []string) (result string) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// Prints a selection of predefined items and return the selected index
func MenuQuestionIndex(label string, items []string) (result int) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// Prints a multi selection of predefined items
func MultiSelectMenuQuestion(label string, items []string) (result []string) {
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// Prints an error message
func Error(description string, exit bool) {
	color.Red("Error: " + description)
	if exit {
		os.Exit(1)
	}
}
