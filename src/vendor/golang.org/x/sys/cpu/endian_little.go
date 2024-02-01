// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh || wasm
=======
//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh
// +build 386 amd64 amd64p32 alpha arm arm64 loong64 mipsle mips64le mips64p32le nios2 ppc64le riscv riscv64 sh
>>>>>>> b37f0a1 (Bump github.com/fatih/color from 1.14.1 to 1.15.0 in /src (#444))
=======
//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh || wasm
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

package cpu

// IsBigEndian records whether the GOARCH's byte order is big endian.
const IsBigEndian = false
