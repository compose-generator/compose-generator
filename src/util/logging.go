/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
*/

package util

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	consoleWarningLogger *log.Logger
	consoleErrorLogger   *log.Logger
	DebugLogger          *log.Logger
	InfoLogger           *log.Logger
	WarningLogger        *log.Logger
	ErrorLogger          *log.Logger
)

func init() {
	// Open logfile

	logfile, err := os.OpenFile(getLogfilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	// Create console loggers
	consoleWarningLogger = log.New(os.Stdout, color.HiYellowString("WARNING: "), log.Ldate|log.Ltime)
	consoleErrorLogger = log.New(os.Stderr, color.RedString("ERROR: "), log.Ldate|log.Ltime)

	// Create file loggers
	DebugLogger = log.New(logfile, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	InfoLogger = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	WarningLogger = log.New(logfile, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	ErrorLogger = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// Write header to file
	InfoLogger.Println("\n" + getLogFileHeader())
}

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// LogError prints an error message to the console
func LogError(message string, exit bool) {
	if exit {
		consoleErrorLogger.Fatalln(message)
	} else {
		consoleErrorLogger.Println(message)
	}
}

// LogWarning prints a warning message to the console
func LogWarning(message string) {
	consoleWarningLogger.Println(message)
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func getLogFileHeader() string {
	headerFields := []string{}
	headerFields = append(headerFields, "Version: "+strings.ReplaceAll(BuildVersion(Version, Commit, Date, BuiltBy), "\n", ", "))
	headerFields = append(headerFields, "Priviledged: "+strconv.FormatBool(IsPrivileged()))
	headerFields = append(headerFields, "Dockerized: "+strconv.FormatBool(isDockerizedEnvironment()))
	headerFields = append(headerFields, "CI: "+strconv.FormatBool(IsCIEnvironment()))
	return strings.Join(headerFields, "\n")
}

func getLogfilePath() string {
	// Create filename
	timestampString := time.Now().Format("2017-09-07 17:06:04.000000")
	timestampString = strings.ReplaceAll(timestampString, " ", "_")
	timestampString = strings.ReplaceAll(timestampString, "-", "_")
	timestampString = strings.ReplaceAll(timestampString, ":", "_")
	timestampString = strings.ReplaceAll(timestampString, ".", "_")
	logfileName := "log_" + timestampString + ".log"
	// Create logfile dir
	logfileDir := getLogfilesPath()
	if err := os.MkdirAll(logfileDir, 0700); err != nil {
		log.Fatal(err)
	}
	return logfileDir + "/" + logfileName
}
