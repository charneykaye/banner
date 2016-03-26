// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package autoload configure the banner loader with defaults
// Import the package. Thats it.
package autoload

import (
	"flag"
	"os"

	"github.com/dimiro1/banner"
)

func init() {
	var (
		filename  string
		isEnabled bool
	)

	flag.StringVar(&filename, "banner", "banner.txt", "banner.txt file")
	flag.BoolVar(&isEnabled, "show-banner", true, "print the banner?")

	flag.Parse()

	banner.Init(os.Stdout, isEnabled, filename)
}