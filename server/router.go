package server

import (
	"Meme_Api/controllers/gimme"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	gimmeRouter := router.Group("gimme")
	{
		gimmeRouter.GET("", gimme.GetOneRandomMeme)

		gimmeRouter.GET("/:interface", gimme.SubredditOrCount)

		gimmeRouter.GET("/:interface/:count", gimme.GetNPostsFromSub)
	}

	return router
}
