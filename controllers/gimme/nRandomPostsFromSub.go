package gimme

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/R3l3ntl3ss/Meme_Api/data"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetNPostsFromSub : Get N no. of posts from a specific subreddit
func (g Controller) GetNPostsFromSub(c *gin.Context) {

	sub := strings.ToLower(c.Param("interface"))
	count, err := strconv.Atoi(c.Param("count"))

	if err != nil {
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
	memes := g.Cache.GetPostsFromCache(sub)

	// If it is not in Cache then get posts from Reddit
	if memes == nil {
		// Get 50 posts from that subreddit
		freshMemes, res := g.R.GetNPosts(sub, data.RedditPostsLimit)

		if freshMemes == nil {
			c.JSON(res.Code, res)
			return
		}

		// Remove Non Image posts from the Array
		freshMemes = utils.RemoveNonImagePosts(freshMemes)

		// Write sub posts to Cache
		g.Cache.WritePostsToCache(sub, freshMemes)

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
	res := response.MultipleMemes{
		Count: len(memes),
		Memes: memes,
	}

	c.JSON(http.StatusOK, res)
	return
}
