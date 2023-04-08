package sendsafelyuploader

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func chunkAndEncryptFiles(fileName string, uploader *Uploader) ([]string, error) {

	var fileNames []string

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	// stat the file so we can access properties of the file
	fileInfo, _ := file.Stat()
	var fileSize = int64(fileInfo.Size())

	// pick a file size and number of parts to divide it into.
	// we can add concurrency later to upload multiple parts at once
	// currently using 2.5 MB slices since that's what the sendsafely doc recommends, but this will definitely need to be adjust into much larger slices
	var fileChunkSize float64 = 2.5 * (1 << 20)

	// determine how many parts we'll need to chunk the file into based on the defined fileChunkSize
	totalPartNum := int64(math.Ceil(float64(fileSize) / float64(fileChunkSize)))

	log.Printf("Splitting to %d pieces.", totalPartNum)

	// split files into "totalPartsNum" number of files
	for i := int64(0); i < totalPartNum; i++ {

		partSize := int(math.Min(fileChunkSize, float64(fileSize-i*int64(fileChunkSize))))
		partBuffer := make([]byte, partSize)
		_, _ = file.Read(partBuffer)

		// write to disk and append the loop counter to each file part
		// hard coded for testing
		//outFileName := "split_files/testfile.txt_" + strconv.FormatUint(i+1, 10)
		outFileName := "split_files/" + filepath.Base(fileName) + "_" + strconv.FormatInt(i, 10)
		_, err := os.Create(outFileName)

		if err != nil {
			log.Fatal(err)
		}

		log.Print("Name of file part: ", outFileName)
		log.Print("Encrypting partBuffer")
		encryptedPartBufferReader, err := EncryptFileParts(uploader.PackageInfo.ServerSecret, uploader.ClientSecret, partBuffer)
		if err != nil {
			return nil, fmt.Errorf("Unable to encrypt file parts: %s", err)
		}

		buf := &bytes.Buffer{}
		_, err = buf.ReadFrom(encryptedPartBufferReader)
		if err != nil {
			return nil, err
		}

		// write/save buffer to disk
		err = ioutil.WriteFile(outFileName, buf.Bytes(), os.ModeAppend)
		if err != nil {
			return nil, err
		}

		fileNames = append(fileNames, outFileName)
		uploader.FileInfo.Parts = int(totalPartNum)
	}

	return fileNames, nil

}

func UploadLogs(args Args) error {

	uploader := CreateUploader("https://secure-upload.yugabyte.com", args.DropzoneIdFlag, "DROP_ZONE")

	// Step 1 - Create a new Dropzone Package
	if err := uploader.createDropzonePackage(); err != nil {
		return fmt.Errorf("Unable to create dropzone package: %s", err)
	}

	// Generate clientSecret and checksum
	uploader.CreateChecksum()

	for _, fileName := range args.FilesFlag {

		// Step 2 - Add a File to the Package
		// This step submits metadata about the file(s) to the SendSafely API
		// The actual upload is performed later in the workflow
		if err := uploader.addFileToPackage(fileName); err != nil {
			return fmt.Errorf("Unable to add file to package: %s", err)
		}

		// Step 3 - Obtain the Upload URLs for each File Part
		if err := uploader.getUploadURL(fileName); err != nil {
			return fmt.Errorf("Unable to get upload URL: %s", err)
		}

		// Step 4 - Encrypt and Upload each File Part
		// Files will be split, encrypted, and uploaded
		fileNames, err := chunkAndEncryptFiles(fileName, uploader)
		if err != nil {
			return fmt.Errorf("Unable to chunk and encrypt files: %s", err)
		}
		if err := uploader.uploadFilePartsToPackage(fileNames); err != nil {
			return fmt.Errorf("Unable to upload file parts: %s", err)
		}

		// Step 5 - Mark the Upload as Complete
		if err := uploader.markPackageComplete(); err != nil {
			return fmt.Errorf("Unable to upload file parts: %s", err)
		}

		// Step 6 - Finalize the Package
		if err := uploader.finalizePackage(); err != nil {
			return fmt.Errorf("Unable to finalize package: %s", err)
		}

		// Step 7 - Invoke the Hosted Dropzone Submission Endpoint
		if err := uploader.submitHostedDropzone(); err != nil {
			return fmt.Errorf("Unable to push file to Hosted Dropzone: %s", err)
		}
	}
	return nil
}
