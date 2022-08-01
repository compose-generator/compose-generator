// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || freebsd || netbsd || openbsd) && gc
<<<<<<< HEAD
=======
// +build darwin freebsd netbsd openbsd
// +build gc
>>>>>>> fd0a574 (Bump github.com/compose-spec/compose-go from 1.2.9 to 1.3.0 in /src (#362))

#include "textflag.h"

// System call support for RISCV64 BSD

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
