package controllers

import (
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

	memes := g.R.GetNPosts("memes", 100)

	c.JSON(http.StatusOK, memes)
	return
}
