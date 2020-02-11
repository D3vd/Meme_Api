package controllers

import (
	"github.com/R3l3ntl3ss/Meme_Api/data"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit"
)

// GimmeController : Gives random meme(s) through /gimme endpoint
type GimmeController struct {
	R reddit.Reddit
}

// GetOneRandomMeme : Returns a single meme from a random subreddit
func (g GimmeController) GetOneRandomMeme(c *gin.Context) {

	// Choose Random Meme Subreddit
	sub := data.MemeSubreddits[GetRandomN(len(data.MemeSubreddits))]

	// Get 50 posts from that Subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Choose one post from the list
	meme := memes[GetRandomN(len(memes))]

	c.JSON(http.StatusOK, meme)
	return
}
