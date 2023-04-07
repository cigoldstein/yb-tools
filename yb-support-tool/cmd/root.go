package cmd

import (
	"fmt"
	"main/log"
	"main/structs"
	"main/uploader"
	"os"

	"github.com/spf13/cobra"
)

var Args structs.Args

var Logger = log.CreateLogger(false, false)

var Verion = "pre-release"

var (
	rootCmd = &cobra.Command{
		Use:   "yb-support-tool",
		Short: "Yugabyte Support Tool",
	}
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload attachment to a support case",
		Long:  `Uploads a file or files up to a limit of 100GB to a support ticket. Requires exising ticket to be created`,
		Run: func(cmd *cobra.Command, args []string) {
			Args.IsDropzoneFlagChanged = cmd.Flags().Changed("dropzone_id")
			uploader.UploadLogs(Args)
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

func init() {
	// root command flags
	rootCmd.Flags().StringVarP(&Args.EmailFlag, "email", "e", "", "Email address of submitter (required)")

	rootCmd.MarkFlagRequired("email")

	//subcommands
	rootCmd.AddCommand(uploadCmd)

	// upload command flags
	uploadCmd.Flags().StringSliceVarP(&Args.FilesFlag, "files", "f", nil, "List of files to upload")
	uploadCmd.Flags().IntVarP(&Args.CaseNumFlag, "case_num", "c", 0, "Zendesk case number to attach files to (required)")

	uploadCmd.MarkFlagRequired("files")
	uploadCmd.MarkFlagRequired("case_num")

	// default dropzone ID is set to Yugabyte Support's anonymous dropzone
	// this can be overridden with the --dropzone_id flag
	uploadCmd.Flags().StringVar(&Args.DropzoneIdFlag, "dropzone_id", "BdFZz_JoZqtqPVueANkspD86KZ_PJsW1kIf_jVHeCO0", "Override default dropzone ID")
	// hide the dropzone_id flag from the help menu
	uploadCmd.Flags().MarkHidden("dropzone_id")

}
