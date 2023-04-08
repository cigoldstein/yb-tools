package sendsafelyuploader

type RequestInfo struct {
	Url                string
	SsApiKey           string
	SsRequestApiTarget string
}

type Args struct {
	FilesFlag             []string
	CaseNumFlag           int
	EmailFlag             string
	DropzoneIdFlag        string
	IsDropzoneFlagChanged bool
}

type FileUpload struct {
	Filename   string `json:"filename"`
	UploadType string `json:"uploadType"`
	Parts      int    `json:"parts"`
	Filesize   int64  `json:"filesize,omitempty,string"`
}

type PackageInfo struct {
	PackageID    string `json:"packageId"`
	PackageCode  string `json:"packageCode"`
	ServerSecret string `json:"serverSecret"`
	Response     string `json:"response"`
}

type FileInfo struct {
	FileID          string `json:"fileId"`
	FileName        string `json:"fileName"`
	FileSize        int64  `json:"fileSize,omitempty,string"`
	Parts           int    `json:"parts"`
	FileUploaded    string `json:"fileUploaded"`
	FileUploadedStr string `json:"fileUploadedStr"`
	FileVersion     string `json:"fileVersion"`
	CreatedByEmail  string `json:"createdByEmail"`
	Response        string `json:"response"`
	Message         string `json:"message"`
}

type UploadUrlInfo struct {
	UploadUrls []struct {
		Part int    `json:"part"`
		URL  string `json:"url"`
	} `json:"uploadUrls"`
	Response string `json:"response"`
}

type FinalizeInfo struct {
	NeedsLink bool   `json:"needsLink"`
	Response  string `json:"response"`
	Message   string `json:"message"`
}

type HostedDropzoneInfo struct {
	Success         string   `json:"success"`
	Data            string   `json:"data"`
	Digest          string   `json:"digest"`
	IntegrationUrls []string `json:"integrationUrls"`
}
