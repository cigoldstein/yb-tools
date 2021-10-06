package structs

// Yugaware API auth
type YugawareAuth struct {
	AuthToken    string
	CustomerUUID string
	UserUUID     string
}

// Universe struct and its lists
type NodeDetails struct {
	NodeIdx   int
	NodeName  string
	CloudInfo struct {
		Private_ip string
		Public_ip  string
	}
}

type Cluster struct {
	Uuid       string
	UserIntent struct {
		Provider      string
		AccessKeyCode string
	}
}

type Universe struct {

	// universe name
	NodePrefix string

	UniverseUUID    string
	UniverseDetails struct {
		NodeDetailsSet []NodeDetails
		Clusters       []Cluster
	}
}
