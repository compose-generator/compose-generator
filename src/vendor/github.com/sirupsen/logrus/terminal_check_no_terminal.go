/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build js nacl plan9

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
