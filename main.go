package main

import (
	"log"
	"os"
	"time"

	"Meme_Api/libraries/reddit"
	"Meme_Api/libraries/redis"
	"Meme_Api/server"

	"github.com/getsentry/sentry-go"
)

func main() {

	// Initialize Libraries
	reddit.Init()
	redis.Init()

	// Initialize and defer Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		log.Fatalln("Sentry Init Error: ", err)
	}

	defer sentry.Flush(2 * time.Second)

	// Initialize Server
	server.Init()
}
