package utils

import (
	"math/rand"
	"time"

	"Meme_Api/models"
)

// GetRandomN : Get a Random Number to a particular range n
func GetRandomN(n int) (num int) {
	rand.Seed(time.Now().Unix())
	num = rand.Int() % n

	return
}

// RemoveNonImagePosts : Remove all posts from Memes List that doesn't end with '.jpg' or '.png'
func RemoveNonImagePosts(memes []models.Meme) []models.Meme {
	var onlyImagePosts []models.Meme

	for _, meme := range memes {
		url := meme.URL
		if url[len(url)-4:] == ".jpg" || url[len(url)-4:] == ".png" || url[len(url)-4:] == ".gif" {
			onlyImagePosts = append(onlyImagePosts, meme)
		}
	}

	return onlyImagePosts
}

// GetNRandomMemes : Get N no. of random memes from a list of Memes
func GetNRandomMemes(memes []models.Meme, n int) []models.Meme {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(memes), func(i, j int) { memes[i], memes[j] = memes[j], memes[i] })

	return memes[:n]
}
