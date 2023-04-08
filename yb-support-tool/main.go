package main

import (
	"fmt"
	"time"

	"github.com/yugabyte/yb-tools/yb-support-tool/cmd"
)

func main() {
	start := time.Now()

	// executes cobra for command line interaction
	cmd.Execute()

	fmt.Print("Execution time: ", time.Since(start))
}
