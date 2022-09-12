package cmd

import (
	"github.com/spf13/cobra"

	"yb-get/log"
	"yb-get/platform"
	"yb-get/structs"
)

var ywInfo structs.YwInfo

var BundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Commands to perform bundle actions",
	Args:  cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		updateLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		platform.Logger = log.CreateLogger(log.Logger.Args.DebugFlag, log.Logger.Args.VerboseFlag)
		cmd.Help()
	},
}

var getBundleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get support bundle from Platform API",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		updateLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		ywInfo = getCustomersUniverses(ywInfo)
		platform.GetBundles(ywInfo)
	},
}

var createBundleCmd = &cobra.Command{
	Use:   "create",
	Short: "use GCP cloud",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		updateLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		ywInfo = getCustomersUniverses(ywInfo)
		platform.CreateBundle()
	},
}

func updateLogger() {
	platform.Logger = log.CreateLogger(log.Logger.Args.DebugFlag, log.Logger.Args.VerboseFlag)
}

func AddBundleCmds() {
	BundleCmd.AddCommand(getBundleCmd, createBundleCmd)
}

func AddBundleFlags() {
	defaultHost := "localhost:80"
	getBundleCmd.Flags().StringVar(&ywInfo.YwHost, "hostname", defaultHost, "Platform for support-bundle API calls")
}

func getCustomersUniverses(ywInfo structs.YwInfo) structs.YwInfo {

	ywInfo = platform.GetCustomers(ywInfo)
	ywInfo = platform.GetUniverses(ywInfo)

	// change loglevel if debug or verbose is set

	return ywInfo

}
