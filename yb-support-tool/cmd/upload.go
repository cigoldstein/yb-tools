package cmd

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/docker/go-units"
	"github.com/schollz/progressbar/v3"
	uploader "github.com/yugabyte/yb-tools/yb-support-tool/sendsafelyuploader"
)

func addUploaderFlags() {
	// upload command flags
	uploadCmd.Flags().IntVarP(&caseNum, "case", "c", 0, "Zendesk case number to attach files to (required)")
	uploadCmd.Flags().StringVarP(&email, "email", "e", "", "Email address of submitter (required)")
	uploadCmd.Flags().IntVarP(&concurrency, "parallelism", "p", defaultConcurency, "parallelism")
	uploadCmd.Flags().IntVar(&retries, "retries", defaultRetries, "number file upload retry attempts")

	uploadCmd.MarkFlagRequired("case")
	uploadCmd.MarkFlagRequired("email")

	// TODO: validate # retries and concurrency

	// default dropzone ID is set to Yugabyte Support's anonymous dropzone
	// this can be overridden with the --dropzone_id flag
	uploadCmd.Flags().StringVar(&dropzoneID, "dropzone_id", YBDropzoneID, "Override default dropzone ID")
	// hide the dropzone_id flag from the help menu
	uploadCmd.Flags().MarkHidden("dropzone_id")

	// default uploader URL is set to Yugabyte Support's instance of sendsafely
	// this can be overridden with the --secure_upload_url flag
	uploadCmd.Flags().StringVar(&uploaderURL, "secure_upload_url", YBUploaderURL, "Override default secure uploader URL")
	// hide the dropzone_id flag from the help menu
	uploadCmd.Flags().MarkHidden("secure_upload_url")

	// partsize for testing
	uploadCmd.Flags().StringVar(&partSizeString, "partSize", "", "Part size used to upload file. Must be between 1KB and 2.5MB")
	uploadCmd.Flags().MarkHidden("partSize")
}

func Upload(email string, caseNum int, files []string) error {

	validateArgs()

	u := uploader.CreateUploader(uploaderURL, dropzoneID, "DROP_ZONE")

	p, err := u.CreateDropzonePackage(uploader.WithConcurrency(concurrency), uploader.WithRetries(retries), uploader.WithChunkSize(partSize))
	if err != nil {
		return fmt.Errorf("Unable to create dropzone package: %s", err)
	}

	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		fmt.Printf("Preparing to upload %s\n", fileName)
		f, err := p.AddFileToPackage(file)
		if err != nil {
			return fmt.Errorf("Unable to add file to package: %s", err)
		}

		bar := progressbar.Default(int64(f.Info.Parts), fmt.Sprintf("Uploading file: %s", f.Info.Name))
		go func() {
			for i := range f.Status {
				_ = bar.Add(i)
			}
		}()

		if err := p.UploadFileParts(f); err != nil {
			return fmt.Errorf("Unable to upload file parts: %s", err)
		}

		if err := p.MarkFileComplete(f); err != nil {
			return err
		}

	}

	fmt.Printf("File uploads complete\nFinalizing Package...\n")

	err = p.FinalizePackage()
	if err != nil {
		return fmt.Errorf("Unable to finalize package: %s", err)
	}

	if verbose {
		fmt.Printf("Package available at: %s\n", p.URL)
	}

	// Step 7 - Invoke the Hosted Dropzone Submission Endpoint
	if err := p.SubmitHostedDropzone(fmt.Sprint(caseNum), email); err != nil {
		return fmt.Errorf("Unable to push file to Hosted Dropzone: %s", err)
	}

	return nil
}

func validateArgs() {
	var err error
	if partSizeString != "" {
		partSize, err = units.FromHumanSize(partSizeString)
		if err != nil {
			printErrorAndExit("Unable to parse --partSize flag: %s\n", err)
		}
	} else {
		partSize = defaultPartSize
	}

	if concurrency < 1 || concurrency > 100 {
		fmt.Printf("Parallism must be a value between 1 and 100, reverting to default of %d\n", defaultConcurency)
		concurrency = defaultConcurency
	}

	if retries < 0 || retries > 100 {
		fmt.Printf("Retries must be a value between 0 and 100, reverting to default of %d\n", defaultRetries)
		retries = defaultRetries
	}

	if caseNum < 1 {
		printErrorAndExit("Case number must be a positive value\n")
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		printErrorAndExit("Invalid email provided: %s\n", email)
	}

}

func printErrorAndExit(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}
