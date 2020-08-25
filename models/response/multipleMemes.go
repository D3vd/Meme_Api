package response

import "github.com/R3l3ntl3ss/Meme_Api/models"

// MultipleMemes : Response for multiple Memes
type MultipleMemes struct {
	Count int       `json:"count"`
	Memes []models.Meme `json:"memes"`
}
