package reddit

import (
	"log"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit/auth"
)

func GetMemes() {

	accessToken := auth.GetAccessToken()

	log.Println(accessToken)

}
