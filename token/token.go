// token implements models and APIs for github authorizations APIs
// see https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const apiRoot = "https://api.github.com"

const ctHeader = "Content-Type"
const appjson = "application/json"

const otpHeader = "X-GitHub-OTP"

// EnsureToken tries to read the token from disk, if not exists it asks for
// credentials and stores the token on disk. If the invocation is successful,
// the caller has the token and the token is persisted on disk in `where`.
func EnsureToken(where string, name string, scopes []string) (token *Token, err error) {
	token, err = LoadToken(where)
	if err == nil && token.Note == name {
		return
	}

	token, err = NewToken(name, scopes)
	if err != nil {
		return
	}

	err = SaveToken(token, where)
	return
}

// SaveToken saves the token to disk
func SaveToken(t *Token, where string) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	file, err := os.Create(where)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(b))
	if err != nil {
		return err
	}

	return nil
}

// LoadToken loads the token from disk
func LoadToken(where string) (*Token, error) {
	file, err := os.Open(where)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, file)
	if err != nil {
		return nil, err
	}

	var t Token
	err = json.Unmarshal(b.Bytes(), &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// NewToken interactively asks for credentials and uses them
// to acquire github personal access token
func NewToken(name string, scopes []string) (*Token, error) {
	req := &TokenRequest{
		Scopes: scopes,
		Note:   name,
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
