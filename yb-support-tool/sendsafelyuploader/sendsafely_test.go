package sendsafelyuploader

import (
	"crypto/rand"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestSendRequest(t *testing.T) {

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

	validUploadURLObject := UploadUrlInfo{
		URLS: []UploadURL{
			UploadURL{
				Part: 1,
				URL:  `https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/f0805497-314d-4eac-ac52-80f4fb40d9eb-1?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1680894043&Signature=3R%2FbQJrY1XXObIa5XUvm6ntk3sE%3D`,
			},
		},

		Response: "SUCCESS"}
	validUploadURLJSON, _ := json.Marshal(validUploadURLObject)

	// check that we hit endpoint "/drop-zone/v2.0/package/%s/file/%s/upload-urls"
	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.URL.Path, "/upload-urls") {
			t.Errorf("Expected to hit path containing /upload-urls, instead got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		// test that r.Body contains upload-urls
		w.Write(validUploadURLJSON)

	}))
	defer validServer.Close()
	u := getUploader(validServer.URL)

	// no error if we have valid package code and fileid
	u.UploadUrlInfo = UploadUrlInfo{}
	u.PackageInfo.PackageCode = "sMN0ghoCOkxFiOBHlXhscb1qS3fMmd7sV0012gIRutU"
	u.FileInfo.FileID = "8c652651-7e0e-4eb8-bb1b-24710bd4ee35"

	err := u.getUploadURL(fileName)
	if err != nil {
		t.Errorf("Expected no error with valid package code and fileid, but recieved %s", err)
	}
	// and we set UploadUrlInfo properly based on response from server
	if cmp.Equal(u.UploadUrlInfo, UploadUrlInfo{}) {
		t.Errorf("Expected UploadUrlInfo to be populated, but it is empty")
	}
	if !cmp.Equal(u.UploadUrlInfo, validUploadURLObject) {
		t.Errorf("Expected to populate UploadUrlInfo with a valid respose, but did not")
	}

	// test that if PackageCode is missing in uploader, getUploadURL fails with error
	u.PackageInfo.PackageCode = ""
	err = u.getUploadURL(fileName)

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

func TestUploadFilePart(t *testing.T) {

	var fileChunkSize float64 = 2.5 * (1 << 20)

	// generate byte slice to represent the byte slice being uploaded
	filePart := make([]byte, int(fileChunkSize))

	// non-zero file
	rand.Read(filePart)

	// Example:
	/*
		{
		  "uploadUrls": [
		    {
		      "part": 1,
		      "url": "https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/8c652651-7e0e-4eb8-bb1b-24710bd4ee35-1?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1664069315&Signature=yOAGja0sPBv75MLa2Jf6YfXYom8%3D"
		    },
		    {
		      "part": 2,
		      "url": "https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/8c652651-7e0e-4eb8-bb1b-24710bd4ee35-2?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1664069315&Signature=xBkpJlOM1OUkoKRwvBf1mPrDlM0%3D"
		    },
		    {
		      "part": 3,
		      "url": "https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/8c652651-7e0e-4eb8-bb1b-24710bd4ee35-3?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1664069315&Signature=kxfBBSQfE7oLu83%2FwALAmKboeL8%3D"
		    },
		    {
		      "part": 4,
		      "url": "https://sendsafely-us-west-2.s3-accelerate.amazonaws.com/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/8c652651-7e0e-4eb8-bb1b-24710bd4ee35-4?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1664069315&Signature=mfnwodVkgMvRSOSVCrRie3yxqcE%3D"
		    }
		  ],
		  "response": "SUCCESS"
		}

	*/

	validPath := "/commercial/e93ec274-e586-4f55-8eab-498a8444cf94/f0805497-314d-4eac-ac52-80f4fb40d9eb-1?AWSAccessKeyId=AKIAJNE5FSA2YFQP4BDA&Expires=1680894043&Signature=3R%2FbQJrY1XXObIa5XUvm6ntk3sE%3D"
	sampleURL := "http://example.com" + validPath
	sURL, _ := url.Parse(sampleURL)

	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != sURL.Path {
			t.Errorf("Expected to request '%s', got: %s", sURL.Path, r.URL.Path)
		}
		if r.Header.Get("ss-request-api") != "DROP_ZONE" {
			t.Errorf("Expected ss-request-api header: DROP_ZONE, got: %s", r.Header.Get("ss-api-key"))
		}

		sentBody, _ := io.ReadAll(r.Body)

		if string(sentBody) != string(filePart) {
			t.Errorf("Expected sent and received messages are not the same, but they are")
		}
		//if len(sentBody) != len(filePart) {
		//
		//}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testOKResponse))

	}))

	validURL := validServer.URL + validPath

	defer validServer.Close()
	u := getUploader(validServer.URL)

	// valid file part and valid URL returns no error
	err := u.uploadFilePart(filePart, validURL)

	if err != nil {
		t.Errorf("Expected no error but instead got %s", err)
	}

	// if I'm given an invalid URL, returns an error
	invalidURL := "1234"
	err = u.uploadFilePart(filePart, invalidURL)
	if err == nil {
		t.Errorf("Expected error on invalid URL but did not receive one")
	}

	// if part []byte is size of zero, returns an error
	err = u.uploadFilePart(make([]byte, 0), validURL)
	if err == nil {
		t.Errorf("Expected error for empty part but did not receive one")
	}

	/*
		// testing non-ok response
		invalidServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(testOKResponse))
		}))
	*/
}
