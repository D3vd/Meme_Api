package reddit

import (
	"log"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit/auth"
)

func GetMemes(subreddit string) {

	accessToken := auth.GetAccessToken()
	url := GetSubredditURL(subreddit, 1)

	body := MakeGetRequest(url, accessToken)

	log.Println(string(body))

}
