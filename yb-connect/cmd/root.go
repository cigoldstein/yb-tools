package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yugabyte/yb-tools/yb-connect/client"
)

var command string

var rootCmd = &cobra.Command{
	// simply print help if no arguments are given
	Use:   "yb-connect",
	Short: "Connect to one or more Yugabyte resources ",
}

var getHostname = &cobra.Command{
	Use:   "hostname [hostname or IP address]",
	Short: "Specify Yugaware host to connect to",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// there will only ever be one arg, so return args[0]
		client.YbConnect(args[0], command)
	},
}

func Execute() {

	rootCmd.AddCommand(getHostname)
	rootCmd.PersistentFlags().StringVarP(&command, "exec", "e", "", "Command to execute (Required)")
	err := rootCmd.MarkPersistentFlagRequired("exec")

	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
