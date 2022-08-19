package uploader

// TODO: create a "createHttpClient function so that there's less reused code on api calls
// TODO: break structs out into separate file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/log"
	"net/http"
)

var logger = log.Log()
var apiRequestInfo requestInfo
var apiFileInfo fileInfo
var apiPackageInfo packageInfo
var apiUploadUrlInfo uploadUrlInfo

type requestInfo struct {
	url                string
	ssApiKeyHeader     string
	ssRequestApiHeader string
}

type packageInfo struct {
	PackageID    string `json:"packageId"`
	PackageCode  string `json:"packageCode"`
	ServerSecret string `json:"serverSecret"`
	Response     string `json:"response"`
}

type fileInfo struct {
	FileID          string `json:"fileId"`
	FileName        string `json:"fileName"`
	FileSize        string `json:"fileSize"`
	Parts           int    `json:"parts"`
	FileUploaded    string `json:"fileUploaded"`
	FileUploadedStr string `json:"fileUploadedStr"`
	FileVersion     string `json:"fileVersion"`
	CreatedByEmail  string `json:"createdByEmail"`
	Response        string `json:"response"`
	Message         string `json:"message"`
}

type uploadUrlInfo struct {
	UploadUrls []struct {
		Part int    `json:"part"`
		URL  string `json:"url"`
	} `json:"uploadUrls"`
	Response string `json:"response"`
}

func createHttpPut() {

}

func createHttpPost() {

}

func createDropzonePackage() packageInfo {

	apiRequestInfo.url = "https://secure-upload.yugabyte.com/drop-zone/v2.0/package/"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, apiRequestInfo.url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", apiRequestInfo.ssApiKeyHeader)
	req.Header.Set("ss-request-api", apiRequestInfo.ssRequestApiHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson packageInfo
	err = json.Unmarshal(body, &bodyJson)

	return bodyJson
}

func addFileToPackage() fileInfo {

	apiRequestInfo.url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file", apiPackageInfo.PackageCode)

	client := &http.Client{}

	type reqBody struct {
		Filename   string `json:"filename"`
		UploadType string `json:"uploadType"`
		Parts      int    `json:"parts"`
		Filesize   int    `json:"filesize"`
	}

	rb := reqBody{
		Filename:   "testfile.txt",
		UploadType: "DROP_ZONE",
		Parts:      2,
		Filesize:   18,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, apiRequestInfo.url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", apiRequestInfo.ssApiKeyHeader)
	req.Header.Set("ss-request-api", apiRequestInfo.ssRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson fileInfo
	err = json.Unmarshal(body, &bodyJson)
	logger.Info(string(body))

	return bodyJson
}

func getUploadUrls() uploadUrlInfo {

	apiRequestInfo.url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-urls", apiPackageInfo.PackageCode, apiFileInfo.FileID)

	client := &http.Client{}

	type reqBody struct {
		Part int `json:"part"`
	}

	rb := reqBody{
		Part: 1,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, apiRequestInfo.url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", apiRequestInfo.ssApiKeyHeader)
	req.Header.Set("ss-request-api", apiRequestInfo.ssRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson uploadUrlInfo
	err = json.Unmarshal(body, &bodyJson)
	logger.Info(string(body))

	return bodyJson

}

func encryptFileParts() {

}

func markPackageComplete() {

}

func UploadLogs(caseNum int, email string, dropzoneId string, isDropzoneFlagChanged bool, files []string) {

	logger.Info(apiRequestInfo.url)

	apiRequestInfo.ssApiKeyHeader = dropzoneId
	apiRequestInfo.ssRequestApiHeader = "DROP_ZONE"

	// set the default here for now but maybe hardcode the dropzone id in the struct instead of cobra
	//apiRequestInfo.ssApiKeyHeader := dropzoneId

	//packageInfo := createDropzonePackage()

	// hard code these values for testing so that we don't create a million empty packages
	// packageInfo would normally be created by the "createDropzonePackage" function above
	// packageInfo is a struct built from the HTTP json response of the api call
	// the values below are actual values from a "createDropzonePackage" call
	apiPackageInfo.PackageID = "X9V2-L8S1"
	apiPackageInfo.PackageCode = "TWC3p9Sq8kyHDNceOl3pk59SvGqqVNiP1fNTFHNvZ00"
	apiPackageInfo.ServerSecret = "ABSEtmVjrkwKsuH7qrbRJr61O5BX76dumg"
	apiPackageInfo.Response = "SUCCESS"

	// fileInfo := addFileToPackage()

	// dump struct
	// logger.Infof("%+v\n", fileInfo)

	// hard code fileInfo for now, similar to packageInfo
	apiFileInfo.FileID = "72448c21-54b7-41b9-93e1-8a6406933f8a"
	apiFileInfo.FileName = "testfile.txt"
	apiFileInfo.FileSize = "18"
	apiFileInfo.Parts = 2
	apiFileInfo.FileUploaded = "Aug 19, 2022 1:03:38 AM"
	apiFileInfo.FileUploadedStr = "Thu Aug 18 at 21:03 (EDT)"
	apiFileInfo.FileVersion = "1"
	apiFileInfo.FileVersion = "1"
	apiFileInfo.CreatedByEmail = "Anonymous Recipient"
	apiFileInfo.Response = "SUCCESS"
	apiFileInfo.Message = "c164b143-db33-436d-ab1c-75809a640dc4"

	urlInfo := getUploadUrls()
	logger.Infof("%+v\n", urlInfo)

}
