package utils

import (
	"math/rand"
	"time"

	"github.com/R3l3ntl3ss/Meme_Api/models"
)

// GetRandomN : Get a Random Number to a particular range n
func GetRandomN(n int) (num int) {
	rand.Seed(time.Now().Unix())
	num = rand.Int() % n

	return
}

// GetNRandomMemes : Get N no. of random memes from a list of Memes
func GetNRandomMemes(memes []models.Meme, n int) []models.Meme {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(memes), func(i, j int) { memes[i], memes[j] = memes[j], memes[i] })

	return memes[:n]
}
