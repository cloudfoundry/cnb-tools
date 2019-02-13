package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cloudfoundry/cnb-tools/utils"

	"github.com/cloudfoundry/cnb-tools/action"
)

const (
	INTEGRATION         = "integration"
	ENVPACK             = "PACK_VERSION"
	DEFAULT_BUILD_IMAGE = "cfbuildpacks/cflinuxfs3-cnb-experimental:build"
	DEFAULT_RUN_IMAGE   = "cfbuildpacks/cflinuxfs3-cnb-experimental:run"
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
	action.InstallTools(envPack)

	buildImage := os.Getenv("CNB_BUILD_IMAGE")
	if buildImage == "" {
		buildImage = DEFAULT_BUILD_IMAGE
	}

	runImage := os.Getenv("CNB_RUN_IMAGE")
	if runImage == "" {
		runImage = DEFAULT_RUN_IMAGE
	}

	for _, image := range []string{runImage, buildImage} {
		cmd := exec.Command("docker", "pull", image)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Run Buildpack Runtime Integration Tests")

	cmd := exec.Command("go", "test", "./integration/...", "-v", "-run", "Integration")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	if err != nil {
		fmt.Printf(utils.RED, "GO Test Failed")
		os.Exit(utils.ExitCode(err))
	} else {
		fmt.Printf(utils.GREEN, "GO Test Succeeded")
	}
}
