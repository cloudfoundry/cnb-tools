package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cnb-tools/install_tools"
)

const (
	INTEGRATION = "integration"
	ENVPACK     = "PACK_VERSION"
)

func main() {
	if _, err := os.Stat(INTEGRATION); os.IsNotExist(err) {
		fmt.Println("** WARNING ** No integration tests specified")
		os.Exit(0)
	}

	envPack := os.Getenv(ENVPACK)
	if envPack != "" {
		fmt.Println("Using the", ENVPACK, "environment variable")
	}

	install_tools.Run()

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error installing pack: %s\n", err.Error())
	}

	// TODO install specific pack version (grab it from environment variable)

}
