/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build !windows

package safeexec

import "os/exec"

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
