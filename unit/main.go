package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cloudfoundry/cnb-tools/utils"
)

func main() {
	fmt.Println("Run Buildpack Unit Tests")

	cmd := exec.Command("go", "test", "./...", "-v", "-run", "Unit")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	if err != nil {
		fmt.Printf(utils.RED, "GO Test Failed")
		os.Exit(utils.ExitCode(err))
	} else {
		fmt.Printf(utils.GREEN, "GO Test Succeeded")
	}
}
