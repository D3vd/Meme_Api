package response

// MultipleMemesResponse : Response for multiple Memes
type MultipleMemesResponse struct {
	Count int               `json:"count"`
	Memes []OneMemeResponse `json:"memes"`
}
