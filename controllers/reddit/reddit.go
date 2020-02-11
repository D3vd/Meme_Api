package reddit

import (
	"log"
	"os"
)

// Reddit : Reddit structure with reddit credentials
type Reddit struct {
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	UserAgent    string `json:"user_agent"`
}

// Init : Initialize the Reddit Structure with App Credentials
func (r *Reddit) Init() {
	// Get Reddit Client Credentials from the environment variables
	clientID := os.Getenv("REDDIT_CLIENT_ID")
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatalln("Reddit Client ID and Secret have not been set")
		return
	}

	r.ClientID = clientID
	r.ClientSecret = clientSecret

	r.UserAgent = "MEME_API"

	accessToken := r.GetAccessToken()

	if accessToken == "" {
		log.Fatalln("Error while getting Access Token")
		return
	}

	r.AccessToken = accessToken
}
