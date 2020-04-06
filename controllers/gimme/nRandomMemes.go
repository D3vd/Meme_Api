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

	// Get 50 posts from that subreddit
	memes := g.R.GetNPosts(sub, 50)

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
