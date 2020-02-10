package reddit

import "strconv"

// GetSubredditURL : Returns API Reddit URL with Limit
func GetSubredditURL(subreddit string, limit int) (url string) {

	url = "https://api.reddit.com/r/" + subreddit + "/hot?limit=" + strconv.Itoa(limit)

	return
}
