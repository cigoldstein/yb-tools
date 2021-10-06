package prompts

import (
	"fmt"
)

func UsernamePrompt() string {

	var Username string

	fmt.Print("Yugaware Username: ")
	fmt.Scanln(&Username)

	return Username

}
