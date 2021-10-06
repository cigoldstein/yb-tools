package client

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/yugabyte/yb-tools/yb-connect/structs"
)

// pull from universe[0].provider and accesskeycode
func sshToNodes(universeList []structs.Universe, command string) {

	baseDirPath := "/opt/yugabyte/yugaware/data/keys/"

	providerUUID := universeList[0].UniverseDetails.Clusters[0].UserIntent.Provider
	accessKeyCode := universeList[0].UniverseDetails.Clusters[0].UserIntent.AccessKeyCode

	privateKeyPath := baseDirPath + providerUUID + "/" + accessKeyCode + ".pem"

	privateKey, err := readPrivateKey(privateKeyPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range universeList[0].UniverseDetails.NodeDetailsSet {

		publicIp := universeList[0].UniverseDetails.NodeDetailsSet[i].CloudInfo.Public_ip
		fmt.Println(publicIp)

		// TODO: adjust timeout
		conn, err := ssh.Dial("tcp", publicIp+":22", &ssh.ClientConfig{
			User:            "yugabyte",
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(privateKey)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second,
		})

		if err != nil {
			fmt.Printf("Failed to dial: %s", err)
			continue
		}
		fmt.Println("Connected!")

		session, err := conn.NewSession()

		if err != nil {
			fmt.Println(err)
		}

		buff, err := session.CombinedOutput(command)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(buff))

	}

}

func readPrivateKey(keyFile string) (ssh.Signer, error) {
	keyBytes, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(keyBytes)
}
