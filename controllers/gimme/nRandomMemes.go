package gimme

import (
	"net/http"
	"strconv"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/data"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetNRandomMemes : Returns N no. of memes from a random subreddit
func (g Controller) GetNRandomMemes(c *gin.Context) {

	count, _ := strconv.Atoi(c.Param("interface"))

	// Check if the count is less than 50
	if count > 50 {
		count = 50
	}

	// Choose a random meme subreddit
	sub := data.MemeSubreddits[utils.GetRandomN(len(data.MemeSubreddits))]

	// Check if the sub is present in the cache
	memes := g.Cache.GetPostsFromCache(sub)

	// If it is not in Cache then get posts from Reddit
	if memes == nil {
		// Get 50 posts from that subreddit
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
	memesLen := len(memes)

	if memesLen == 0 {
		res := response.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while getting Memes",
		}

		c.JSON(http.StatusInternalServerError, res)
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
