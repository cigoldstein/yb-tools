package supportbundle

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

func GetCustomers(ywInfo *YwInfo) *http.Response {
	apiEndpoint := "/api/v1/customers"

	url := fmt.Sprintf("https://%s/%s", ywInfo.YwHost, apiEndpoint)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-AUTH-YW-API-TOKEN", ywInfo.YwAuthToken)

	// self-signed cert
	// TODO: client.Transport instead of changing the DefaultTransport? cmd flag --insecure or -k?
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("err: ", err)
	}

	return resp
}
