// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !netbsd && !openbsd && arm64
<<<<<<< HEAD
=======
// +build !linux,!netbsd,!openbsd,arm64
>>>>>>> 69ead24 (Bump github.com/spf13/viper from 1.13.0 to 1.14.0 in /src (#397))

package cpu

func doinit() {}
