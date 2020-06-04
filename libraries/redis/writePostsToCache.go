package redis

import (
	"encoding/json"
	"github.com/R3l3ntl3ss/Meme_Api/data"
	"github.com/R3l3ntl3ss/Meme_Api/models"
)

// WritePostsToCache : Takes a List of Memes and writes it to Cache
func (r Redis) WritePostsToCache(sub string, memes []models.Meme) (ok bool) {
	memesBinary, err := json.Marshal(memes)

	if err != nil {
		return false
	}

	err = r.Client.Set(sub, memesBinary, data.CacheExpirationTime).Err()

	if err != nil {
		return false
	}

	return true
}
