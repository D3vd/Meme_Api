package server

import (
	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/controllers"
	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create a Reddit Object and Initialize it
	var r reddit.Reddit
	r.Init()

	gimme := router.Group("gimme")
	{
		meme := controllers.GimmeController{
			R: r,
		}

		gimme.GET("/", meme.GetOneRandomMeme)
	}

	return router
}
