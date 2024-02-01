// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !aix && !linux && (ppc64 || ppc64le)
<<<<<<< HEAD
<<<<<<< HEAD
=======
// +build !aix
// +build !linux
// +build ppc64 ppc64le
>>>>>>> cb35177 (Bump github.com/spf13/viper from 1.14.0 to 1.15.0 in /src (#417))
=======
>>>>>>> 7a0c493 (Bump github.com/docker/docker in /src (#533))

package cpu

func archInit() {
	PPC64.IsPOWER8 = true
	Initialized = true
}
