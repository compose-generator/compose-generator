package util

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// P prints a normal text to the console
func P(text string) {
	if _, err := color.New(color.FgWhite).Print(text); err != nil {
		Error("Could not print white text", err, false)
	}
}

// Pl prints a normal text line to the console
func Pl(text string) {
	color.White(text)
}

// Pel prints an empty line to the console
func Pel() {
	fmt.Println()
}

// StartProcess displays a loading animation until StopProcess is called
func StartProcess(text string) (s *spinner.Spinner) {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + text
	s.FinalMSG = color.GreenString("⠿") + " " + text + color.GreenString(" done\n")
	s.HideCursor = true
	s.Start()
	return
}

// StopProcess stops the spinner, which was started by calling StartProcess
func StopProcess(s *spinner.Spinner) {
	s.Stop()
}

// Heading prints heading to console
func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	if _, err := green.Println(text); err != nil {
		Error("Could not print heading", err, false)
	}
}

// Success prints a success message to console
func Success(text string) {
	green := color.New(color.FgGreen).Add(color.Italic)
	if _, err := green.Println(text); err != nil {
		Error("Could not print success message", err, false)
	}
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
