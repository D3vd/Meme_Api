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

	// TODO -- Check if the subreddit data is present in Cache

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

	// TODO -- Write subreddit posts to Cache

	g.Cache.WritePostsToCache(sub, memes)

	// Check if the Memes list has any posts
	if len(memes) == 0 {
		response := response.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while getting Memes",
		}

		c.JSON(http.StatusInternalServerError, response)
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
