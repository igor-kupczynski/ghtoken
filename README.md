# ghtoken

A simple library to get github personal access token.

Can request new token from github, store it on disk and load it from disk.
Additionally, it can ask the user for their login credentials in terminal.

User can find their list of defined tokens in [https://github.com/settings/tokens]().

Demo app:

```go
package main

import (
	"fmt"
	"github.com/igor-kupczynski/ghtoken"
	"log"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ghtoken: %v\n", err)
	}

	t, err := ghtoken.EnsureToken(
		fmt.Sprintf("%s/.github-token.json", home),
		"github.com/igor-kupczynski/ghtoken",
		[]string{"repo"},
	)
	if err != nil {
		log.Fatalf("ghtoken: %v\n", err)
	}

	fmt.Printf("Token: %#v\n", t)
}
```
