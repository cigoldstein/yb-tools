package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/structs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func createDropzonePackage(Uploader *structs.Uploader) {

	Uploader.RequestInfo.Url = "https://secure-upload.yugabyte.com/drop-zone/v2.0/package/"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, Uploader.RequestInfo.Url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", Uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", Uploader.RequestInfo.SsRequestApiHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.PackageInfo
	err = json.Unmarshal(body, &bodyJson)

	// this returns to Uploader.PackageInfo
	Uploader.PackageInfo = bodyJson
}

func addFileToPackage(fileName string, uploader *structs.Uploader) {

	uploader.RequestInfo.Url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file", uploader.PackageInfo.PackageCode)

	client := &http.Client{}

	type reqBody struct {
		Filename   string `json:"filename"`
		UploadType string `json:"uploadType"`
		Parts      int    `json:"parts"`
		Filesize   int    `json:"filesize"`
	}

	rb := reqBody{
		Filename:   filepath.Base(fileName),
		UploadType: "DROP_ZONE",

		// TODO: need to make these dynamic based on the file sizes
		Parts:    4,
		Filesize: 3550775,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, uploader.RequestInfo.Url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.FileInfo
	err = json.Unmarshal(body, &bodyJson)

	// returns to Uploader.FileInfo
	uploader.FileInfo = bodyJson
}

func getUploadUrls(uploader *structs.Uploader) {

	uploader.RequestInfo.Url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-urls", uploader.PackageInfo.PackageCode, uploader.FileInfo.FileID)

	client := &http.Client{}

	rb := structs.RequestBody{
		FileName:   uploader.Args.FilesFlag,
		UploadType: "DROP_ZONE",
		Part:       uploader.FileInfo.Parts,
		FileSize:   uploader.FileInfo.FileSize,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, uploader.RequestInfo.Url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.UploadUrlInfo
	err = json.Unmarshal(body, &bodyJson)

	uploader.UploadUrlInfo = bodyJson

}

func uploadFilePartsToPackage(fileNames []string, uploader *structs.Uploader) {
	Logger.Info("Uploading ", uploader.FileInfo.Parts, " file parts")
	client := &http.Client{}

	for _, uploadUrl := range uploader.UploadUrlInfo.UploadUrls {

		Logger.Info("fileNames[uploadUrl.Part-1]", fileNames[uploadUrl.Part-1])

		filePart, err := os.ReadFile(fileNames[uploadUrl.Part-1])

		req, err := http.NewRequest(http.MethodPut, uploadUrl.URL, bytes.NewBuffer(filePart))
		if err != nil {
			panic(err)
		}

		req.Header.Set("ss-request-api", "DROP_ZONE")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		var bodyJson structs.UploadUrlInfo
		err = json.Unmarshal(body, &bodyJson)
	}
}

func markPackageComplete(uploader *structs.Uploader) {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-complete", uploader.PackageInfo.PackageCode, uploader.FileInfo.FileID)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	req.Header.Set("ss-api-key", uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.FileInfo
	err = json.Unmarshal(body, &bodyJson)
	Logger.Debug(string(body))

}

func finalizePackage(uploader *structs.Uploader) {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/finalize", uploader.PackageInfo.PackageCode)

	type Rb struct {
		Checksum string
	}

	var rb Rb

	rb.Checksum = uploader.Secrets.Checksum

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	req.Header.Set("ss-api-key", uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.FinalizeInfo
	err = json.Unmarshal(body, &bodyJson)

	uploader.FinalizeInfo = bodyJson

	Logger.Info("Please use the following secure link to access your file(s): ", uploader.FinalizeInfo.Message+"#keyCode="+uploader.Secrets.ClientSecret)

}

func submitHostedDropzone(uploader *structs.Uploader) {
	dropzoneData := url.Values{}
	dropzoneData.Set("name", "4095")
	dropzoneData.Set("email", "cgoldstein@yugabyte.com")
	dropzoneData.Set("packageCode", uploader.PackageInfo.PackageCode)
	dropzoneData.Set("publicApiKey", "BdFZz_JoZqtqPVueANkspD86KZ_PJsW1kIf_jVHeCO0")

	encodedDzData := dropzoneData.Encode()

	client := &http.Client{}

	dropzoneUrl := "https://secure-upload.yugabyte.com/auth/json/?action=submitHostedDropzone"

	req, err := http.NewRequest(http.MethodPost, dropzoneUrl, strings.NewReader(encodedDzData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.HostedDropzoneInfo
	err = json.Unmarshal(body, &bodyJson)
	Logger.Info("submitDZBody: ", bodyJson)

	uploader.HostedDropzoneInfo = bodyJson
}
