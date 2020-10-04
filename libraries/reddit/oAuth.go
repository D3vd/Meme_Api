package reddit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	rm "Meme_Api/libraries/reddit/models"

	"github.com/getsentry/sentry-go"
)

// GetAccessToken : Get temporary Access Token based on App client ID and Secret
func GetAccessToken() (accessToken string) {

	encodedCredentials := EncodeCredentials()

	// Reddit URL to get access token
	url := "https://www.reddit.com/api/v1/access_token"

	// Set grant type to client_credentials as POST body
	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error while Creating Request")
		return ""
	}

	// Set Headers including the User Agent and the Authorization with the encoded credentials
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", "Basic "+encodedCredentials)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.reddit.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "29")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	// Make Request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error while making connecting to : " + url)
		return ""
	}

	// Close the response body
	defer res.Body.Close()

	// Read the response
	body, _ := ioutil.ReadAll(res.Body)

	var accessTokenBody rm.AccessTokenBody

	if err := json.Unmarshal(body, &accessTokenBody); err != nil {
		sentry.CaptureException(err)
		log.Println("Error while Unmarshalling AccessTokenBody", err)
		return ""
	}

	return accessTokenBody.AccessToken
}
