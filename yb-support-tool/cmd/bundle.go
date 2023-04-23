package cmd

import (
	"fmt"

	bundle "github.com/yugabyte/yb-tools/yb-support-tool/supportbundle"

	"github.com/spf13/cobra"
)

var ywInfo bundle.YwInfo

var BundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Commands to perform bundle actions",
	Args:  cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		updateLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// platform.Logger = log.CreateLogger(log.Logger.Args.DebugFlag, log.Logger.Args.VerboseFlag)
		cmd.Help()
	},
}

var getBundleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get support bundle from Platform API",
	Args:  cobra.NoArgs,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	updateLogger()
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// ywInfo = getCustomersUniverses(ywInfo)
		// platform.GetBundles(ywInfo)
		fmt.Print("Get log bundle")
	},
}

var createBundleCmd = &cobra.Command{
	Use:   "create",
	Short: "Create support bundle from Platform API",
	Args:  cobra.NoArgs,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	updateLogger()
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// ywInfo = getCustomersUniverses(ywInfo)
		// platform.CreateBundle()
		fmt.Print("Create log bundle")
	},
}

func updateLogger() {
	// platform.Logger = log.CreateLogger(log.Logger.Args.DebugFlag, log.Logger.Args.VerboseFlag)
}

func addBundleSubcommands() {
	BundleCmd.AddCommand(getBundleCmd, createBundleCmd)
}

func addBundleFlags() {
	defaultHost := "localhost:80"
	getBundleCmd.Flags().StringVar(&ywInfo.YwHost, "hostname", defaultHost, "Platform for support-bundle API calls")
}

func getCustomersUniverses(ywInfo bundle.YwInfo) bundle.YwInfo {

	// ywInfo = platform.GetCustomers(ywInfo)
	// ywInfo = platform.GetUniverses(ywInfo)

	// change loglevel if debug or verbose is set

	return ywInfo

}
