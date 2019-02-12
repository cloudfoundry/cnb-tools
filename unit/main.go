package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const RED = "\n\033[0;31m%s\033[0m\n"
const GREEN = "\n\033[0;32m%s\033[0m\n"

func main() {
	fmt.Println("Run Buildpack Unit Tests")

	cmd := exec.Command("go", "test", "./...", "-v", "-run", "Unit")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	if err != nil {
		fmt.Printf(RED, "GO Test Failed")
		os.Exit(ExitCode(err))
	} else {
		fmt.Printf(GREEN, "GO Test Succeeded")
	}
}

func ExitCode(err error) int {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		}
	}
	return 1
}
