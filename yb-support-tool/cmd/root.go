package cmd

import (
	"fmt"
	"os"

	"github.com/docker/go-units"

	"github.com/spf13/cobra"
)

var Version = "pre-release"

// uploader variables
var (
	verbose bool

	// dropzone details
	dropzoneID  string
	uploaderURL string

	// uploader variables
	caseNum int
	email   string

	// package settings
	concurrency    int
	retries        int
	partSize       int64
	partSizeString string
)

const (
	defaultRetries = 5
	// testing shows these to be a good balance of speed and lower memory / CPU usage
	defaultConcurency = 10
	defaultPartSize   = 10 * units.MiB
)

// globals
const (
	YBUploaderURL = "https://secure-upload.yugabyte.com"
	YBDropzoneID  = "BdFZz_JoZqtqPVueANkspD86KZ_PJsW1kIf_jVHeCO0"
)

var (
	rootCmd = &cobra.Command{
		Use:   "yb-support-tool",
		Short: "Yugabyte Support Tool",
	}
	uploadCmd = &cobra.Command{
		Use:   "upload -c [case number] -e [email] [files]",
		Short: "Upload attachment to a support case",
		Long:  `Uploads a file or files up to a limit of 100GB to a support ticket. Requires exising ticket to be created`,
		Args:  cobra.RangeArgs(1, 10),
		Run: func(cmd *cobra.Command, args []string) {
			if err := Upload(email, caseNum, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
			os.Exit(0)
		},
	}
)

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

// nolint: errcheck
func init() {
	// root command flags

	//subcommands
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(BundleCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Print verbose logs")

	// uploader flags
	addUploaderFlags()

	// bundle subcommands cmds, get and create and flags
	addBundleSubcommands()
	addBundleFlags()

}
