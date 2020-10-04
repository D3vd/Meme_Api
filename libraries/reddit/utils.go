package reddit

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
)

// EncodeCredentials : Return base64 Encoded client ID and Secret required for authentication
func EncodeCredentials() (encodedCredentials string) {
	data := ClientID + ":" + ClientSecret
	encodedCredentials = base64.StdEncoding.EncodeToString([]byte(data))
	return
}

// MakeGetRequest : Makes a GET Request to Reddit API with Access Token
func MakeGetRequest(url string) (responseBody []byte, errorCode int) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+AccessToken)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "oauth.reddit.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error while making request", err)
		return nil, res.StatusCode
	}
	// Close the response body
	defer res.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error while Parsing Response Body")
		return nil, res.StatusCode
	}

	return body, res.StatusCode
}

// GetSubredditAPIURL : Returns API Reddit URL with Limit
func GetSubredditAPIURL(subreddit string, limit int) (url string) {
	url = "https://oauth.reddit.com/r/" + subreddit + "/hot?limit=" + strconv.Itoa(limit)
	return
}
