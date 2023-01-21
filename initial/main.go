package main

/**
A little script for init container, that copies the `runner` binary
to the shared volume for the application container to pick up
*/

import (
	"megacron/system/functions"
	"os"
)

func main() {
	functions.Log("Starting initialization..")
	initial()
	functions.Log("Initialization complete.")
}

// initial Initializing the side card daemon
func initial() {
	prepareRunner()
}

func prepareRunner() bool {
	err := functions.CopyFile(
		os.Getenv("RUNNER_BINARY_PATH"),
		os.Getenv("RUNNER_BINARY_WORK_PATH"),
		true,
	)
	return err == nil
}
