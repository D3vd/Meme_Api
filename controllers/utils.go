package controllers

import (
	"github.com/R3l3ntl3ss/Meme_Api/models"
	"math/rand"
	"time"
)

func GetRandomN(n int) (num int) {
	rand.Seed(time.Now().Unix())
	num = rand.Int() % n

	return
}

func GetNRandomMemes(memes []models.Meme, n int) []models.Meme {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(memes), func(i, j int) { memes[i], memes[j] = memes[j], memes[i] })

	return memes[:n]
}
