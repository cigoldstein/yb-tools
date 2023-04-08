package sendsafelyuploader

import (
	"bytes"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func chunkAndEncryptFiles(fileName string, uploader *Uploader) []string {

	var fileNames []string

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	// stat the file so we can access properties of the file
	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()

	// pick a file size and number of parts to divide it into.
	// we can add concurrency later to upload multiple parts at once
	// currently using 2.5 MB slices since that's what the sendsafely doc recommends, but this will definitely need to be adjust into much larger slices
	const fileChunkSize = 2.5 * (1 << 20)

	// determine how many parts we'll need to chunk the file into based on the defined fileChunkSize
	totalPartNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunkSize)))

	log.Printf("Splitting to %d pieces.", totalPartNum)

	// split files into "totalPartsNum" number of files
	for i := uint64(0); i < totalPartNum; i++ {

		partSize := int(math.Min(fileChunkSize, float64(fileSize-int64(i*fileChunkSize))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)

		// write to disk and append the loop counter to each file part
		// hard coded for testing
		//outFileName := "split_files/testfile.txt_" + strconv.FormatUint(i+1, 10)
		outFileName := "split_files/" + filepath.Base(fileName) + "_" + strconv.FormatUint(i, 10)
		_, err := os.Create(outFileName)

		if err != nil {
			log.Fatal(err)
		}

		log.Print("Name of file part: ", outFileName)
		log.Print("Encrypting partBuffer")
		encryptedPartBufferReader := EncryptFileParts(uploader.PackageInfo.ServerSecret, uploader.ClientSecret, partBuffer)
		buf := &bytes.Buffer{}
		buf.ReadFrom(encryptedPartBufferReader)

		// write/save buffer to disk
		ioutil.WriteFile(outFileName, buf.Bytes(), os.ModeAppend)

		fileNames = append(fileNames, outFileName)
		uploader.FileInfo.Parts = int(totalPartNum)
	}

	return fileNames

}

func UploadLogs(args Args) {

	uploader := CreateUploader("https://secure-upload.yugabyte.com", args.DropzoneIdFlag, "DROP_ZONE")

	// Step 1 - Create a new Dropzone Package
	uploader.createDropzonePackage()

	// Generate clientSecret and checksum
	uploader.CreateChecksum()

	for _, fileName := range args.FilesFlag {

		// Step 2 - Add a File to the Package
		// This step submits metadata about the file(s) to the SendSafely API
		// The actual upload is performed later in the workflow
		uploader.addFileToPackage(fileName)

		// Step 3 - Obtain the Upload URLs for each File Part
		uploader.getUploadUrls()

		// Step 4 - Encrypt and Upload each File Part
		// Files will be split, encrypted, and uploaded
		fileNames := chunkAndEncryptFiles(fileName, uploader)
		uploadFilePartsToPackage(fileNames, uploader)

		// Step 5 - Mark the Upload as Complete
		markPackageComplete(uploader)

		// Step 6 - Finalize the Package
		finalizePackage(uploader)

		// Step 7 - Invoke the Hosted Dropzone Submission Endpoint
		submitHostedDropzone(uploader)

		log.Print("Done")
	}
}
