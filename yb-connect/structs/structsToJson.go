package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

func YugawareAuthJsonToStruct(body []byte) YugawareAuth {

	var yugawareAuth YugawareAuth
	err := json.Unmarshal(body, &yugawareAuth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse authentication response: %v\n", err)
		os.Exit(1)
	}

	return yugawareAuth
}
