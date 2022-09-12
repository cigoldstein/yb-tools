package platform

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"yb-get/structs"
)

func Prompt(label string, choices []structs.PromptChoices) structs.PromptChoices {

	Logger.Infof("Multiple \"%s\" available, prompting", label)
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select %s", label),
		Items: choices,
	}

	// i is the menu number item that was picked. this will be used for the index of []choices
	i, _, err := prompt.Run()

	if err != nil {
		Logger.Error("Prompt failed", err)
		os.Exit(1)
	}

	return choices[i]
}
