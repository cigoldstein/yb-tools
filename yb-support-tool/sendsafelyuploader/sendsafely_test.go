package sendsafelyuploader

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// testing variables
var (
	testKey    = "testKey"
	testTarget = "DROP_ZONE"

	testOKResponse = `{"value":"OK"}`

	testInvalidResponse = `invalid content`
)

func getUploader(URL string) *Uploader {
	return CreateUploader(URL, testKey, testTarget)
}

func TestSendReqeust(t *testing.T) {

	testPath := "/testPath"
	testMethod := http.MethodPut
	testBody := []byte("{type: test}")

	// test OK response
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != testPath {
			t.Errorf("Expected to request '%s', got: %s", testPath, r.URL.Path)
		}
		if r.Header.Get("ss-api-key") != testKey {
			t.Errorf("Expected ss-api-key header: %s, got: %s", testKey, r.Header.Get("ss-api-key"))

		}
		if r.Header.Get("ss-request-api") != testTarget {
			t.Errorf("Expected ss-request-api header: %s, got: %s", testTarget, r.Header.Get("ss-api-key"))
		}

		sentBody, _ := io.ReadAll(r.Body)

		if string(sentBody) != string(testBody) {
			t.Errorf("Expected sent message: %s, got: %s", string(testBody), string(sentBody))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testOKResponse))

	}))
	defer serverOK.Close()

	uploader := CreateUploader(serverOK.URL, testKey, testTarget)

	resp, err := uploader.sendRequest(testMethod, testPath, testBody)

	if err != nil {
		t.Errorf("Expected no error, instead got %s", err)
	}

	if string(resp) != testOKResponse {
		t.Errorf("Expected to get an OK response, instead got %s", string(resp))
	}

	// test Invalid response

	serverInvalid := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(testInvalidResponse))

	}))
	defer serverInvalid.Close()

	uploader = CreateUploader(serverInvalid.URL, testKey, testTarget)

	resp, err = uploader.sendRequest(testMethod, testPath, testBody)

	if err == nil {
		t.Errorf("Expected an error due to invalid HTTP response, instead got %s", err)
	}

	if string(resp) != "" {
		t.Errorf("Expected response '%s', instead got '%s'", testInvalidResponse, string(resp))
	}

}

func TestAddFileToPackage(t *testing.T) {

	fileName := "mytestfile"

	// test something that can't be unmarshalled
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testInvalidResponse))

	}))
	defer server.Close()

	u := getUploader(server.URL)

	err := u.addFileToPackage(fileName)

	if !strings.Contains(err.Error(), testInvalidResponse) {
		t.Errorf("Expected error response to show what server response failed to unmarshal, instead got: '%s'", err)
	}

}

func TestGetUploadURL(t *testing.T) {
	fileName := "mytestfile"

	validUploadURLInfo, _ := json.Marshal(
		UploadUrlInfo{
			URLS: []UploadURL{
				UploadURL{
					Part: 1,
					URL:  `https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/f0805497-314d-4eac-ac52-80f4fb40d9eb-1?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1680894043&Signature=3R%2FbQJrY1XXObIa5XUvm6ntk3sE%3D`,
				},
			},

			Response: "SUCCESS"})

	// check that we hit endpoint "/drop-zone/v2.0/package/%s/file/%s/upload-urls"
	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// test that r.Body contains upload-urls
		w.Write(validUploadURLInfo)

	}))
	defer validServer.Close()
	u := getUploader(validServer.URL)

	// test that if PackageCode is missing in uploader, getUploadURL fails with error
	u.PackageInfo.PackageCode = ""
	err := u.getUploadURL(fileName)

	if err == nil {
		t.Error("Expected an error from function because PackageCode empty, but instead function has no err")
	}
	// error string should contain filename for reference
	if !strings.Contains(err.Error(), fileName) {
		t.Errorf("Expected failure message to contain filename: '%s', instead got message: '%s'", fileName, err.Error())
	}

	// test that if we get `"RESPONSE": "FAIL"` return, that we get a valid error

	invalidUploadURLInfo, _ := json.Marshal(
		UploadUrlInfo{
			Response: "FAIL",
			Message:  "An error occurred.  Error Id {some error code}",
		})

	failServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(invalidUploadURLInfo))

	}))
	defer failServer.Close()

	uf := getUploader(failServer.URL)

	uf.PackageInfo.PackageCode = "sMN0ghoCOkxFiOBHlXhscb1qS3fMmd7sV0012gIRutU"
	err = uf.getUploadURL(fileName)

	if err == nil {
		t.Error("Expected to get response message from server that we got a FAILURE, instead got not error")
	}
}
