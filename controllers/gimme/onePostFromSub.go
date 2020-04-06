package gimme

import (
	"net/http"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
	"github.com/gin-gonic/gin"
)

// GetOnePostFromSub : Get one post from a specific subreddit
func (g Controller) GetOnePostFromSub(c *gin.Context) {

	sub := c.Param("interface")

	// Get 50 posts from that Subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Check if memes is nil because of error
	if memes == nil {
		response := response.Error{
			Code:    500,
			Message: "Error while getting memes from subreddit. Please try again",
		}

		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	// Choose one post from the list
	meme := memes[utils.GetRandomN(len(memes))]

	response := response.OneMemeResponse{
		PostLink:  meme.PostLink,
		Subreddit: meme.SubReddit,
		Title:     meme.Title,
		URL:       meme.URL,
	}

	c.JSON(http.StatusOK, response)
	return
}
