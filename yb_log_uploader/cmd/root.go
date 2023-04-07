package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/log"
	"main/structs"
	"main/uploader"
	"os"
)

var Args structs.Args

var Logger = log.CreateLogger(false, false)

var rootCmd = &cobra.Command{
	Use:   "yb_log_uploader",
	Short: "utility to upload logs to yugabyte support",
	Run: func(cmd *cobra.Command, args []string) {
		Args.IsDropzoneFlagChanged = cmd.Flags().Changed("dropzone_id")
		uploader.UploadLogs(Args)
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	Logger.Info("Processing command line arguments")
	rootCmd.Flags().StringSliceVarP(&Args.FilesFlag, "files", "f", nil, "List of files to upload")
	rootCmd.Flags().IntVarP(&Args.CaseNumFlag, "case_num", "c", 0, "Zendesk case number to attach files to (required)")
	rootCmd.Flags().StringVarP(&Args.EmailFlag, "email", "e", "", "Email address of submitter (required)")

	// default dropzone ID is set to Yugabyte Support's anonymous dropzone
	// this can be overridden with the --dropzone_id flag
	rootCmd.Flags().StringVar(&Args.DropzoneIdFlag, "dropzone_id", "BdFZz_JoZqtqPVueANkspD86KZ_PJsW1kIf_jVHeCO0", "Override default dropzone ID")

	rootCmd.MarkFlagRequired("files")
	rootCmd.MarkFlagRequired("case_num")
	rootCmd.MarkFlagRequired("email")

	// hide the dropzone_id flag from the help menu
	rootCmd.Flags().MarkHidden("dropzone_id")
}
