package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/log"
	"main/uploader"
	"os"
)

var logger = log.Log()

var filesFlag []string
var caseNumFlag int = 0
var emailFlag string
var dropzoneIdFlag string

var rootCmd = &cobra.Command{
	Use:   "yb_log_uploader",
	Short: "utility to upload logs to yugabyte support",
	Run: func(cmd *cobra.Command, args []string) {
		isDropzoneFlagChanged := cmd.Flags().Changed("dropzone_id")
		uploader.UploadLogs(caseNumFlag, emailFlag, dropzoneIdFlag, isDropzoneFlagChanged, filesFlag)
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
	rootCmd.Flags().StringSliceVarP(&filesFlag, "files", "f", nil, "List of files to upload")
	rootCmd.Flags().IntVarP(&caseNumFlag, "case_num", "c", 0, "Zendesk case number to attach files to (required)")
	rootCmd.Flags().StringVarP(&emailFlag, "email", "e", "", "Email address of submitter (required)")
	rootCmd.Flags().StringVar(&dropzoneIdFlag, "dropzone_id", "S4dsLt2meOtq1iWgBhqJYsuEe2nzvYuv03j_Y6LqhY0", "Override default dropzone ID")

	rootCmd.MarkFlagRequired("case_num")
	rootCmd.MarkFlagRequired("email")
	rootCmd.Flags().MarkHidden("dropzone_id")
}
