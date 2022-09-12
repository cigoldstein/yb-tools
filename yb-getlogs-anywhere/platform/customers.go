package platform

import (
	"encoding/json"
	"io"
	"os"

	"go.uber.org/zap"

	"yb-get/log"
	"yb-get/structs"
)

var Logger = log.CreateLogger(false, false)

func GetCustomers(ywInfo structs.YwInfo) structs.YwInfo {

	zap.L().Sugar().Info("test")
	var api structs.Api
	api.Host = ywInfo.YwHost
	api.Endpoint = "api/v1/customers"
	api.AuthToken = ywInfo.YwAuthToken
	api.Method = "GET"
	api.ReqBodyJson = ""

	Logger.Infof("Looking up customers for \"%s\"", api.Host)
	resp := getRequestBody(api, false)

	// check http status code and report the error if it fails
	err := checkRespStatusCode(resp)
	if err != nil {
		Logger.Error("API call failed: ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Fatal("Error reading response body: ", err)
	}

	err = json.Unmarshal(body, &ywInfo.Customers)
	if err != nil {
		Logger.Errorf("Error unmarshalling response body: %s: %e", string(body), err)
		os.Exit(1)
	}

	// if there's more than one customerUUID, prompt the user for which one to use
	switch {
	case len(ywInfo.Customers) > 1:
		customerPrompt(ywInfo)
		return ywInfo
	case len(ywInfo.Customers) == 1:
		// if there's only one, set it in ywInfo
		ywInfo.YwCustomerUUID = ywInfo.Customers[0].Uuid
		return ywInfo
	default:
		Logger.Infof("No customers found for host: \"%s\"", api.Host)
		os.Exit(0)
	}

	return ywInfo
}

func customerPrompt(apiInfo structs.YwInfo) structs.YwInfo {

	var customersPromptChoices []structs.PromptChoices

	//generate list for prompt
	for _, v := range apiInfo.Customers {

		var customer structs.PromptChoices

		customer.UUID = v.Uuid
		customer.Name = v.Name

		customersPromptChoices = append(customersPromptChoices, customer)
	}

	promptResponse := Prompt("customerUUID", customersPromptChoices)
	apiInfo.YwCustomerUUID = promptResponse.UUID

	return apiInfo
}
