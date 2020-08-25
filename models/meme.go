package models

// Meme : Basic structure of a Meme Response
type Meme struct {
	PostLink  string `json:"postLink"`
	SubReddit string `json:"subreddit"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Author    string `json:"author"`
	Ups       int    `json:"ups"`
	NSFW      bool   `json:"nsfw"`
	Spoiler   bool   `json:"spoiler"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	Author            string      `json:"author"`
	Body              string      `json:"body"`
	Ups               int         `json:"ups"`
}
