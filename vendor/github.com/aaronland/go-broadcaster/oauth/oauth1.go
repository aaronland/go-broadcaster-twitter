package oauth

import (
	"context"
	"encoding/json"
)

type OAuth1Credentials struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken    string `json:"access_token"`
	AccessSecret   string `json:"access_token_secret"`
}

func NewOAuth1CredentialsFromString(ctx context.Context, str_creds string) (*OAuth1Credentials, error) {

	var creds *OAuth1Credentials

	err := json.Unmarshal([]byte(str_creds), &creds)

	if err != nil {
		return nil, err
	}

	return creds, nil
}
