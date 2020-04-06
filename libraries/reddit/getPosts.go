package reddit

import (
	"encoding/json"
	"log"

	"github.com/R3l3ntl3ss/Meme_Api/models"

	"github.com/R3l3ntl3ss/Meme_Api/models/reddit"
)

// GetNPosts : Get (N) no. of posts from Reddit with Subreddit Name and Limit
func (r *Reddit) GetNPosts(subreddit string, count int) []models.Meme {

	url := GetSubredditAPIURL(subreddit, count)

	body, statusCode := r.MakeGetRequest(url)

	// Check if the access Token has been expired
	if statusCode == 401 {
		// Get new access token
		r.GetNewAccessToken()

		// Make request Again
		body, _ = r.MakeGetRequest(url)
	}

	var redditResponse reddit.Response

	if err := json.Unmarshal(body, &redditResponse); err != nil {
		log.Println("Error while Parsing Reddit Response")
		return nil
	}

	var memes []models.Meme

	for _, post := range redditResponse.Data.Children {
		meme := models.Meme{
			Title:     post.Data.Title,
			PostLink:  post.Data.GetShortLink(),
			SubReddit: post.Data.Subreddit,
			URL:       post.Data.URL,
		}

		memes = append(memes, meme)
	}

	return memes
}
