package reddit

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/R3l3ntl3ss/Meme_Api/models"
	"github.com/R3l3ntl3ss/Meme_Api/models/reddit"
)

// GetComments : Get (N) no. of posts from Reddit with Subreddit Name and Limit
func (r *Reddit) GetNComments(id string, limit int) ([]models.Comment, response.Error) {
	url := "https://oauth.reddit.com/comments/" + id + "?limit=" + strconv.Itoa(limit)
	body, _ := r.MakeGetRequest(url)

	var redditResponse reddit.Comments
	if err := json.Unmarshal(body, &redditResponse); err != nil {
		log.Println("Error while Parsing Reddit Response")

		res := response.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while getting memes from subreddit. Please try again",
		}

		return nil, res
	}

	var comments []models.Comment
	for _, post := range redditResponse[1].Data.Children {
		comment := models.Comment{
			Author:    post.Data.Author,
			Ups:       post.Data.Ups,
			Body:      post.Data.Body,

		}
		comments = append(comments, comment)
        }

	return comments, response.Error{}
}

// GetNPosts : Get (N) no. of posts from Reddit with Subreddit Name and Limit
func (r *Reddit) GetNPosts(subreddit string, count int) ([]models.Meme, response.Error) {

	url := GetSubredditAPIURL(subreddit, count)

	body, statusCode := r.MakeGetRequest(url)

	// Check if the access Token has been expired
	if statusCode == 401 {
		// Get new access token
		r.GetNewAccessToken()

		// Make request Again
		body, _ = r.MakeGetRequest(url)
	}

	// Handle Subreddit Errors for Forbidden and Not Found
	if statusCode == 403 {
		res := response.Error{
			Code:    http.StatusForbidden,
			Message: "Unable to Access Subreddit. Subreddit is Locked or Private",
		}

		return nil, res
	}

	if statusCode == 404 {
		res := response.Error{
			Code:    http.StatusNotFound,
			Message: "This subreddit does not exist.",
		}

		return nil, res
	}

	var redditResponse reddit.Response

	if err := json.Unmarshal(body, &redditResponse); err != nil {
		log.Println("Error while Parsing Reddit Response")

		res := response.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while getting memes from subreddit. Please try again",
		}

		return nil, res
	}

	// Check if there are posts in the subreddit
	if len(redditResponse.Data.Children) == 0 {
		res := response.Error{
			Code:    http.StatusNotFound,
			Message: "This subreddit has no posts or doesn't exist.",
		}

		return nil, res
	}

	var memes []models.Meme

	for _, post := range redditResponse.Data.Children {
                comments, _ := r.GetNComments(post.Data.ID, 10)
		meme := models.Meme{
			Title:     post.Data.Title,
			PostLink:  post.Data.GetShortLink(),
			SubReddit: post.Data.Subreddit,
			URL:       post.Data.URL,
			Author:    post.Data.Author,
			Ups:       post.Data.Ups,
			NSFW:      post.Data.Over18,
			Spoiler:   post.Data.Spoiler,
			Comments:  comments,
		}

		memes = append(memes, meme)
	}

	return memes, response.Error{}
}
