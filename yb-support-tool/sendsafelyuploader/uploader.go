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
		encryptedPartBufferReader := EncryptFileParts(uploader.PackageInfo.ServerSecret, uploader.Secrets.ClientSecret, partBuffer)
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

	var Uploader Uploader

	Uploader.RequestInfo.SsApiKeyHeader = args.DropzoneIdFlag
	Uploader.RequestInfo.SsRequestApiHeader = "DROP_ZONE"

	// Step 1 - Create a new Dropzone Package
	createDropzonePackage(&Uploader)

	// Generate clientSecret and checksum
	CreateClientSecret(&Uploader)
	CreateChecksum(&Uploader)

	for _, fileName := range args.FilesFlag {

		// Step 2 - Add a File to the Package
		// This step submits metadata about the file(s) to the SendSafely API
		// The actual upload is performed later in the workflow
		addFileToPackage(fileName, &Uploader)

		// Step 3 - Obtain the Upload URLs for each File Part
		getUploadUrls(&Uploader)

		// Step 4 - Encrypt and Upload each File Part
		// Files will be split, encrypted, and uploaded
		fileNames := chunkAndEncryptFiles(fileName, &Uploader)
		uploadFilePartsToPackage(fileNames, &Uploader)

		// Step 5 - Mark the Upload as Complete
		markPackageComplete(&Uploader)

		// Step 6 - Finalize the Package
		finalizePackage(&Uploader)

		// Step 7 - Invoke the Hosted Dropzone Submission Endpoint
		submitHostedDropzone(&Uploader)

		log.Print("Done")
	}
}
