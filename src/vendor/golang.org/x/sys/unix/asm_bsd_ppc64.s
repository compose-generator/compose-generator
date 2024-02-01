// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || freebsd || netbsd || openbsd) && gc
<<<<<<< HEAD
<<<<<<< HEAD
=======
// +build darwin freebsd netbsd openbsd
// +build gc
>>>>>>> cb35177 (Bump github.com/spf13/viper from 1.14.0 to 1.15.0 in /src (#417))
=======
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

#include "textflag.h"

//
// System call support for ppc64, BSD
//

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT	·Syscall(SB),NOSPLIT,$0-56
	JMP	syscall·Syscall(SB)

TEXT	·Syscall6(SB),NOSPLIT,$0-80
	JMP	syscall·Syscall6(SB)

TEXT	·Syscall9(SB),NOSPLIT,$0-104
	JMP	syscall·Syscall9(SB)

TEXT	·RawSyscall(SB),NOSPLIT,$0-56
	JMP	syscall·RawSyscall(SB)

TEXT	·RawSyscall6(SB),NOSPLIT,$0-80
	JMP	syscall·RawSyscall6(SB)
