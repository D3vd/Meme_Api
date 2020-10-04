package server

import (
	"Meme_Api/api/gimme"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	router := gin.New()

	// Gin and CORS Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// /gimme routes
	gimmeRouter := router.Group("gimme")
	gimme.Routes(gimmeRouter)

	return router
}
