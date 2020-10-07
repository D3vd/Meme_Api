package gimme

import (
	"fmt"
	"net/http"
	"strings"

	"Meme_Api/data"
	"Meme_Api/libraries/reddit"
	"Meme_Api/libraries/redis"

	"Meme_Api/models/response"
	"Meme_Api/utils"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// GetOnePostFromSub : Get one post from a specific subreddit
func GetOnePostFromSub(c *gin.Context) {

	sub := strings.ToLower(c.Param("interface"))

	// Check if the sub is present in the cache
	memes := redis.GetPostsFromCache(sub)

	// If it is not in Cache then get posts from Reddit
	if memes == nil {
		// Get 50 posts from that Subreddit
		freshMemes, res := reddit.GetNPosts(sub, data.RedditPostsLimit)

		// Check if memes is nil because of error
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
	if len(memes) == 0 {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("r/%s has no Posts with Images", sub),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Choose one post from the list
	meme := memes[utils.GetRandomN(len(memes))]

	res := response.OneMeme{
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

	c.JSON(http.StatusOK, res)
}
