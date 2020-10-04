package redis

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

// Redis : Redis structure with client
var (
	RedisURL string
	Client   *redis.Client
)

// Init : Initialize New Redis Connection
func Init() {
	redisURL := os.Getenv("REDISCLOUD_URL")

	if redisURL == "" {
		log.Fatalln("Redis Cloud URL has been set")
		return
	}

	options, err := redis.ParseURL(redisURL)

	if err != nil {
		log.Fatalln("Invalid Redis Cloud URL. Unable to Parse")
	}

	redisClient := redis.NewClient(options)

	if redisClient.Ping().String() != "ping: PONG" {
		log.Fatalln("Error while contacting Redis DB. Please check the Redis Cloud URL.")
	}

	Client = redisClient
}
