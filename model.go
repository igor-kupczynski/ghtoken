package ghtoken

import "time"

// TokenRequest represents a request to create new personal access token
// as per https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
type TokenRequest struct {
	Scopes []string `json:"scopes,omitempty"`
	Note   string   `json:"note,omitempty"`
}

// Token represents a reponse to personal access token creation request
// as per https://developer.github.com/v3/oauth_authorizations/#response-5
type Token struct {
	Id             int       `json:"id,omitempty"`
	Url            string    `json:"url,omitempty"`
	Scopes         []string  `json:"scopes,omitempty"`
	Token          string    `json:"token,omitempty"`
	TokenLastEight string    `json:"token_last_eight,omitempty"`
	HashedToken    string    `json:"hashed_token,omitempty"`
	App            App       `json:"app,omitempty"`
	Note           string    `json:"note,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

// App is part of the Token reponse
type App struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	ClientId string `json:"client_id"`
}
