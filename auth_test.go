package goebay

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAuth(t *testing.T) {

	err := godotenv.Load()

	if err != nil {
		t.Error(err)
	}

  	if err != nil {
		t.Error("Test Failed: error getting data from .env")
  	}

	clientID := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")

	resp, err := GetOAuthToken(OAuthParams{clientID: clientID, clientSecret: clientSecret, IsSandbox: true})

	if err != nil {
		t.Error(err)
	}
	
	if len(resp) <= 0 {
		t.Error("no response")
	}
}