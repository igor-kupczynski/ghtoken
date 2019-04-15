// main contains a demo program for the github-token library/
// User can find their list of defined tokens in https://github.com/settings/tokens
package main

import (
	"fmt"
	"github.com/igor-kupczynski/github-cli/token"
	"log"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("github-token: %v\n", err)
	}

	t, err := token.EnsureToken(
		fmt.Sprintf("%s/.github-token.json", home),
		"github.com/igor-kupczynski/github-token",
		[]string{"repo"},
	)
	if err != nil {
		log.Fatalf("github-token: %v\n", err)
	}

	fmt.Printf("Token: %#v\n", t)
}
