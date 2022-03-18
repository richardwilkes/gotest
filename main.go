package main

import (
	"os"

	"github.com/richardwilkes/toolbox/cmdline"
)

func main() {
	cmdline.CopyrightStartYear = "2016"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License 2.0"
	cl := cmdline.New(true)
	cl.Parse(os.Args[1:])
}
