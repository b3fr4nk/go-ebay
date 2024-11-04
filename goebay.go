package goebay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	tokenURL = "https://api.ebay.com/identity/v1/oauth2/token"
	sandboxTokenURL = "https://api.sandbox.ebay.com/identity/v1/oauth2/token"
)

type OAuthResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}

type OAuthParams struct {
	clientID string
	clientSecret string
	IsSandbox bool
}

func encodeCredentials(clientID, clientSecret string) string {
    auth := clientID + ":" + clientSecret
    return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetOAuthToken(p OAuthParams) (string, error) {
	data := []byte("grant_type=client_credentials&scope=https://api.ebay.com/oauth/api_scope")
	var reqURL string
	if p.IsSandbox {
		reqURL = sandboxTokenURL
	} else {
		reqURL = tokenURL
	}
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodeCredentials(p.clientID, p.clientSecret))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get token %s", body)
	}

	var result OAuthResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.AccessToken, nil

}