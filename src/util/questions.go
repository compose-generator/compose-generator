package util

import (
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

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
func TextQuestionWithSuggestions(question string, fn Suggest) (result string) {
	prompt := &survey.Input{
		Message: question,
		Suggest: fn,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// TextQuestionWithDefaultAndSuggestions prints simple text question with default value and a suggestion function
func TextQuestionWithDefaultAndSuggestions(question string, defaultValue string, fn Suggest) (result string) {
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

func MenuQuestionWithDefault(label string, items []string, defaultItem string) (result string) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
		Default: defaultItem,
	}
	handleInterrupt(survey.AskOne(prompt, &result))
	return
}

// MenuQuestionIndex prints a selection of predefined items and returns the selected index
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
		Message:  label,
		Options:  items,
		PageSize: 15,
	}
	handleInterrupt(survey.AskOne(prompt, &result, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "yellow+hb"
	})))
	return
}

// MultiSelectMenuQuestionIndex prints a multi selection of predefined items and returns the selected indices
func MultiSelectMenuQuestionIndex(label string, items []string, defaultItems []string) (result []int) {
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  items,
		Default:  defaultItems,
		PageSize: 15,
	}
	handleInterrupt(survey.AskOne(prompt, &result, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "yellow+hb"
	})))
	return
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

// Function for settings suggestions to a question for autocompletion
type Suggest func(toComplete string) []string

func handleInterrupt(err error) {
	if err == terminal.InterruptErr {
		Pel()
		os.Exit(0)
	}
	if err != nil {
		panic(err)
	}
}
