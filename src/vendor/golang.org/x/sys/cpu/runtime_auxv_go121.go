// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21
<<<<<<< HEAD
<<<<<<< HEAD
=======
// +build go1.21
>>>>>>> b37f0a1 (Bump github.com/fatih/color from 1.14.1 to 1.15.0 in /src (#444))
=======
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

package cpu

import (
	_ "unsafe" // for linkname
)

//go:linkname runtime_getAuxv runtime.getAuxv
func runtime_getAuxv() []uintptr

func init() {
	getAuxvFn = runtime_getAuxv
}
