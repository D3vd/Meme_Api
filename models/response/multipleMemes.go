package response

// MultipleMemes : Response for multiple Memes
type MultipleMemes struct {
	Count int       `json:"count"`
	Memes []OneMeme `json:"memes"`
}
