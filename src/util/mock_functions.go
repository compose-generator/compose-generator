/*
Copyright © 2021 Compose Generator Contributors
All rights reserved.
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
	"github.com/sethvargo/go-password/password"
)

// Logging
var logError = LogError
var logWarning = LogWarning

// Text output
var println = fmt.Println
var white = color.White

// Environment
var commandExists = CommandExists
var isDockerRunning = IsDockerRunning
var getEnv = os.Getenv
var executable = osext.Executable
var currentUser = user.Current
var isDockerizedEnvironment = IsDockerizedEnvironment
var isDevVersion = IsDevVersion
var isLinux = IsLinux
var isWindows = IsWindows
var getwd = os.Getwd
var getOuterVolumePathOnDockerizedEnvironmentMockable = getOuterVolumePathOnDockerizedEnvironment

// Conditions
var evaluateCondition = EvaluateCondition

// Other
var newClientWithOpts = client.NewClientWithOpts
var imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
	return cli.ImageList(ctx, opts)
}
var executeCommand = exec.Command
var getCommandOutput = func(cmd *exec.Cmd) ([]byte, error) {
	return cmd.Output()
}
var generatePassword = password.Generate
