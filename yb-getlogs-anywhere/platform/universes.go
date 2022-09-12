package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"yb-get/structs"
)

func GetUniverses(ywInfo structs.YwInfo) structs.YwInfo {

	var api structs.Api
	api.Host = ywInfo.YwHost
	api.Endpoint = fmt.Sprintf("api/v1/customers/%s/universes", ywInfo.YwCustomerUUID)
	api.AuthToken = ywInfo.YwAuthToken
	api.Method = "GET"
	api.ReqBodyJson = ""

	// universes will need to be fetched per customerUUID
	for _, customer := range ywInfo.Customers {
		Logger.Infof("Looking up universes for customerUUID: \"%s\"", customer.Uuid)
		resp := getRequestBody(api, false)

		// check http status code and report the error if it fails
		err := checkRespStatusCode(resp)
		if err != nil {
			Logger.Fatal("API call failed: ", err)
			os.Exit(1)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			Logger.Fatal("Error reading response body: ", err)
		}

		err = json.Unmarshal(body, &ywInfo.Universes)
		if err != nil {
			Logger.Errorf("Error unmarshalling response body: %s: %e", string(body), err)
			os.Exit(1)
		}
	}

	Logger.Debug("Universe response: ", ywInfo.Universes)
	switch {
	case len(ywInfo.Universes) > 1:
		// prompt if there's more than one universe provisioned
		ywInfo = universePrompt(ywInfo)
		return ywInfo
	case len(ywInfo.Universes) == 1:
		ywInfo.YwBundleUUID = ywInfo.Universes[0].UniverseUUID
		return ywInfo
	default:
		Logger.Infof("No universes found for customerUUID: \"%s\"", ywInfo.YwCustomerUUID)
		os.Exit(0)
	}

	return ywInfo
}

func universePrompt(apiInfo structs.YwInfo) structs.YwInfo {

	var universesPromptChoices []structs.PromptChoices

	//generate list for prompt
	for _, v := range apiInfo.Universes {

		var universe structs.PromptChoices

		universe.UUID = v.UniverseUUID
		universe.Name = v.Name

		universesPromptChoices = append(universesPromptChoices, universe)

	}

	promptResponse := Prompt("universeUUID", universesPromptChoices)
	apiInfo.YwUniverseUUID = promptResponse.UUID

	return apiInfo

}
