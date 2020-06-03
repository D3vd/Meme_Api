package gimme

import (
	"fmt"
	"net/http"
	"strings"

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
		memes, res := g.R.GetNPosts(sub, 50)

		// Check if memes is nil because of error
		if memes == nil {
			c.JSON(res.Code, res)
			return
		}

		// Remove Non Image posts from the Array
		memes = utils.RemoveNonImagePosts(memes)

		// Write sub posts to Cache
		g.Cache.WritePostsToCache(sub, memes)
	}

	// Check if the Memes list has any posts
	if len(memes) == 0 {
		response := response.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("r/%s has no Posts with Images", sub),
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Choose one post from the list
	meme := memes[utils.GetRandomN(len(memes))]

	response := response.OneMeme{
		PostLink:  meme.PostLink,
		Subreddit: meme.SubReddit,
		Title:     meme.Title,
		URL:       meme.URL,
	}

	c.JSON(http.StatusOK, response)
	return
}
