package reddit

import (
	"encoding/json"
	"log"
	"net/http"

	rm "Meme_Api/libraries/reddit/models"
	"Meme_Api/models"

	"github.com/getsentry/sentry-go"
)

// GetNPosts : Get (N) no. of posts from Reddit with Subreddit Name and Limit
func GetNPosts(subreddit string, count int) ([]models.Meme, rm.CustomRedditError) {

	url := GetSubredditAPIURL(subreddit, count)

	body, statusCode := MakeGetRequest(url)

	// Check if the access Token has been expired
	if statusCode == 401 {
		// Get new access token
		GetNewAccessToken()

		// Make request Again
		body, statusCode = MakeGetRequest(url)
	}

	// If Reddit is down return 503 error
	if statusCode == 500 {
		res := rm.CustomRedditError{
			Code:    http.StatusServiceUnavailable,
			Message: "Reddit is unreachable at the moment",
		}

		sentry.CaptureMessage("Reddit is Down!!")
		return nil, res
	}

	// Handle Subreddit Errors for Forbidden and Not Found
	if statusCode == 403 {
		res := rm.CustomRedditError{
			Code:    http.StatusForbidden,
			Message: "Unable to Access Subreddit. Subreddit is Locked or Private",
		}

		return nil, res
	}

	if statusCode == 404 {
		res := rm.CustomRedditError{
			Code:    http.StatusNotFound,
			Message: "This subreddit does not exist.",
		}

		return nil, res
	}

	if statusCode != 200 {
		res := rm.CustomRedditError{
			Code:    http.StatusServiceUnavailable,
			Message: "Unknown error while getting posts. Please try again",
		}

		sentry.CaptureMessage(string(body))
		return nil, res
	}

	var redditResponse rm.Response

	if err := json.Unmarshal(body, &redditResponse); err != nil {
		sentry.CaptureException(err)
		log.Println("Error while Parsing Reddit Response")

		res := rm.CustomRedditError{
			Code:    http.StatusInternalServerError,
			Message: "Error while getting memes from subreddit. Please try again",
		}

		return nil, res
	}

	// Check if there are posts in the subreddit
	if len(redditResponse.Data.Children) == 0 {
		res := rm.CustomRedditError{
			Code:    http.StatusNotFound,
			Message: "This subreddit has no posts or doesn't exist.",
		}

		return nil, res
	}

	var memes []models.Meme

	for _, post := range redditResponse.Data.Children {
		meme := models.Meme{
			Title:     post.Data.Title,
			PostLink:  post.Data.GetShortLink(),
			SubReddit: post.Data.Subreddit,
			URL:       post.Data.URL,
			Author:    post.Data.Author,
			Ups:       post.Data.Ups,
			NSFW:      post.Data.Over18,
			Spoiler:   post.Data.Spoiler,
			Preview:   post.Data.GetCleanPreviewImages(),
		}

		memes = append(memes, meme)
	}

	return memes, rm.CustomRedditError{}
}
