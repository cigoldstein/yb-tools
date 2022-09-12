package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"yb-get/log"
)

var rootCmd = &cobra.Command{
	// simply print help if no arguments are given
	Use:   "yb-getlogs",
	Short: "Collect logs from Yugabyte systems",
}

func Execute() {

	// Persistent Flags
	// verbose adds additional information to the log lines (loglevel,
	rootCmd.PersistentFlags().BoolVarP(&log.Logger.Args.DebugFlag, "debug", "d", false, "Debug-level logging and verbose formatting")
	rootCmd.PersistentFlags().BoolVarP(&log.Logger.Args.VerboseFlag, "verbose", "v", false, "Verbose formatting")

	// bundle cmds, get and create
	rootCmd.AddCommand(BundleCmd)
	AddBundleCmds()
	AddBundleFlags()

	// grab auth token from env variable
	ywInfo.YwAuthToken = os.Getenv("YW_API_TOKEN")

	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
