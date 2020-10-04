package main

import (
	"Meme_Api/libraries/reddit"
	"Meme_Api/libraries/redis"
	"Meme_Api/server"
)

func main() {

	// Initialize Libraries
	reddit.Init()
	redis.Init()

	// Initialize Server
	server.Init()
}
