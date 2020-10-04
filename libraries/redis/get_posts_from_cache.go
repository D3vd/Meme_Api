package redis

import (
	"encoding/json"

	"Meme_Api/models"
)

// GetPostsFromCache : Get memes from Cache based on sub
func GetPostsFromCache(sub string) ([]models.Meme, error) {

	memesBinary, err := Client.Get(sub).Bytes()

	if err != nil {
		return nil, err
	}

	var memes []models.Meme

	err = json.Unmarshal(memesBinary, &memes)

	if err != nil {
		return nil, err
	}

	return memes, nil
}
