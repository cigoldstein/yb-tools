package client

import (
	"fmt"

	"github.com/yugabyte/yb-tools/yb-connect/validation"
	"github.com/yugabyte/yb-tools/yb-connect/yugawareApi"
)

func YbConnect(ywHost string, command string) {

	// make sure we're running as root or with sudo
	validation.VerifyRootUser()

	// Yugaware API actions
	yugawareUsername, yugawarePassword := yugawareApi.ConfigHttpClient(ywHost)
	fmt.Println()
	yugawareAuth := yugawareApi.YugawareLogin(yugawareUsername, yugawarePassword)
	universeList := yugawareApi.GetUniverseList(yugawareAuth)

	// ssh to each node
	sshToNodes(universeList, command)

}
