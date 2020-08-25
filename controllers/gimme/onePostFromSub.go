package gimme

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/R3l3ntl3ss/Meme_Api/data"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetOnePostFromSub : Get one post from a specific subreddit
func (g Controller) GetOnePostFromSub(c *gin.Context) {

	sub := strings.ToLower(c.Param("interface"))

	// Check if the sub is present in the cache
	memes := g.Cache.GetPostsFromCache(sub)

	// If it is not in Cache then get posts from Reddit
	if memes == nil {
		// Get 50 posts from that Subreddit
		freshMemes, res := g.R.GetNPosts(sub, data.RedditPostsLimit)

		// Check if memes is nil because of error
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
	c.JSON(http.StatusOK, meme)
	return
}
