package response

// OneMemeResponse : Response for a single Meme
type OneMemeResponse struct {
	PostLink  string `json:"postLink"`
	Subreddit string `json:"subreddit"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
