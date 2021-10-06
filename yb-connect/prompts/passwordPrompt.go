package prompts

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func PasswordPrompt() string {

	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}

	return string(password)
}
