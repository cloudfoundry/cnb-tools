package main

import (
	"fmt"
	"os"
)

const INTEGRATION = "integration"

func main() {
	if _, err := os.Stat(INTEGRATION); os.IsNotExist(err) {
		fmt.Println("** WARNING ** No integration tests specified")
		os.Exit(0)
	}

	// TODO install specific pack version (grab it from environment variable)

}
