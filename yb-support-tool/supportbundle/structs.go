package supportbundle

import (
	"time"
)

type Logger struct {
	Args struct {
		DebugFlag   bool
		VerboseFlag bool
	}
}

type YwInfo struct {
	YwHost           string
	YwAuthToken      string
	YwCustomerUUID   string
	YwUniverseUUID   string
	YwBundleUUID     string
	YwBundleFilename string

	Customers []CustomerResp
	Universes []UniverseResp
	Bundles   []BundleResp

	Logger struct {
		DebugFlag   bool
		VerboseFlag bool
	}
}

type Api struct {
	Host        string
	Endpoint    string
	AuthToken   string
	Method      string
	ReqBodyJson string
}

type PromptChoices struct {
	UUID string
	Name string
}

type CreateReqBodyJson struct {
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Components []string `json:"components"`
}

type CustomerResp struct {
	Uuid         string `json:"uuid"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate"`
	Features     struct {
	} `json:"features"`
	CustomerId int `json:"customerId"`
}

type UniverseResp struct {
	UniverseUUID string `json:"universeUUID"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate"`
	Version      int    `json:"version"`
	Resources    struct {
		PricePerHour      float64  `json:"pricePerHour"`
		EbsPricePerHour   float64  `json:"ebsPricePerHour"`
		NumCores          float64  `json:"numCores"`
		MemSizeGB         float64  `json:"memSizeGB"`
		VolumeCount       int      `json:"volumeCount"`
		VolumeSizeGB      int      `json:"volumeSizeGB"`
		NumNodes          int      `json:"numNodes"`
		Gp3FreePiops      int      `json:"gp3FreePiops"`
		Gp3FreeThroughput int      `json:"gp3FreeThroughput"`
		AzList            []string `json:"azList"`
	} `json:"resources"`
	UniverseDetails struct {
		NodeExporterUser        string `json:"nodeExporterUser"`
		UniverseUUID            string `json:"universeUUID"`
		ExpectedUniverseVersion int    `json:"expectedUniverseVersion"`
		EncryptionAtRestConfig  struct {
			EncryptionAtRestEnabled bool   `json:"encryptionAtRestEnabled"`
			OpType                  string `json:"opType"`
			Type                    string `json:"type"`
		} `json:"encryptionAtRestConfig"`
		CommunicationPorts struct {
			MasterHttpPort       int `json:"masterHttpPort"`
			MasterRpcPort        int `json:"masterRpcPort"`
			TserverHttpPort      int `json:"tserverHttpPort"`
			TserverRpcPort       int `json:"tserverRpcPort"`
			YbControllerHttpPort int `json:"ybControllerHttpPort"`
			YbControllerrRpcPort int `json:"ybControllerrRpcPort"`
			RedisServerHttpPort  int `json:"redisServerHttpPort"`
			RedisServerRpcPort   int `json:"redisServerRpcPort"`
			YqlServerHttpPort    int `json:"yqlServerHttpPort"`
			YqlServerRpcPort     int `json:"yqlServerRpcPort"`
			YsqlServerHttpPort   int `json:"ysqlServerHttpPort"`
			YsqlServerRpcPort    int `json:"ysqlServerRpcPort"`
			NodeExporterPort     int `json:"nodeExporterPort"`
		} `json:"communicationPorts"`
		ExtraDependencies struct {
			InstallNodeExporter bool `json:"installNodeExporter"`
		} `json:"extraDependencies"`
		FirstTry bool `json:"firstTry"`
		Clusters []struct {
			Uuid        string `json:"uuid"`
			ClusterType string `json:"clusterType"`
			UserIntent  struct {
				UniverseName      string   `json:"universeName"`
				Provider          string   `json:"provider"`
				ProviderType      string   `json:"providerType"`
				ReplicationFactor int      `json:"replicationFactor"`
				RegionList        []string `json:"regionList"`
				InstanceType      string   `json:"instanceType"`
				NumNodes          int      `json:"numNodes"`
				YbSoftwareVersion string   `json:"ybSoftwareVersion"`
				AccessKeyCode     string   `json:"accessKeyCode"`
				DeviceInfo        struct {
					VolumeSize   int    `json:"volumeSize"`
					NumVolumes   int    `json:"numVolumes"`
					StorageClass string `json:"storageClass"`
					StorageType  string `json:"storageType"`
				} `json:"deviceInfo"`
				AssignPublicIP            bool   `json:"assignPublicIP"`
				AssignStaticPublicIP      bool   `json:"assignStaticPublicIP"`
				UseTimeSync               bool   `json:"useTimeSync"`
				EnableYCQL                bool   `json:"enableYCQL"`
				YsqlPassword              string `json:"ysqlPassword"`
				YcqlPassword              string `json:"ycqlPassword"`
				EnableYSQLAuth            bool   `json:"enableYSQLAuth"`
				EnableYCQLAuth            bool   `json:"enableYCQLAuth"`
				EnableYSQL                bool   `json:"enableYSQL"`
				EnableYEDIS               bool   `json:"enableYEDIS"`
				EnableNodeToNodeEncrypt   bool   `json:"enableNodeToNodeEncrypt"`
				EnableClientToNodeEncrypt bool   `json:"enableClientToNodeEncrypt"`
				EnableVolumeEncryption    bool   `json:"enableVolumeEncryption"`
				EnableIPV6                bool   `json:"enableIPV6"`
				EnableExposingService     string `json:"enableExposingService"`
				AwsArnString              string `json:"awsArnString"`
				UseHostname               bool   `json:"useHostname"`
				UseSystemd                bool   `json:"useSystemd"`
				MasterGFlags              struct {
				} `json:"masterGFlags"`
				TserverGFlags struct {
				} `json:"tserverGFlags"`
				InstanceTags struct {
				} `json:"instanceTags"`
				YbcPackagePath string `json:"ybcPackagePath,omitempty"`
			} `json:"userIntent"`
			PlacementInfo struct {
				CloudList []struct {
					Uuid       string `json:"uuid"`
					Code       string `json:"code"`
					RegionList []struct {
						Uuid   string `json:"uuid"`
						Code   string `json:"code"`
						Name   string `json:"name"`
						AzList []struct {
							Uuid              string `json:"uuid"`
							Name              string `json:"name"`
							ReplicationFactor int    `json:"replicationFactor"`
							Subnet            string `json:"subnet"`
							NumNodesInAZ      int    `json:"numNodesInAZ"`
							IsAffinitized     bool   `json:"isAffinitized"`
						} `json:"azList"`
					} `json:"regionList"`
				} `json:"cloudList"`
			} `json:"placementInfo"`
			Index   int `json:"index"`
			Regions []struct {
				Uuid      string  `json:"uuid"`
				Code      string  `json:"code"`
				Name      string  `json:"name"`
				YbImage   string  `json:"ybImage"`
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
				Zones     []struct {
					Uuid   string `json:"uuid"`
					Code   string `json:"code"`
					Name   string `json:"name"`
					Active bool   `json:"active"`
					Subnet string `json:"subnet"`
				} `json:"zones"`
				Active bool `json:"active"`
				Config struct {
				} `json:"config"`
			} `json:"regions"`
		} `json:"clusters"`
		CurrentClusterType       string        `json:"currentClusterType"`
		NodePrefix               string        `json:"nodePrefix"`
		RootCA                   string        `json:"rootCA"`
		RootAndClientRootCASame  bool          `json:"rootAndClientRootCASame"`
		UserAZSelected           bool          `json:"userAZSelected"`
		ResetAZConfig            bool          `json:"resetAZConfig"`
		UpdateInProgress         bool          `json:"updateInProgress"`
		BackupInProgress         bool          `json:"backupInProgress"`
		UpdateSucceeded          bool          `json:"updateSucceeded"`
		UniversePaused           bool          `json:"universePaused"`
		NextClusterIndex         int           `json:"nextClusterIndex"`
		AllowInsecure            bool          `json:"allowInsecure"`
		SetTxnTableWaitCountFlag bool          `json:"setTxnTableWaitCountFlag"`
		ItestS3PackagePath       string        `json:"itestS3PackagePath"`
		RemotePackagePath        string        `json:"remotePackagePath"`
		NodesResizeAvailable     bool          `json:"nodesResizeAvailable"`
		UseNewHelmNamingStyle    bool          `json:"useNewHelmNamingStyle"`
		UseYbcForBackups         bool          `json:"useYbcForBackups"`
		ImportedState            string        `json:"importedState"`
		Capability               string        `json:"capability"`
		TargetXClusterConfigs    []interface{} `json:"targetXClusterConfigs"`
		SourceXClusterConfigs    []interface{} `json:"sourceXClusterConfigs"`
		NodeDetailsSet           []struct {
			NodeIdx   int    `json:"nodeIdx"`
			NodeName  string `json:"nodeName"`
			NodeUuid  string `json:"nodeUuid"`
			CloudInfo struct {
				PrivateIp          string `json:"private_ip"`
				SecondaryPrivateIp string `json:"secondary_private_ip"`
				PublicIp           string `json:"public_ip"`
				InstanceType       string `json:"instance_type"`
				SubnetId           string `json:"subnet_id"`
				Az                 string `json:"az"`
				Region             string `json:"region"`
				Cloud              string `json:"cloud"`
				AssignPublicIP     bool   `json:"assignPublicIP"`
				UseTimeSync        bool   `json:"useTimeSync"`
			} `json:"cloudInfo"`
			AzUuid                string   `json:"azUuid"`
			PlacementUuid         string   `json:"placementUuid"`
			DisksAreMountedByUUID bool     `json:"disksAreMountedByUUID"`
			YbPrebuiltAmi         bool     `json:"ybPrebuiltAmi"`
			State                 string   `json:"state"`
			IsMaster              bool     `json:"isMaster"`
			MasterHttpPort        int      `json:"masterHttpPort"`
			MasterRpcPort         int      `json:"masterRpcPort"`
			IsTserver             bool     `json:"isTserver"`
			TserverHttpPort       int      `json:"tserverHttpPort"`
			TserverRpcPort        int      `json:"tserverRpcPort"`
			YbControllerHttpPort  int      `json:"ybControllerHttpPort"`
			YbControllerRpcPort   int      `json:"ybControllerRpcPort"`
			IsRedisServer         bool     `json:"isRedisServer"`
			RedisServerHttpPort   int      `json:"redisServerHttpPort"`
			RedisServerRpcPort    int      `json:"redisServerRpcPort"`
			IsYqlServer           bool     `json:"isYqlServer"`
			YqlServerHttpPort     int      `json:"yqlServerHttpPort"`
			YqlServerRpcPort      int      `json:"yqlServerRpcPort"`
			IsYsqlServer          bool     `json:"isYsqlServer"`
			YsqlServerHttpPort    int      `json:"ysqlServerHttpPort"`
			YsqlServerRpcPort     int      `json:"ysqlServerRpcPort"`
			NodeExporterPort      int      `json:"nodeExporterPort"`
			CronsActive           bool     `json:"cronsActive"`
			AllowedActions        []string `json:"allowedActions"`
		} `json:"nodeDetailsSet"`
		CreatingUser struct {
			Uuid               string    `json:"uuid"`
			CustomerUUID       string    `json:"customerUUID"`
			Email              string    `json:"email"`
			CreationDate       time.Time `json:"creationDate"`
			AuthTokenIssueDate int64     `json:"authTokenIssueDate"`
			Timezone           string    `json:"timezone"`
			Role               string    `json:"role"`
			IsPrimary          bool      `json:"isPrimary"`
			UserType           string    `json:"userType"`
			LdapSpecifiedRole  bool      `json:"ldapSpecifiedRole"`
		} `json:"creatingUser,omitempty"`
	} `json:"universeDetails"`
	UniverseConfig struct {
		TakeBackups    string `json:"takeBackups"`
		UseCustomImage string `json:"useCustomImage"`
	} `json:"universeConfig"`
	SampleAppCommandTxt string  `json:"sampleAppCommandTxt"`
	PricePerHour        float64 `json:"pricePerHour"`
}

type BundleResp struct {
	BundleUUID    string `json:"bundleUUID"`
	Path          string `json:"path"`
	ScopeUUID     string `json:"scopeUUID"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	BundleDetails struct {
		Components []string `json:"components"`
	} `json:"bundleDetails"`
	Status         string `json:"status"`
	CreationDate   string `json:"creationDate"`
	ExpirationDate string `json:"expirationDate"`
}

type ApiErrorResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
