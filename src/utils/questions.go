package utils

import (
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

// Function for settings suggestions to a question for autocompletion
type suggest func(toComplete string) []string

// Heading prints heading to console
func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

// TextQuestion prints simple text question
func TextQuestion(question string) (result string) {
	prompt := &survey.Input{
		Message: question,
	}
	survey.AskOne(prompt, &result)
	return
}

// TextQuestionWithDefault prints simple text question with default value
func TextQuestionWithDefault(question string, defaultValue string) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: defaultValue,
	}
	survey.AskOne(prompt, &result)
	return
}

// TextQuestionWithSuggestions prints simple text question with default value and a suggestion function
func TextQuestionWithSuggestions(question string, defaultValue string, sf suggest) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: defaultValue,
		Suggest: sf,
	}
	survey.AskOne(prompt, &result)
	return
}

// YesNoQuestion prints simple yes/no question with default value
func YesNoQuestion(question string, defaultValue bool) (result bool) {
	prompt := &survey.Confirm{
		Message: question,
		Default: defaultValue,
	}
	survey.AskOne(prompt, &result)
	return
}

// MenuQuestion prints a selection of predefined items
func MenuQuestion(label string, items []string) (result string) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// MenuQuestionIndex prints a selection of predefined items and return the selected index
func MenuQuestionIndex(label string, items []string) (result int) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// MultiSelectMenuQuestion prints a multi selection of predefined items
func MultiSelectMenuQuestion(label string, items []string) (result []string) {
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	survey.AskOne(prompt, &result)
	return
}

// Error prints an error message
func Error(description string, exit bool) {
	color.Red("Error: " + description)
	if exit {
		os.Exit(1)
	}
}
