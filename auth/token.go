package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const apiRoot = "https://api.github.com"

const ctHeader = "Content-Type"
const appjson = "application/json"

const otpHeader = "X-GitHub-OTP"

// NewToken interactively asks for credentials and uses them
// to acquire github personal access token
func NewToken() (*Token, error) {
	req := &TokenRequest{
		Scopes: []string{"public_repo"},
		Note:   "igor-kupczynski/github-cli",
	}

	fmt.Print("Provide your github credentials. " +
		"They will be used to retrieve the token and won't be stored.\n")

	username := ask("username")
	password := askPassword("password")

	t, err := createToken(req, username, password, "")
	authError, ok := err.(*AuthError)
	if ok && authError.OtpNeeded {
		otp := ask("one-time password")
		t, err = createToken(req, username, password, otp)
	}

	return t, err
}

type AuthError struct {
	OtpNeeded bool
}

func (a *AuthError) Error() string {
	otp := ""
	if a.OtpNeeded {
		otp = ", please provide one-time password"
	}
	return fmt.Sprintf("401 Unauthorized%s", otp)
}

func createToken(tokenRequest *TokenRequest, username, password, otp string) (*Token, error) {
	buf, err := json.Marshal(tokenRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiRoot+"/authorizations", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Add(ctHeader, appjson)
	if otp != "" {
		req.Header.Add(otpHeader, otp)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 401 {
		otpNeeded := false
		if otp := resp.Header.Get("X-GitHub-OTP"); strings.HasPrefix(otp, "required;") {
			otpNeeded = true
		}
		return nil, &AuthError{OtpNeeded: otpNeeded}
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("can't get access token with request [%s]: %s", buf, resp.Status)
	}

	defer resp.Body.Close()

	var token Token
	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
