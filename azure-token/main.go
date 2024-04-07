package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("az", "account", "get-access-token")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// start with upper case if it's a public field
	var azAccountGet struct {
		AccessToken  string `json:"accessToken"`
		ExpiresOn    string `json:"expiresOn"`
		Expires_on   int    `json:"expires_on"`
		Subscription string `json:"subscription"`
		Tenant       string `json:"tenant"`
		TokenType    string `json:"tokenType"`
	}

	err = json.Unmarshal(stdoutStderr, &azAccountGet)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s\n", azAccountGet.Subscription)
	fmt.Printf("Authorization: %s %s", azAccountGet.TokenType, azAccountGet.AccessToken)
}