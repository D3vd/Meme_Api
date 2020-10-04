package gimme

import "github.com/gin-gonic/gin"

// Routes for /gimme
func Routes(group *gin.RouterGroup) {

	group.GET("", GetOneRandomMeme)

	// Checks if /<subreddit> or /<count>
	group.GET("/:interface", SubredditOrCount)

	group.GET("/:interface/:count", GetNPostsFromSub)
}
