package redis

import (
	"Meme_Api/data"
	"Meme_Api/models"
	"encoding/json"
)

// WritePostsToCache : Takes a List of Memes and writes it to Cache
func WritePostsToCache(sub string, memes []models.Meme) (ok bool) {
	memesBinary, err := json.Marshal(memes)

	if err != nil {
		return false
	}

	err = Client.Set(sub, memesBinary, data.CacheExpirationTime).Err()

	if err != nil {
		return false
	}

	return true
}
