package server

import (
	"github.com/R3l3ntl3ss/Meme_Api/controllers"
	"github.com/R3l3ntl3ss/Meme_Api/libraries/reddit"
	"github.com/gin-gonic/gin"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create a Reddit Object and Initialize it
	r := &reddit.Reddit{}
	r.Init()

	gimme := router.Group("gimme")
	{
		g := controllers.GimmeController{
			R: r,
		}

		gimme.GET("/", g.GetOneRandomMeme)

		gimme.GET("/:interface/", g.SubredditOrCount)

		gimme.GET("/:interface/:count/", g.GetNPostsFromSub)
	}

	return router
}
