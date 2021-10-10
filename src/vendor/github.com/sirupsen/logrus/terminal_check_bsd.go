/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build darwin dragonfly freebsd netbsd openbsd
// +build !js

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
