package cmd

import (
	"fmt"
	"os"

	uploader "github.com/yugabyte/yb-tools/yb-support-tool/sendsafelyuploader"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = CreateLogger(false, false)

var Version = "pre-release"

var Args uploader.Args

// globals

var (
	dropzoneUrl   = "https://secure-upload.yugabyte.com/drop-zone/v2.0/package/"
	SSUploaderURL = "https://secure-upload.yugabyte.com"
)

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

func init() {
	// root command flags
	rootCmd.Flags().StringVarP(&Args.EmailFlag, "email", "e", "", "Email address of submitter (required)")

	rootCmd.MarkFlagRequired("email")

	//subcommands
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(versionCmd)

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

// CreateLogger creates a logger
func CreateLogger(debugFlag bool, verboseFlag bool) zap.SugaredLogger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var defaultConsoleLogLevel zapcore.Level

	// configure verbosity and log-level
	switch {
	case debugFlag:
		defaultConsoleLogLevel = zap.DebugLevel
		config.FunctionKey = "true"
	case verboseFlag:
		defaultConsoleLogLevel = zap.InfoLevel
		config.FunctionKey = "true"
	default:
		defaultConsoleLogLevel = zap.InfoLevel
		config.TimeKey = ""
		config.CallerKey = ""
	}

	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultConsoleLogLevel)
	zapNew := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.FatalLevel))
	return *zapNew.Sugar()
}
