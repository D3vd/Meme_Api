package server

import (
	"github.com/R3l3ntl3ss/Meme_Api/controllers"
	"github.com/gin-gonic/gin"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gimme := router.Group("gimme")
	{
		newMeme := new(controllers.GimmeController)

		gimme.GET("/", newMeme.GetOneRandomMeme)
	}

	return router
}
