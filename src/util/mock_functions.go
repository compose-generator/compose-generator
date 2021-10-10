/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package util

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
	"github.com/kardianos/osext"
)

var println = fmt.Println
var white = color.White
var red = color.Red
var hiYellow = color.HiYellow
var printError = Error
var printWarning = Warning
var commandExists = CommandExists
var isDockerRunning = IsDockerRunning
var exitProgram = os.Exit
var getEnv = os.Getenv
var fileExists = FileExists
var executable = osext.Executable
var currentUser = user.Current
var newClientWithOpts = client.NewClientWithOpts
var imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
	return cli.ImageList(ctx, opts)
}
var executeCommand = exec.Command
var getCommandOutput = func(cmd *exec.Cmd) ([]byte, error) {
	return cmd.Output()
}
