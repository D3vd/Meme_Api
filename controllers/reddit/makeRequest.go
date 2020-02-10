package reddit

import (
	"io/ioutil"
	"log"
	"net/http"
)

// MakeGetRequest : Makes a GET Request to Reddit API with Access Token
func MakeGetRequest(url string, accessToken string) (responseBody []byte) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "MEME_API")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.reddit.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error while making request", err)
		return nil
	}
	// Close the response body
	defer res.Body.Close()

	// Read the response
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
