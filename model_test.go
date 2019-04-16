package ghtoken

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestMarshallRequest(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want TokenRequest
	}{
		{
			name: "should marshal the TokenRequest zero value",
			in:   `{}`,
			want: TokenRequest{},
		},
		{
			name: "should marshal the hydrated TokenRequest",
			in: `
{
	"note": "personal access token for app foobar",
	"scopes": ["repo", "gist"]
}`[1:],
			want: TokenRequest{
				Note:   "personal access token for app foobar",
				Scopes: []string{"repo", "gist"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// unmarshall
			var tokenRequest TokenRequest
			err := json.Unmarshal([]byte(test.in), &tokenRequest)
			if err != nil {
				t.Fatalf("Can't unmarhall: %#v\n", err)
			}
			if !reflect.DeepEqual(tokenRequest, test.want) {
				t.Fatalf("Want: %#v,\nhave: %#v\n", test.want, tokenRequest)
			}

			// marshall -- symmetry
			bytes, err := json.Marshal(tokenRequest)
			if err != nil {
				t.Fatalf("Can't marhall: %#v\n", err)
			}
			var tokenRequest2 TokenRequest
			err = json.Unmarshal(bytes, &tokenRequest2)
			if err != nil {
				t.Fatalf("Can't unmarhall: %#v\n", err)
			}
			if !reflect.DeepEqual(tokenRequest2, test.want) {
				t.Fatalf("Want: %#v,\nhave: %#v\n", test.want, tokenRequest2)
			}

		})
	}
}

var (
	updated, _ = time.Parse(time.RFC3339, "2011-09-06T20:39:23Z")
	created, _ = time.Parse(time.RFC3339, "2011-09-06T17:26:27Z")
)

func TestMarshall(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Token
	}{
		{
			name: "should marshal the Token zero value",
			in:   `{}`,
			want: Token{},
		},
		{
			name: "should marshal the hydrated Token",
			in: `
{
  "id": 1,
  "url": "https://api.github.com/authorizations/1",
  "scopes": [
    "public_repo"
  ],
  "token": "abcdefgh12345678",
  "token_last_eight": "12345678",
  "hashed_token": "25f94a2a5c7fbaf499c665bc73d67c1c87e496da8985131633ee0a95819db2e8",
  "app": {
    "url": "http://my-github-app.com",
    "name": "my github app",
    "client_id": "abcde12345fghij67890"
  },
  "note": "optional note",
  "updated_at": "2011-09-06T20:39:23Z",
  "created_at": "2011-09-06T17:26:27Z"
}`[1:],
			want: Token{
				Id:             1,
				Url:            "https://api.github.com/authorizations/1",
				Scopes:         []string{"public_repo"},
				Token:          "abcdefgh12345678",
				TokenLastEight: "12345678",
				HashedToken:    "25f94a2a5c7fbaf499c665bc73d67c1c87e496da8985131633ee0a95819db2e8",
				App: App{
					Url:      "http://my-github-app.com",
					Name:     "my github app",
					ClientId: "abcde12345fghij67890",
				},
				Note:      "optional note",
				UpdatedAt: updated,
				CreatedAt: created,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// unmarshall
			var token Token
			err := json.Unmarshal([]byte(test.in), &token)
			if err != nil {
				t.Fatalf("Can't unmarhall: %#v\n", err)
			}
			if !reflect.DeepEqual(token, test.want) {
				t.Fatalf("Want: %#v,\nhave: %#v\n", test.want, token)
			}

			// marshall -- symmetry
			bytes, err := json.Marshal(token)
			if err != nil {
				t.Fatalf("Can't marhall: %#v\n", err)
			}
			var token2 Token
			err = json.Unmarshal(bytes, &token2)
			if err != nil {
				t.Fatalf("Can't unmarhall: %#v\n", err)
			}
			if !reflect.DeepEqual(token2, test.want) {
				t.Fatalf("Want: %#v,\nhave: %#v\n", test.want, token2)
			}

		})
	}
}
