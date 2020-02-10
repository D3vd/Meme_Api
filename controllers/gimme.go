package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit"
)

// GimmeController : Gives random meme(s) through /gimme endpoint
type GimmeController struct{}

// GetOneRandomMeme : Returns a single meme from a random subreddit
func (g GimmeController) GetOneRandomMeme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Test": "Hello World"})
	reddit.GetMemes()
	return
}
