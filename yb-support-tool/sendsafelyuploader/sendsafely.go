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
	Client           *http.Client

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
		Client:           &http.Client{},
	}
	return &u
}

// uses uploader credentials to send the provided body to the provided enpoint
// returns the response body as an array of bytes and any errors
func (u *Uploader) sendRequest(method, endpoint string, body []byte) ([]byte, error) {
	fmt.Println(u.URL + endpoint)
	req, err := http.NewRequest(method, u.URL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("ss-api-key", u.APIKey)
	req.Header.Set("ss-request-api", u.RequestApiTarget)

	// TODO - does this need to be set only sometimes?
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Invalid HTTP Status Code: %d; Unable to read response body", resp.StatusCode)
		}
		return nil, fmt.Errorf("Invalid HTTP status response; HTTP Error code %d; request body %s", resp.StatusCode, string(respBody))
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

func (u *Uploader) createDropzonePackage() error {

	endpoint := "/drop-zone/v2.0/package/"

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

	endpoint := fmt.Sprintf("/drop-zone/v2.0/package/%s/file", u.PackageInfo.PackageCode)

	rb := FileUpload{
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

func (u *Uploader) getUploadURL(fileName string) error {

	endpoint := fmt.Sprintf("/drop-zone/v2.0/package/%s/file/%s/upload-urls", u.PackageInfo.PackageCode, u.FileInfo.FileID)

	rb := FileUpload{
		Filename:   fileName,
		UploadType: "DROP_ZONE",
		Parts:      u.FileInfo.Parts,
		Filesize:   u.FileInfo.FileSize,
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

func (u *Uploader) uploadFilePartsToPackage(fileNames []string) error {
	log.Print("Uploading ", u.FileInfo.Parts, " file parts")
	client := &http.Client{}

	for _, uploadUrl := range u.UploadUrlInfo.UploadUrls {

		log.Print("fileNames[uploadUrl.Part-1]", fileNames[uploadUrl.Part-1])

		filePart, err := os.ReadFile(fileNames[uploadUrl.Part-1])
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPut, uploadUrl.URL, bytes.NewBuffer(filePart))
		if err != nil {
			return err
		}

		req.Header.Set("ss-request-api", "DROP_ZONE")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var bodyJson UploadUrlInfo
		err = json.Unmarshal(body, &bodyJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Uploader) markPackageComplete() error {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/file/%s/upload-complete", u.PackageInfo.PackageCode, u.FileInfo.FileID)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("ss-api-key", u.APIKey)
	req.Header.Set("ss-request-api", u.RequestApiTarget)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var bodyJson FileInfo
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return err
	}
	return nil
}

func (u *Uploader) finalizePackage() error {
	client := &http.Client{}

	url := fmt.Sprintf("https://secure-upload.yugabyte.com/drop-zone/v2.0/package/%s/finalize", u.PackageInfo.PackageCode)

	type Rb struct {
		Checksum string
	}

	var rb Rb

	rb.Checksum = u.Checksum

	rbJson, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rbJson))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	req.Header.Set("ss-api-key", u.APIKey)
	req.Header.Set("ss-request-api", u.RequestApiTarget)
	req.Header.Set("content-type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var bodyJson FinalizeInfo
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return err
	}

	u.FinalizeInfo = bodyJson

	log.Print("Please use the following secure link to access your file(s): ", u.FinalizeInfo.Message+"#keyCode="+u.ClientSecret)
	return nil
}

func (u *Uploader) submitHostedDropzone() error {
	dropzoneData := url.Values{}
	dropzoneData.Set("name", "4095")
	dropzoneData.Set("email", "cgoldstein@yugabyte.com")
	dropzoneData.Set("packageCode", u.PackageInfo.PackageCode)
	dropzoneData.Set("publicApiKey", "BdFZz_JoZqtqPVueANkspD86KZ_PJsW1kIf_jVHeCO0")

	encodedDzData := dropzoneData.Encode()

	client := &http.Client{}

	dropzoneUrl := "https://secure-upload.yugabyte.com/auth/json/?action=submitHostedDropzone"

	req, err := http.NewRequest(http.MethodPost, dropzoneUrl, strings.NewReader(encodedDzData))
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var bodyJson HostedDropzoneInfo
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return err
	}

	log.Print("submitDZBody: ", bodyJson)

	u.HostedDropzoneInfo = bodyJson
	return nil
}
