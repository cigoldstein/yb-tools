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
	"strings"
)

func createDropzonePackage(Uploader structs.Uploader) structs.PackageInfo {

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
	return bodyJson
}

func addFileToPackage(fileName string, Uploader structs.Uploader) structs.FileInfo {

	Uploader.RequestInfo.Url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file", Uploader.PackageInfo.PackageCode)

	client := &http.Client{}

	type reqBody struct {
		Filename   string `json:"filename"`
		UploadType string `json:"uploadType"`
		Parts      int    `json:"parts"`
		Filesize   int    `json:"filesize"`
	}

	rb := reqBody{
		Filename:   fileName,
		UploadType: "DROP_ZONE",
		Parts:      4,
		Filesize:   3550775,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, Uploader.RequestInfo.Url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", Uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", Uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.FileInfo
	err = json.Unmarshal(body, &bodyJson)
	Logger.Info(string(body))

	// returns to Uploader.FileInfo
	return bodyJson
}

func getUploadUrls(Uploader structs.Uploader) structs.UploadUrlInfo {

	Uploader.RequestInfo.Url = fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-urls", Uploader.PackageInfo.PackageCode, Uploader.FileInfo.FileID)

	client := &http.Client{}

	rb := structs.RequestBody{
		FileName:   Uploader.Args.FilesFlag,
		UploadType: "DROP_ZONE",
		Part:       Uploader.FileInfo.Parts,
		FileSize:   Uploader.FileInfo.FileSize,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, Uploader.RequestInfo.Url, bytes.NewBuffer(rbJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", Uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", Uploader.RequestInfo.SsRequestApiHeader)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson structs.UploadUrlInfo
	err = json.Unmarshal(body, &bodyJson)
	Logger.Debug("Body: ", string(body))

	return bodyJson

}

func uploadFilePartsToPackage(uploader structs.Uploader, fileNames []string) {
	Logger.Info("Uploading ", uploader.FileInfo.Parts, " file parts")
	client := &http.Client{}

	for _, uploadUrl := range uploader.UploadUrlInfo.UploadUrls {

		// TODO: change fileParts from list to a map with part:file instead of relying on slice index
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
		Logger.Info("Body: ", string(body))
	}
}

func markPackageComplete(Uploader structs.Uploader) {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-complete", Uploader.PackageInfo.PackageCode, Uploader.FileInfo.FileID)
	Logger.Info("packageCode", Uploader.PackageInfo.PackageCode)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	req.Header.Set("ss-api-key", Uploader.RequestInfo.SsApiKeyHeader)
	req.Header.Set("ss-request-api", Uploader.RequestInfo.SsRequestApiHeader)
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

func finalizePackage(uploader structs.Uploader) structs.FinalizeInfo {
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
	Logger.Info("Body: ", bodyJson)

	return bodyJson

}

func submitHostedDropzone(packageCode string) structs.HostedDropzoneInfo {
	dropzoneData := url.Values{}
	dropzoneData.Set("name", "4095")
	dropzoneData.Set("email", "cgoldstein@yugabyte.com")
	dropzoneData.Set("packageCode", packageCode)
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

	return bodyJson

}
