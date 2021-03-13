package utils

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/fatih/color"
)

// Function for settings suggestions to a question for autocompletion
type suggest func(toComplete string) []string

// P prints a normal text to the console
func P(text string) {
	color.New(color.FgWhite).Print(text)
}

// Pl prints a normal text line to the console
func Pl(text string) {
	color.White(text)
}

// Pel prints an empty line to the console
func Pel() {
	fmt.Println()
}

// Done prints 'done' in green to the console
func Done() {
	color.Green("done")
}

// Heading prints heading to console
func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

// SuccessMessage prints a success message to console
func SuccessMessage(text string) {
	green := color.New(color.FgGreen).Add(color.Italic)
	green.Println(text)
}

// TextQuestion prints simple text question
func TextQuestion(question string) (result string) {
	prompt := &survey.Input{
		Message: question,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// TextQuestionWithValidator prints simple text question with a validation
func TextQuestionWithValidator(question string, fn survey.Validator) (result string) {
	prompt := &survey.Question{
		Prompt: &survey.Input{
			Message: question,
		},
		Validate: fn,
	}
	handleInterrupt(survey.Ask([]*survey.Question{prompt}, &result))
	return
}

// TextQuestionWithDefaultAndValidator prints simple text question with a validation and default value
func TextQuestionWithDefaultAndValidator(question string, defaultValue string, fn survey.Validator) (result string) {
	prompt := &survey.Question{
		Prompt: &survey.Input{
			Message: question,
			Default: defaultValue,
		},
		Validate: fn,
	}
	handleInterrupt(survey.Ask([]*survey.Question{prompt}, &result))
	return
}

// TextQuestionWithDefault prints simple text question with default value
func TextQuestionWithDefault(question string, defaultValue string) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: defaultValue,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// TextQuestionWithSuggestions prints simple text question with a suggestion function
func TextQuestionWithSuggestions(question string, fn suggest) (result string) {
	prompt := &survey.Input{
		Message: question,
		Suggest: fn,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// TextQuestionWithDefaultAndSuggestions prints simple text question with default value and a suggestion function
func TextQuestionWithDefaultAndSuggestions(question string, defaultValue string, fn suggest) (result string) {
	prompt := &survey.Input{
		Message: question,
		Default: defaultValue,
		Suggest: fn,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// YesNoQuestion prints simple yes/no question with default value
func YesNoQuestion(question string, defaultValue bool) (result bool) {
	prompt := &survey.Confirm{
		Message: question,
		Default: defaultValue,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// MenuQuestion prints a selection of predefined items
func MenuQuestion(label string, items []string) (result string) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// MenuQuestionIndex prints a selection of predefined items and return the selected index
func MenuQuestionIndex(label string, items []string) (result int) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// MultiSelectMenuQuestion prints a multi selection of predefined items
func MultiSelectMenuQuestion(label string, items []string) (result []string) {
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// Error prints an error message
func Error(description string, exit bool) {
	color.Red("Error: " + description)
	if exit {
		os.Exit(1)
	}
}

// Warning prints an warning message
func Warning(description string) {
	color.Red("Warning: " + description)
}

func handleInterrupt(err error) {
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
}
