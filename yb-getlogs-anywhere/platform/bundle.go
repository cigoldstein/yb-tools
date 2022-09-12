package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"

	"yb-get/structs"
)

func CreateBundle() {

}

func GetBundles(ywInfo structs.YwInfo) {

	var api structs.Api
	api.Host = ywInfo.YwHost
	api.Endpoint = fmt.Sprintf("api/v1/customers/%s/universes/%s/support_bundle", ywInfo.YwCustomerUUID, ywInfo.YwUniverseUUID)
	api.AuthToken = ywInfo.YwAuthToken
	api.Method = "GET"
	api.ReqBodyJson = ""

	// look up bundles for the unviverse that was selected by the user
	Logger.Infof("Looking up bundles for universeUUID: \"%s\"", ywInfo.YwUniverseUUID)
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

	err = json.Unmarshal(body, &ywInfo.Bundles)
	if err != nil {
		Logger.Errorf("Error unmarshalling response body: %s: %e", string(body), err)
		os.Exit(1)
	}

	switch {
	case len(ywInfo.Bundles) > 1:
		// prompt if there's more than one universe provisioned
		ywInfo = bundlesPrompt(ywInfo)
	case len(ywInfo.Bundles) == 1:
		ywInfo.YwBundleUUID = ywInfo.Bundles[0].BundleUUID
		ywInfo.YwBundleFilename = ywInfo.Bundles[0].Path
	default:
		Logger.Infof("No bundles found for universeUUID: \"%s\"", ywInfo.YwUniverseUUID)
		os.Exit(0)
	}

	downloadBundle(ywInfo)

}

func downloadBundle(ywInfo structs.YwInfo) {

	var api structs.Api
	api.Host = ywInfo.YwHost
	api.Endpoint = fmt.Sprintf("api/v1/customers/%s/universes/%s/support_bundle/%s/download", ywInfo.YwCustomerUUID, ywInfo.YwUniverseUUID, ywInfo.YwBundleUUID)
	api.AuthToken = ywInfo.YwAuthToken
	api.Method = "GET"
	api.ReqBodyJson = ""

	bundleFile, err := os.OpenFile(filepath.Base(ywInfo.YwBundleFilename), os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		Logger.Fatalf("Unable to open file for writing: file: %s, err: %e", ywInfo.YwBundleFilename, err)
	}
	defer func(bundleFile *os.File) {
		err = bundleFile.Close()
		if err != nil {
			Logger.Error("Unable to close file", err)
		}
	}(bundleFile)

	resp := getRequestBody(api, true)

	err = checkRespStatusCode(resp)
	if err != nil {
		Logger.Error("API call failed: ", err)
	}

	// write out to bundleFile (file on disk) and bar (progress bar) at the same time
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Downloading %s: ", filepath.Base(ywInfo.YwBundleFilename)),
	)
	_, err = io.Copy(io.MultiWriter(bundleFile, bar), resp.Body)
	if err != nil {
		Logger.Error("Unable to write file: ", err)
	}
	fmt.Println()

	err = resp.Body.Close()
	if err != nil {
		Logger.Error("Unable to close response body: ", err)
	}

	Logger.Infof("Finished downloading %s", ywInfo.YwBundleFilename)

}

func bundlesPrompt(apiInfo structs.YwInfo) structs.YwInfo {

	var bundlesPromptChoices []structs.PromptChoices

	//generate list for prompt
	for _, v := range apiInfo.Bundles {

		var bundle structs.PromptChoices

		bundle.UUID = v.BundleUUID
		bundle.Name = filepath.Base(v.Path)

		bundlesPromptChoices = append(bundlesPromptChoices, bundle)
	}

	promptResponse := Prompt("bundleUUID", bundlesPromptChoices)
	apiInfo.YwBundleUUID = promptResponse.UUID
	apiInfo.YwBundleFilename = promptResponse.Name

	return apiInfo
}
