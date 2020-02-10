package models

// Meme : Basic structure of a Meme Response
type Meme struct {
	PostLink  string `json:"postLink"`
	SubReddit string `json:"subreddit"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
