package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry/cnb-tools/action"
)

const (
	USAGE   = "Usage:   install_tools.sh <version: optional>\n"
	EXAMPLE = "Example: install_tools.sh 0.0.9\n"
)

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, USAGE)
		fmt.Fprint(os.Stderr, EXAMPLE)
		os.Exit(0)
	}
}

func main() {
	action.InstallTools()
}
