package gimme

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/libraries/reddit"
)

// Controller : Gives random meme(s) through /gimme endpoint
type Controller struct {
	R *reddit.Reddit
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
