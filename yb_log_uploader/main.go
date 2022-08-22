package main

import (
	"main/cmd"
	"main/log"
	"time"
)

var logger = log.Log()

func init() {
}

func main() {
	start := time.Now()
	logger.Info("Starting log upload")

	cmd.Execute()

	logger.Info("Execution time: ", time.Since(start))
}
