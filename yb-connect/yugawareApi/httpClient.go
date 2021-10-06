package yugawareApi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"github.com/yugabyte/yb-tools/pkg/util"
	"github.com/yugabyte/yb-tools/yb-connect/prompts"
	"github.com/yugabyte/yb-tools/yb-connect/structs"
	"golang.org/x/net/publicsuffix"
)

var (
	apiBaseUrl       string
	disableCertCheck bool = true
	httpClient       *http.Client
	httpTimeout      time.Duration
	yugawareAuth     structs.YugawareAuth
)

func ConfigHttpClient(yugawareHostname string) (string, string) {

	yugawareUsername := prompts.UsernamePrompt()
	yugawarePassword, err := util.PasswordPrompt()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create cookie jar for HTTP client: %v\n", err)
	}

	transport := http.DefaultTransport.(*http.Transport)
	if disableCertCheck {
		transport = http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	httpClient = &http.Client{Jar: jar, Timeout: httpTimeout, Transport: transport}
	apiBaseUrl = "https://" + yugawareHostname + "/api"

	return yugawareUsername, yugawarePassword
}

func YugawareLogin(yugawareUsername string, yugawarePassword string) structs.YugawareAuth {

	fmt.Println("Logging into Yugaware server")

	loginUrl := apiBaseUrl + "/login"
	response, err := httpClient.PostForm(loginUrl, url.Values{"email": {yugawareUsername}, "password": {yugawarePassword}})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Yugaware login failed: %v\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse login response body: %v\n", err)
	}

	if response.StatusCode == 401 {
		fmt.Fprintf(os.Stderr, "Yugaware login failed: %s\n", body)
		os.Exit(1)
	}

	yugawareAuth = structs.YugawareAuthJsonToStruct(body)

	return yugawareAuth
}

func GetUniverseList(yugawareAuth structs.YugawareAuth) []structs.Universe {

	universeUrl := apiBaseUrl + "/customers/" + yugawareAuth.CustomerUUID + "/universes"

	//fmt.Println("Retrieving Universe list")
	//fmt.Sprintf("Retrieving Universe list from %v", universeUrl)

	response, err := httpClient.Get(universeUrl)
	if err != nil {
		fmt.Printf("Failed to retrieve Universe list: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read Universe list: %v", err)
	}

	var universeList []structs.Universe

	err = json.Unmarshal(body, &universeList)
	if err != nil {
		fmt.Printf("Failed to parse Universe list: %v", err)
	}

	return universeList
}
