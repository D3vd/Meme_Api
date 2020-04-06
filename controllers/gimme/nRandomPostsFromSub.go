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
		// TODO: Handle Error properly
		c.JSON(http.StatusBadRequest, "Give proper count value")
		return
	}

	// Check if the count is less than 50
	if count > 50 {
		count = 50
	}

	// Get 50 posts from that subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Get N no. of posts from that list
	memes = utils.GetNRandomMemes(memes, count)

	var memesResponse []response.OneMemeResponse

	for _, meme := range memes {
		memeResponse := response.OneMemeResponse{
			PostLink:  meme.PostLink,
			Subreddit: meme.SubReddit,
			Title:     meme.Title,
			URL:       meme.URL,
		}

		memesResponse = append(memesResponse, memeResponse)
	}

	response := response.MultipleMemesResponse{
		Count: len(memesResponse),
		Memes: memesResponse,
	}

	c.JSON(http.StatusOK, response)
	return
}
