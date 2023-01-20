// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc
<<<<<<< HEAD
=======
// +build darwin dragonfly freebsd linux,!ppc64,!ppc64le netbsd openbsd solaris
// +build gc
>>>>>>> cb35177 (Bump github.com/spf13/viper from 1.14.0 to 1.15.0 in /src (#417))

package unix

import "syscall"

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
