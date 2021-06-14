package gimme

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"Meme_Api/data"
	"Meme_Api/libraries/reddit"
	"Meme_Api/libraries/redis"

	"Meme_Api/models/response"
	"Meme_Api/utils"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// GetNPostsFromSub : Get N no. of posts from a specific subreddit
func GetNPostsFromSub(c *gin.Context) {

	sub := strings.ToLower(c.Param("interface"))
	count, err := strconv.Atoi(c.Param("count"))

	if err != nil || count <= 0 {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Count Value",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Check if the count is less than 50
	if count > 50 {
		count = 50
	}

	// Check if the sub is present in the cache
	memes := redis.GetPostsFromCache(sub)

	// If it is not in Cache then get posts from Reddit
	if memes == nil {
		// Get 50 posts from that subreddit
		freshMemes, res := reddit.GetNPosts(sub, data.RedditPostsLimit)

		if freshMemes == nil {
			c.JSON(res.Code, res)
			return
		}

		// Remove Non Image posts from the Array
		freshMemes = utils.RemoveNonImagePosts(freshMemes)

		// Write sub posts to Cache
		if err := redis.WritePostsToCache(sub, freshMemes); err != nil {
			sentry.CaptureException(err)
		}

		// Set Memes to Fresh Memes
		memes = freshMemes
	}

	// Check if the Memes list has any posts
	memesLen := len(memes)

	if memesLen == 0 {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("r/%s has no Posts with Images", sub),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Check if the returned memes length is lesser count
	if memesLen < count {
		count = memesLen
	}

	// Get N no. of posts from that list
	memes = utils.GetNRandomMemes(memes, count)

	var memesResponse []response.OneMeme

	for _, meme := range memes {
		memeResponse := response.OneMeme{
			PostLink:  meme.PostLink,
			Subreddit: meme.SubReddit,
			Title:     meme.Title,
			URL:       meme.URL,
			NSFW:      meme.NSFW,
			Spoiler:   meme.Spoiler,
			Author:    meme.Author,
			Ups:       meme.Ups,
			Preview:   meme.Preview,
		}

		memesResponse = append(memesResponse, memeResponse)
	}

	res := response.MultipleMemes{
		Count: len(memesResponse),
		Memes: memesResponse,
	}

	c.JSON(http.StatusOK, res)
}
