package platform

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"yb-get/structs"
)

func getRequestBody(api structs.Api, downloadBundle bool) *http.Response {

	Logger.Debug("API struct: ", api)
	url := fmt.Sprintf("https://%s/%s", api.Host, api.Endpoint)
	req, _ := http.NewRequest(api.Method, url, bytes.NewBuffer([]byte(api.ReqBodyJson)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-AUTH-YW-API-TOKEN", api.AuthToken)

	// self-signed cert
	// TODO: client.Transport instead of changing the DefaultTransport?
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}

	Logger.Debug("API Request: ", url)
	resp, err := client.Do(req)
	if err != nil {
		Logger.Error("err: ", err)
	}

	return resp
}

func checkRespStatusCode(resp *http.Response) error {

	var apiErrorResp structs.ApiErrorResp

	switch {
	case resp.StatusCode != 200:

		// read in body from response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s", resp.Status, err))
		}

		// unmarshal the error message JSON from the API response
		err = json.Unmarshal(body, &apiErrorResp)
		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s", resp.Status, err))
		}

		return errors.New(fmt.Sprintf("%s: %s", resp.Status, apiErrorResp.Error))

	default:
		return nil
	}
}
