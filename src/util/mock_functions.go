package util

import (
	"context"
	"fmt"
	"os"
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
var exitProgram = os.Exit
var getEnv = os.Getenv
var fileExists = FileExists
var executable = osext.Executable
var currentUser = user.Current
var newClientWithOpts = client.NewClientWithOpts
var imageList = func(cli *client.Client, ctx context.Context, opts types.ImageListOptions) ([]types.ImageSummary, error) {
	return cli.ImageList(ctx, opts)
}
