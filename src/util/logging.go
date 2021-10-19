package util

import (
	"log"
	"os"
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

	logfile, err := os.OpenFile(getLogfilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

func getLogfilePath() string {
	timestampString := time.Now().Format("2017-09-07 17:06:04.000000")
	timestampString = strings.ReplaceAll(timestampString, " ", "_")
	timestampString = strings.ReplaceAll(timestampString, "-", "_")
	timestampString = strings.ReplaceAll(timestampString, ":", "_")
	timestampString = strings.ReplaceAll(timestampString, ".", "_")
	logfileName := "log_" + timestampString + ".log"
	return getLogfilesPath() + "/" + logfileName
}
