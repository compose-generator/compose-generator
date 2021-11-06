/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// P prints a normal text to the console
func P(text string) {
	if _, err := color.New(color.FgWhite).Print(text); err != nil {
		ErrorLogger.Println("Could not print white text: " + err.Error())
		logError("Could not print white text", false)
	}
}

// Pl prints a normal text line to the console
func Pl(text string) {
	white(text)
}

// Pel prints an empty line to the console
func Pel() {
	if _, err := println(); err != nil {
		ErrorLogger.Println("Could not print empty line: " + err.Error())
		logError("Could not print empty line", true)
	}
}

// StartProcess displays a loading animation until StopProcess is called
func StartProcess(text string) (s *spinner.Spinner) {
	if IsCIEnvironment() {
		if _, err := color.New(color.FgWhite).Print(text); err != nil {
			ErrorLogger.Println("Could not print process start: " + err.Error())
			logError("Could not print to console", false)
		}
		return nil
	}
	charSet := 14
	finalChar := "⠿"
	if isWindows() {
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
		ErrorLogger.Println("Could not print heading: " + err.Error())
		logError("Could not print heading", false)
	}
}

// Success prints a success message to console
func Success(text string) {
	green := color.New(color.FgGreen).Add(color.Italic)
	if _, err := green.Println(text); err != nil {
		ErrorLogger.Println("Could not print success message: " + err.Error())
		logError("Could not print success message", false)
	}
}

// IsUrl checks if the input is a valid url: true = valid, false = invalid
func IsUrl(str string) bool {
	url, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	if net.ParseIP(url.Host) == nil {
		return strings.Contains(url.Host, ".")
	}
	return true
}
