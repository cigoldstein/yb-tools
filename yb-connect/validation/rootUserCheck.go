package validation

import (
	"fmt"
	"os"
	"os/user"
)

func VerifyRootUser() {

	var isRootUser bool

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	if currentUser.Username == "root" {
		isRootUser = true
	}

	if !isRootUser {
		fmt.Println("ERROR: must be run as root. Please re-run as root or with sudo.")
		os.Exit(1)
	}
}
