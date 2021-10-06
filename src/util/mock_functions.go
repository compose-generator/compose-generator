package util

import (
	"fmt"
	"os"
	"os/user"

	"github.com/docker/docker/client"
	"github.com/fatih/color"
	"github.com/kardianos/osext"
)

var println = fmt.Println
var white = color.White
var red = color.Red
var hiYellow = color.HiYellow
var printError = Error
var exitProgram = os.Exit
var getEnv = os.Getenv
var fileExists = FileExists
var executable = osext.Executable
var currentUser = user.Current
var newClientWithOpts = client.NewClientWithOpts
