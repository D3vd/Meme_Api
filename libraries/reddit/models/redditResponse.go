package models

import (
	"strings"
)

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
	Preview           Preview     `json:"preview"`
	Spoiler           bool        `json:"spoiler"`
	Locked            bool        `json:"locked"`
	ID                string      `json:"id"`
	Author            string      `json:"author"`
	URL               string      `json:"url"`
	CreatedUtc        float64     `json:"created_utc"`
	Media             interface{} `json:"media"`
	IsVideo           bool        `json:"is_video"`
}

// Preview :
type Preview struct {
	Images  []PreviewImages `json:"images"`
	Enabled bool            `json:"enabled"`
}

// PreviewImages :
type PreviewImages struct {
	Source      PreviewImage   `json:"source"`
	Resolutions []PreviewImage `json:"resolutions"`
	Variants    interface{}    `json:"variants"`
	ID          string         `json:"id"`
}

// PreviewImage :
type PreviewImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// GetCleanPreviewImages : Removes `amp;` from preview image url
func (d Data) GetCleanPreviewImages() (urls []string) {
	images := d.Preview.Images
	var links []string

	if len(images) != 0 && len(images[0].Resolutions) != 0 {
		for _, r := range images[0].Resolutions {
			links = append(links, r.URL)
		}
	} else {
		links = append(links, d.URL)
	}

	for _, l := range links {
		urls = append(urls, strings.ReplaceAll(l, "amp;", ""))
	}

	return
}

// GetShortLink : Get the short URL for the post
func (d Data) GetShortLink() (url string) {
	return "https://redd.it/" + d.ID
}
