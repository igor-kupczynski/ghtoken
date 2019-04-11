package auth

import "time"

type TokenRequest struct {
	Scopes       []string `json:"scopes,omitempty"`
	Note         string   `json:"note,omitempty"`
	NoteUrl      string   `json:"note_url,omitempty"`
	ClientId     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
	Fingerprint  string   `json:"fingerprint,omitempty"`
}

type App struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	ClientId string `json:"client_id"`
}

type Token struct {
	Id             int       `json:"id,omitempty"`
	Url            string    `json:"url,omitempty"`
	Scopes         []string  `json:"scopes,omitempty"`
	Token          string    `json:"token,omitempty"`
	TokenLastEight string    `json:"token_last_eight,omitempty"`
	HashedToken    string    `json:"hashed_token,omitempty"`
	App            App       `json:"app,omitempty"`
	Note           string    `json:"note,omitempty"`
	NoteUrl        string    `json:"note_url,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	Fingerprint    string    `json:"fingerprint,omitempty"`
}
