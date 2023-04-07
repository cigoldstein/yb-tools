package main

import (
	"main/cmd"
	"main/log"
	"time"
)

var Logger = log.CreateLogger(false, false)

func main() {
	start := time.Now()

	// executes cobra for command line interaction
	cmd.Execute()

	Logger.Info("Execution time: ", time.Since(start))
}
