package token

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// Get azure token using azcli
func Get() string {
	// start with upper case if it's a public field
	var azAccountGet struct {
		AccessToken  string `json:"accessToken"`
		ExpiresOn    string `json:"expiresOn"`
		Expires_on   int    `json:"expires_on"`
		Subscription string `json:"subscription"`
		Tenant       string `json:"tenant"`
		TokenType    string `json:"tokenType"`
	}

	cmd := exec.Command("az", "account", "get-access-token")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(stdoutStderr, &azAccountGet)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s\n", azAccountGet.Subscription)

	return fmt.Sprintf("Authorization: %s %s", azAccountGet.TokenType, azAccountGet.AccessToken)
}
