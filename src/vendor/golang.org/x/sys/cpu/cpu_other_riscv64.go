// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && riscv64
<<<<<<< HEAD
<<<<<<< HEAD
=======
// +build !linux,riscv64
>>>>>>> fd0a574 (Bump github.com/compose-spec/compose-go from 1.2.9 to 1.3.0 in /src (#362))
=======
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

package cpu

func archInit() {
	Initialized = true
}
