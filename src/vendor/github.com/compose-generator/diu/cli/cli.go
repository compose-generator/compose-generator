/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package cli

import (
	"os/exec"
	"strings"
)

// ExecuteAndWaitWithOutput executes a command and return the command output as string
func ExecuteAndWaitWithOutput(c ...string) string {
	cmd := exec.Command(c[0], c[1:]...)
	output, _ := cmd.CombinedOutput()
	return strings.TrimRight(string(output), "\r\n")
}
