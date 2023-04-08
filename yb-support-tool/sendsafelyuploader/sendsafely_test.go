package sendsafelyuploader

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// testing variables
var (
	testKey    = "testKey"
	testTarget = "DROP_ZONE"
)

func TestSendReqeust(t *testing.T) {

	testPath := "/testPath"
	testMethod := http.MethodPut
	testBody := []byte("{type: test}")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(`{"value":"fixed"}`))

	}))
	defer server.Close()

	uploader := CreateUploader(server.URL, testKey, testTarget)

	uploader.sendRequest(testMethod, testPath, testBody)
}
