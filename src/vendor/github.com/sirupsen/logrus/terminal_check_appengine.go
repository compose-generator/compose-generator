/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
