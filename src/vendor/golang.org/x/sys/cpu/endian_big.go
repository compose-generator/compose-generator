// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build armbe || arm64be || m68k || mips || mips64 || mips64p32 || ppc || ppc64 || s390 || s390x || shbe || sparc || sparc64
<<<<<<< HEAD
<<<<<<< HEAD
=======
// +build armbe arm64be m68k mips mips64 mips64p32 ppc ppc64 s390 s390x shbe sparc sparc64
>>>>>>> b37f0a1 (Bump github.com/fatih/color from 1.14.1 to 1.15.0 in /src (#444))
=======
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

package cpu

// IsBigEndian records whether the GOARCH's byte order is big endian.
const IsBigEndian = true
