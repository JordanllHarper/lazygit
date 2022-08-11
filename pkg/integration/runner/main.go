package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/jesseduffield/generics/slices"
	"github.com/jesseduffield/lazygit/pkg/integration"
	"github.com/jesseduffield/lazygit/pkg/integration/helpers"
)

// see docs/Integration_Tests.md
// This file can be invoked directly, but you might find it easier to go through
// test/lazyintegration/main.go, which provides a convenient gui wrapper to integration tests.
//
// If invoked directly, you can specify a test by passing it as the first argument.
// You can also specify that you want to record a test by passing MODE=record
// as an env var.

func main() {
	mode := integration.GetModeFromEnv()
	includeSkipped := os.Getenv("INCLUDE_SKIPPED") == "true"
	var testsToRun []*helpers.Test

	if len(os.Args) > 1 {
	outer:
		for _, testName := range os.Args[1:] {
			// check if our given test name actually exists
			for _, test := range integration.Tests {
				if test.Name() == testName {
					testsToRun = append(testsToRun, test)
					continue outer
				}
			}
			log.Fatalf("test %s not found. Perhaps you forgot to add it to `pkg/integration/integration_tests/tests.go`?", testName)
		}
	} else {
		testsToRun = integration.Tests
	}

	testNames := slices.Map(testsToRun, func(test *helpers.Test) string {
		return test.Name()
	})

	err := integration.RunTestsNew(
		log.Printf,
		runCmdInTerminal,
		func(test *helpers.Test, f func() error) {
			if !slices.Contains(testNames, test.Name()) {
				return
			}
			if err := f(); err != nil {
				log.Print(err.Error())
			}
		},
		mode,
		includeSkipped,
	)
	if err != nil {
		log.Print(err.Error())
	}
}

func runCmdInTerminal(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
