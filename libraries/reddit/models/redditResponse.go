package reddit

// Response : Main container for the Reddit Response
type Response struct {
	Kind string   `json:"kind"`
	Data MainData `json:"data"`
}

// MainData :
type MainData struct {
	Modhash  string      `json:"modhash"`
	Dist     int         `json:"dist"`
	Children []Children  `json:"children"`
	After    string      `json:"after"`
	Before   interface{} `json:"before"`
}

// Children :
type Children struct {
	Kind string `json:"kind"`
	Data Data   `json:"data,omitempty"`
}

// Data :
type Data struct {
	Subreddit         string      `json:"subreddit"`
	Selftext          string      `json:"selftext"`
	AuthorFullname    string      `json:"author_fullname"`
	Title             string      `json:"title"`
	Downs             int         `json:"downs"`
	Name              string      `json:"name"`
	Ups               int         `json:"ups"`
	IsOriginalContent bool        `json:"is_original_content"`
	Score             int         `json:"score"`
	Over18            bool        `json:"over_18"`
	Preview           interface{} `json:"preview"`
	Spoiler           bool        `json:"spoiler"`
	Locked            bool        `json:"locked"`
	ID                string      `json:"id"`
	Author            string      `json:"author"`
	URL               string      `json:"url"`
	CreatedUtc        float64     `json:"created_utc"`
	Media             interface{} `json:"media"`
	IsVideo           bool        `json:"is_video"`
}

// GetShortLink : Get the short URL for the post
func (d Data) GetShortLink() (url string) {
	return "https://redd.it/" + d.ID
}
