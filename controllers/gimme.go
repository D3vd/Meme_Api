package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GimmeController : Gives random meme(s) through /gimme endpoint
type GimmeController struct{}

// GetOneRandomMeme : Returns a single meme from a random subreddit
func (g GimmeController) GetOneRandomMeme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Test": "Hello World"})
	return
}
