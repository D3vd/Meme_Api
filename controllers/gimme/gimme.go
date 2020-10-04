package gimme

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"Meme_Api/libraries/reddit"
	"Meme_Api/libraries/redis"
)

// Controller : Gives random meme(s) through /gimme endpoint
type Controller struct {
	R     *reddit.Reddit
	Cache *redis.Redis
}

// SubredditOrCount : Find if the route is /<subreddit>/ or /<count>/
func (g Controller) SubredditOrCount(c *gin.Context) {
	route := c.Param("interface")

	if _, err := strconv.Atoi(route); err == nil {
		g.GetNRandomMemes(c)
	} else {
		g.GetOnePostFromSub(c)
	}

	return
}
