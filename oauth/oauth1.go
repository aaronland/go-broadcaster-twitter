package oauth

import (
	"context"
	"encoding/json"
	"fmt"
)

// type OAuth1Credentials is a struct containing OAuth1 keys and secrets.
type OAuth1Credentials struct {
	// ConsumerKey is an OAuth1 consumer (application) key.	
	ConsumerKey    string `json:"consumer_key"`
	// ConsumerSecret is an OAuth1 consumer (application) secret.	
	ConsumerSecret string `json:"consumer_secret"`
	// AccessToken is an OAuth1 access token.		
	AccessToken    string `json:"access_token"`
	// AccessSecret is an OAuth1 access secret.			
	AccessSecret   string `json:"access_token_secret"`
}

// NewOAuth1CredentialsFromString derives a `OAuth1Credentials` struct from a JSON-encoded string.
func NewOAuth1CredentialsFromString(ctx context.Context, str_creds string) (*OAuth1Credentials, error) {

	var creds *OAuth1Credentials

	err := json.Unmarshal([]byte(str_creds), &creds)

	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal credentials, %w", err)
	}

	return creds, nil
}
