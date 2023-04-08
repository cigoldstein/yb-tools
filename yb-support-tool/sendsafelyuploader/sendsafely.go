package sendsafelyuploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Uploader struct {
	// TODO PULL args into their own inputs to the uploader via setter / getter
	Args             Args
	URL              string
	APIKey           string
	RequestApiTarget string
	ClientSecret     string
	Checksum         string

	PackageInfo        PackageInfo
	FileInfo           FileInfo
	UploadUrlInfo      UploadUrlInfo
	FinalizeInfo       FinalizeInfo
	HostedDropzoneInfo HostedDropzoneInfo
}

func CreateUploader(SSUrl, SSAPIKey, SSRequestTarget string) *Uploader {
	u := Uploader{URL: SSUrl,
		APIKey:           SSAPIKey,
		RequestApiTarget: SSRequestTarget,
		ClientSecret:     createClientSecret(),
	}
	return &u
}

// uses uploader credentials to send the provided body to the provided enpoint
// returns the response body as an array of bytes and any errors
func (u *Uploader) sendRequest(method, endpoint string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, u.URL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", u.APIKey)
	req.Header.Set("ss-request-api", u.RequestApiTarget)

	// TODO - does this need to be set only sometimes?
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)

}

func (u *Uploader) createDropzonePackage() error {

	endpoint := "drop-zone/v2.0/package/"

	body, err := u.sendRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &u.PackageInfo); err != nil {
		return err
	}
	return nil

}

func (u *Uploader) addFileToPackage(fileName string) error {

	endpoint := fmt.Sprintf("drop-zone/v2.0/package/%s/file", u.PackageInfo.PackageCode)

	//todo pull out to main body
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
		return (err)
	}

	body, err := u.sendRequest(http.MethodPut, endpoint, rbJson)
	if err != nil {
		return (err)
	}

	err = json.Unmarshal(body, &u.FileInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *Uploader) getUploadUrls() error {

	endpoint := fmt.Sprintf("/drop-zone/v2.0/package/%s/file/%s/upload-urls", u.PackageInfo.PackageCode, u.FileInfo.FileID)

	rb := RequestBody{
		FileName:   u.Args.FilesFlag,
		UploadType: "DROP_ZONE",
		Part:       u.FileInfo.Parts,
		FileSize:   u.FileInfo.FileSize,
	}

	rbJson, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	body, err := u.sendRequest(http.MethodPost, endpoint, rbJson)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &u.UploadUrlInfo)
	if err != nil {
		return err
	}
	return nil

}

func uploadFilePartsToPackage(fileNames []string, uploader *Uploader) {
	log.Print("Uploading ", uploader.FileInfo.Parts, " file parts")
	client := &http.Client{}

	for _, uploadUrl := range uploader.UploadUrlInfo.UploadUrls {

		log.Print("fileNames[uploadUrl.Part-1]", fileNames[uploadUrl.Part-1])

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

		var bodyJson UploadUrlInfo
		err = json.Unmarshal(body, &bodyJson)
	}
}

func markPackageComplete(uploader *Uploader) {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-complete", uploader.PackageInfo.PackageCode, uploader.FileInfo.FileID)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	req.Header.Set("ss-api-key", uploader.APIKey)
	req.Header.Set("ss-request-api", uploader.RequestApiTarget)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson FileInfo
	err = json.Unmarshal(body, &bodyJson)
	log.Print(string(body))

}

func finalizePackage(uploader *Uploader) {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/finalize", uploader.PackageInfo.PackageCode)

	type Rb struct {
		Checksum string
	}

	var rb Rb

	rb.Checksum = uploader.Checksum

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

	req.Header.Set("ss-api-key", uploader.APIKey)
	req.Header.Set("ss-request-api", uploader.RequestApiTarget)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson FinalizeInfo
	err = json.Unmarshal(body, &bodyJson)

	uploader.FinalizeInfo = bodyJson

	log.Print("Please use the following secure link to access your file(s): ", uploader.FinalizeInfo.Message+"#keyCode="+uploader.ClientSecret)

}

func submitHostedDropzone(uploader *Uploader) {
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

	var bodyJson HostedDropzoneInfo
	err = json.Unmarshal(body, &bodyJson)
	log.Print("submitDZBody: ", bodyJson)

	uploader.HostedDropzoneInfo = bodyJson
}
