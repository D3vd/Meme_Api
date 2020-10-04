package reddit

import (
	"log"
	"os"
)

// Reddit : Reddit structure with reddit credentials
var (
	AccessToken  string
	ClientID     string
	ClientSecret string
	UserAgent    string
)

// Init : Initialize the Reddit Structure with App Credentials
func Init() {
	// Get Reddit Client Credentials from the environment variables
	clientID := os.Getenv("REDDIT_CLIENT_ID")
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatalln("Reddit Client ID and Secret have not been set")
		return
	}

	ClientID = clientID
	ClientSecret = clientSecret

	UserAgent = "MEME_API"

	accessToken := GetAccessToken()

	if accessToken == "" {
		log.Fatalln("Error while getting Access Token")
		return
	}

	AccessToken = accessToken
}

// GetNewAccessToken : Function to Generate New Access Token once the old one expires
func GetNewAccessToken() (ok bool) {
	newAccessToken := GetAccessToken()

	if newAccessToken == "" {
		log.Println("Unable to get new Access Token")
		return false
	}

	AccessToken = newAccessToken
	return true
}
