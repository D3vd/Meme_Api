package models

// Meme : Basic structure of a Meme Response
type Meme struct {
	PostLink  string   `json:"postLink"`
	SubReddit string   `json:"subreddit"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	Author    string   `json:"author"`
	Ups       int      `json:"ups"`
	NSFW      bool     `json:"nsfw"`
	Spoiler   bool     `json:"spoiler"`
	Preview   []string `json:"preview"`
}
