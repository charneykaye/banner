// Based on https://github.com/Sirupsen/logrus
// Based on ssh/terminal:
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package banner

// These constants are declared here, rather than importing
// them from the syscall package as some syscall packages, even
// on linux, for example gccgo, do not declare them.
const ioctlReadTermios = 0x5401 // syscall.TCGETS