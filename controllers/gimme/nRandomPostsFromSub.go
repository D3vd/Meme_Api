package gimme

import (
	"net/http"
	"strconv"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetNPostsFromSub : Get N no. of posts from a specific subreddit
func (g Controller) GetNPostsFromSub(c *gin.Context) {

	sub := c.Param("interface")
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

	// Get 50 posts from that subreddit
	memes := g.R.GetNPosts(sub, 50)

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
