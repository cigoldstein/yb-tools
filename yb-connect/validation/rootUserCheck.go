package validation

import (
	"fmt"
	"os/user"
)

func VerifyRootUser() error {

	var isRootUser bool

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	if currentUser.Username == "root" {
		isRootUser = true
	}

	if !isRootUser {
		return fmt.Errorf("must be run as root. Please re-run as root or with sudo.")
	}

	return nil
}
