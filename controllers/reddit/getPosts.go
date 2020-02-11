package reddit

import (
	"encoding/json"
	"log"

	"github.com/R3l3ntl3ss/Meme_Api/models"

	redditModels "github.com/R3l3ntl3ss/Meme_Api/models/reddit"
)

func (r Reddit) GetNPosts(subreddit string, count int) []models.Meme {

	url := GetSubredditAPIURL(subreddit, count)

	body := r.MakeGetRequest(url)

	var redditResponse redditModels.RedditResponse

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
