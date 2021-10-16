/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var getErrorMessageMockable = getErrorMessage

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// P prints a normal text to the console
func P(text string) {
	if _, err := color.New(color.FgWhite).Print(text); err != nil {
		printError("Could not print white text", err, false)
	}
}

// Pl prints a normal text line to the console
func Pl(text string) {
	white(text)
}

// Pel prints an empty line to the console
func Pel() {
	if _, err := println(); err != nil {
		printError("Could not print empty line", err, true)
	}
}

// StartProcess displays a loading animation until StopProcess is called
func StartProcess(text string) (s *spinner.Spinner) {
	if IsCIEnvironment() {
		if _, err := color.New(color.FgWhite).Print(text); err != nil {
			Error("Could not print to console", err, false)
		}
		return nil
	}
	charSet := 14
	finalChar := "⠿"
	if runtime.GOOS == "windows" {
		charSet = 9
		finalChar = "-"
	}
	s = spinner.New(spinner.CharSets[charSet], 100*time.Millisecond)
	s.Suffix = " " + text
	s.FinalMSG = color.GreenString(finalChar) + " " + text + color.GreenString(" done\n")
	s.HideCursor = true
	s.Start()
	return
}

// StopProcess stops the spinner, which was started by calling StartProcess
func StopProcess(s *spinner.Spinner) {
	if IsCIEnvironment() {
		color.Green(" done")
	} else {
		s.Stop()
	}
}

// Heading prints heading to console
func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	if _, err := green.Println(text); err != nil {
		printError("Could not print heading", err, false)
	}
}

// Success prints a success message to console
func Success(text string) {
	green := color.New(color.FgGreen).Add(color.Italic)
	if _, err := green.Println(text); err != nil {
		printError("Could not print success message", err, false)
	}
}

// Error prints an error message
func Error(description string, err error, exit bool) {
	red(getErrorMessageMockable(description, err))
	if exit {
		exitProgram(1)
	}
}

// Warning prints an warning message
func Warning(description string) {
	hiYellow("Warning: " + description)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getErrorMessage(description string, err error) string {
	if err != nil {
		return "Error: " + description + ": " + err.Error()
	}
	return "Error: " + description
}
