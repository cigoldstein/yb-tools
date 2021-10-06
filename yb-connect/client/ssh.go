package client

import (
	"fmt"
	"os"
	"time"

	"github.com/go-logr/logr"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
	"golang.org/x/crypto/ssh"
)

// pull from universe[0].provider and accesskeycode
func SshToNodes(log logr.Logger, universe *models.UniverseResp, command string) error {

	baseDirPath := "/opt/yugabyte/yugaware/data/keys/"

	providerUUID := universe.UniverseDetails.Clusters[0].UserIntent.Provider
	accessKeyCode := universe.UniverseDetails.Clusters[0].UserIntent.AccessKeyCode

	privateKeyPath := baseDirPath + providerUUID + "/" + accessKeyCode + ".pem"

	privateKey, err := readPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}

	for _, node := range universe.UniverseDetails.NodeDetailsSet {
		runSshCommand(log, node, privateKey, command)
	}

	return nil
}

func runSshCommand(log logr.Logger, node *models.NodeDetailsResp, privateKey ssh.Signer, command string) {

	publicIp := node.CloudInfo.PublicIP
	log.Info("connecting to host", "host", node.CloudInfo)

	// TODO: adjust timeout
	conn, err := ssh.Dial("tcp", publicIp+":22", &ssh.ClientConfig{
		User:            "yugabyte",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(privateKey)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second,
	})

	if err != nil {
		log.Error(err, "failed to connect to host")
		return
	}

	defer conn.Close()

	fmt.Println("Connected!")

	session, err := conn.NewSession()

	if err != nil {
		log.Error(err, "failed to create session")
		return
	}

	buff, err := session.CombinedOutput(command)

	if err != nil {
		log.Error(err, "unable to execute command", "out", string(buff))
		return
	}

	fmt.Println(string(buff))

}

func readPrivateKey(keyFile string) (ssh.Signer, error) {
	keyBytes, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(keyBytes)
}
