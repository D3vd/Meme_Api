package gimme

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/utils"
	"github.com/R3l3ntl3ss/Meme_Api/data"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
)

// GetOneRandomMeme : Returns a single meme from a random subreddit
func (g Controller) GetOneRandomMeme(c *gin.Context) {

	// Choose Random Meme Subreddit
	sub := data.MemeSubreddits[utils.GetRandomN(len(data.MemeSubreddits))]

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
