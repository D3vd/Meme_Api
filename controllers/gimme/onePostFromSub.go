package gimme

import (
	"fmt"
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
			Code:    http.StatusServiceUnavailable,
			Message: "Error while getting memes from subreddit. Please try again",
		}

		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	// Remove Non Image posts from the Array
	memes = utils.RemoveNonImagePosts(memes)

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
