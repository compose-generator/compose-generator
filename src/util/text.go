package util

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

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

// Success prints a success message to console
func Success(text string) {
	green := color.New(color.FgGreen).Add(color.Italic)
	green.Println(text)
}

// Error prints an error message
func Error(description string, err error, exit bool) {
	color.Red(getErrorMessage(description, err))
	if exit {
		os.Exit(1)
	}
}

// Warning prints an warning message
func Warning(description string) {
	color.Red("Warning: " + description)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getErrorMessage(description string, err error) string {
	if err != nil {
		return "Error: " + description + ": " + err.Error()
	}
	return "Error: " + description
}
