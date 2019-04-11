package main

import (
	"fmt"
	"github.com/igor-kupczynski/github-cli/auth"
	"log"
)

func main() {
	token, err := auth.NewToken()
	if err != nil {
		log.Fatalf("github-cli: %v\n", err)
	}

	fmt.Printf("Token: %#v\n", token)
}
