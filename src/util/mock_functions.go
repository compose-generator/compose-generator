package util

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var println = fmt.Println
var white = color.White
var red = color.Red
var hiYellow = color.HiYellow
var printError = Error
var exitProgram = os.Exit
