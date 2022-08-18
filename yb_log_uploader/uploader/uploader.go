package uploader

import (
	"fmt"
	"io/ioutil"
	"main/log"
	"net/http"
)

var logger = log.Log()

func createDropzonePackage(dropzoneId string) {

	url := "https://secure-upload.yugabyte.com/drop-zone/v2.0/package/"
	ssApiKeyHeader := dropzoneId
	ssRequestApiHeader := "DROP_ZONE"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("ss-api-key", ssApiKeyHeader)
	req.Header.Set("ss-request-api", ssRequestApiHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	logger.Info(resp.StatusCode)
	fmt.Println(string(body))

}

func UploadLogs(caseNum int, email string, dropzoneId string, isDropzoneFlagChanged bool) {
	createDropzonePackage(dropzoneId)
}
