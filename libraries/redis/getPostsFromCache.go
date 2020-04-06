package redis

import (
	"encoding/json"

	"github.com/R3l3ntl3ss/Meme_Api/models"
)

// GetPostsFromCache : Get memes from Cache based on sub
func (r Redis) GetPostsFromCache(sub string) []models.Meme {

	memesBinary, err := r.Client.Get(sub).Bytes()

	if err != nil {
		return nil
	}

	var memes []models.Meme

	err = json.Unmarshal(memesBinary, &memes)

	if err != nil {
		return nil
	}

	return memes
}
