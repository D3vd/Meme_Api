package gimme

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// SubredditOrCount : Find if the route is /<subreddit>/ or /<count>/
func SubredditOrCount(c *gin.Context) {
	route := c.Param("interface")

	if _, err := strconv.Atoi(route); err == nil {
		GetNRandomMemes(c)
	} else {
		GetOnePostFromSub(c)
	}
}
