package auth

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

// ask asks the user to provide a field value on the terminal.
//
// In case of an error returns the data it read so far.
func ask(field string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter %s: ", field)
	otp, _ := reader.ReadString('\n')

	return strings.TrimSpace(otp)
}

// askPassword asks the user to provide a password on the terminal.
//
// In case of an error returns an empty string.
func askPassword(field string) string {
	fmt.Printf("Enter %s: ", field)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return ""
	}

	return string(bytePassword)
}
