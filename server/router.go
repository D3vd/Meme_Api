package server

import (
	"github.com/R3l3ntl3ss/Meme_Api/controllers/gimme"
	"github.com/R3l3ntl3ss/Meme_Api/libraries/reddit"
	"github.com/R3l3ntl3ss/Meme_Api/libraries/redis"
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

	// Create a Redis Object and Initialize it
	cache := &redis.Redis{}
	cache.Init()

	gimmeRouter := router.Group("gimme")
	{
		g := gimme.Controller{
			R:     r,
			Cache: cache,
		}

		gimmeRouter.GET("/", g.GetOneRandomMeme)

		gimmeRouter.GET("/:interface/", g.SubredditOrCount)

		gimmeRouter.GET("/:interface/:count/", g.GetNPostsFromSub)
	}

	return router
}
