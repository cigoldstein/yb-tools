package uploader

// TODO: create a "createHttpClient function so that there's less reused code on api calls
// TODO: break structs out into separate file

import (
	"bytes"
	"io/ioutil"
	"main/log"
	"main/structs"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

var Logger = log.CreateLogger(false, false)

func createHttpPut() {

}

func createHttpPost() {

}

func chunkAndEncryptFiles(path, fileName string, Uploader structs.Uploader) []string {

	var fileNames []string

	// hard code one file for now, will be a list coming from CLI
	fileToChunk := filepath.Join(path, fileName)

	// will be inaccurate until the hard-coded file above is removed
	Logger.Debug("Files to chunk: ", Uploader.Args.FilesFlag)

	file, err := os.Open(fileToChunk)

	if err != nil {
		Logger.Error(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 2.5 * (1 << 20)

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	Logger.Infof("Splitting to %d pieces.", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "split_files/testfile.txt_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			Logger.Error(err)
			os.Exit(1)
		}

		Logger.Info("Name of file part: ", fileName)
		// encrypt file part
		//unencryptedFilePart, err := os.ReadFile(fileName)
		//if err != nil {
		//	Logger.Error(err)
		//	os.Exit(1)
		//}

		Logger.Info("Encrypting partBuffer")
		encryptedPartBufferReader := EncryptFileParts(Uploader.PackageInfo.ServerSecret, Uploader.Secrets.ClientSecret, partBuffer)
		buf := &bytes.Buffer{}
		buf.ReadFrom(encryptedPartBufferReader)

		// write/save buffer to disk
		ioutil.WriteFile(fileName, buf.Bytes(), os.ModeAppend)

		fileNames = append(fileNames, fileName)

	}

	return fileNames

}

//
//func uploadFilePartsToPackage(fileNames []string, urlInfo uploadUrlInfo) {
//
//	Logger.Info("File parts to upload: ", fileNames)
//
//	for i, fileName := range fileNames {
//		Logger.Info(i, " ", fileName)
//		Logger.Info(urlInfo.UploadUrls[i].URL)
//
//		apiRequestInfo.url = fmt.Sprintf(urlInfo.UploadUrls[i].URL)
//
//		client := &http.Client{}
//
//		file, err := ioutil.ReadFile(fileName)
//		rb := bytes.NewReader(file)
//
//		if err != nil {
//			Logger.Error("unable to read file")
//			os.Exit(1)
//		}
//
//		req, err := http.NewRequest(http.MethodPut, "https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/42a3ddc4-66df-4a08-a7b3-c8e4c9b7fb77-1?AWSAccessKeyId\\u003dAKIAJNE5FSA2YFQP4BDA\\u0026Expires\\u003d1661265995\\u0026Signature\\u003d5JgAAsR1hN8OIud5AKfYzfM6PQM%3D\"},{\"part\":2,\"url\":\"https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/42a3ddc4-66df-4a08-a7b3-c8e4c9b7fb77-2?AWSAccessKeyId\\u003dAKIAJNE5FSA2YFQP4BDA\\u0026Expires\\u003d1661265995\\u0026Signature\\u003d%2FsX3Hrvg5HizfLtjTRaZ%2BjsC8zo%3D", rb)
//		if err != nil {
//			panic(err)
//		}
//		req.Header.Set("ss-api-key-header", apiRequestInfo.ssApiKeyHeader)
//		req.Header.Set("ss-request-api-header", apiRequestInfo.ssRequestApiHeader)
//
//		resp, err := client.Do(req)
//		if err != nil {
//			panic(err)
//		}
//
//		// TODO: returns a 200? lies.
//		Logger.Info(resp.StatusCode, " | ", resp.Header)
//
//	}
//}

func UploadLogs(args structs.Args) {

	filePath := "/home/craig/github/cigoldstein/yb-tools/yb_log_uploader"
	fileName := "testfile.txt"

	var Uploader structs.Uploader

	Uploader.RequestInfo.SsApiKeyHeader = args.DropzoneIdFlag
	Uploader.RequestInfo.SsRequestApiHeader = "DROP_ZONE"

	// Step 1: Create the dropzone package
	Uploader.PackageInfo = createDropzonePackage(Uploader)
	Logger.Debug("PackageInfo: ", Uploader.PackageInfo)

	// Now that we have the packageCode, we'll generate our clientSecret and checksum
	Uploader.Secrets.ClientSecret = CreateClientSecret()
	Uploader.Secrets.Checksum = CreateChecksum([]byte(Uploader.PackageInfo.PackageCode), []byte(Uploader.Secrets.ClientSecret))

	// This step submits information about the file to the SendSafely API
	// The actual upload is performed later in the workflow
	Uploader.FileInfo = addFileToPackage(fileName, Uploader)
	Logger.Debug(Uploader.FileInfo)

	Uploader.UploadUrlInfo = getUploadUrls(Uploader)

	fileNames := chunkAndEncryptFiles(filePath, fileName, Uploader)

	Logger.Debug("fileNames: ", fileNames)

	//var fileNames []string
	//fileNames = append(fileNames, "split_files/testfile.txt_0", "split_files/testfile.txt_1")

	uploadFilePartsToPackage(Uploader, fileNames)

	markPackageComplete(Uploader)

	Uploader.FinalizeInfo = finalizePackage(Uploader)

	Uploader.HostedDropzoneInfo = submitHostedDropzone(Uploader.PackageInfo.PackageCode)

	Logger.Info("Done")
}
