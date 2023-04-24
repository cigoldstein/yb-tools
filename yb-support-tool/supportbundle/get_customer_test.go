package supportbundle

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomer(t *testing.T) {

	ywInfo.YwAuthToken = "e5c6eb2f-7e30-4d5e-b0a2-f197b01d9f79"
	apiEndpoint := "/api/v1/customers"

	// TODO: change this to test for multiple customers once multi-tenancy is available in YBA
	validResponse :=
		`[
			{
			   "uuid":"6553ea6d-485c-4ae8-861a-736c2c29ec46",
			   "code":"dev",
			   "name":"John Smith",
			   "creationDate":"2023-02-08T18:22:38+0000",
			   "features":{},
			   "customerId":1
			}
		 ])`

	t.Log(validResponse)
	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ywInfo.YwHost = r.URL.Hostname()
		t.Log(ywInfo.YwHost)

		if r.URL.Path != apiEndpoint {
			t.Errorf("Expected to hit endpoint %s, instead got %s", apiEndpoint, r.URL.Path)
		}
		if len(r.Header.Get("X-AUTH-YW-API-TOKEN")) == 0 {
			t.Error("Expected X-AUTH-YW-API-TOKEN header but it was not provided")
		}
		if r.Header.Get("X-AUTH-YW-API-TOKEN") != ywInfo.YwAuthToken {
			t.Errorf("Expected X-AUTH-YW-API-TOKEN header: %s, got %s", ywInfo.YwAuthToken, r.Header.Get("X-AUTH-YW-API-TOKEN header"))
		}
	}))
	defer validServer.Close()

	resp := GetCustomers(&ywInfo)

	t.Log(resp)
}
