package gimme

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetNPostsFromSub : Get N no. of posts from a specific subreddit
func (g Controller) GetNPostsFromSub(c *gin.Context) {

	sub := strings.ToLower(c.Param("interface"))
	count, err := strconv.Atoi(c.Param("count"))

	if err != nil {
		response := response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Count Value",
		}

		c.JSON(http.StatusBadRequest, response)
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
		memes = g.R.GetNPosts(sub, 50)

		if memes == nil {
			response := response.Error{
				Code:    http.StatusServiceUnavailable,
				Message: "Error while getting memes from subreddit. Please try again",
			}

			c.JSON(http.StatusServiceUnavailable, response)
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

	// Get N no. of posts from that list
	memes = utils.GetNRandomMemes(memes, count)

	var memesResponse []response.OneMeme

	for _, meme := range memes {
		memeResponse := response.OneMeme{
			PostLink:  meme.PostLink,
			Subreddit: meme.SubReddit,
			Title:     meme.Title,
			URL:       meme.URL,
		}

		memesResponse = append(memesResponse, memeResponse)
	}

	response := response.MultipleMemes{
		Count: len(memesResponse),
		Memes: memesResponse,
	}

	c.JSON(http.StatusOK, response)
	return
}
