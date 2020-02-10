package reddit

import (
	"encoding/json"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit/auth"
	"github.com/R3l3ntl3ss/Meme_Api/models"

	redditModels "github.com/R3l3ntl3ss/Meme_Api/models/reddit"
)

func GetMemes(subreddit string) []models.Meme {

	accessToken := auth.GetAccessToken()
	url := GetSubredditURL(subreddit, 10)

	body := MakeGetRequest(url, accessToken)

	var redditResponse redditModels.RedditResponse

	json.Unmarshal(body, &redditResponse)

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
