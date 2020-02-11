package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/R3l3ntl3ss/Meme_Api/controllers/reddit"
	"github.com/R3l3ntl3ss/Meme_Api/data"
	"github.com/R3l3ntl3ss/Meme_Api/models/response"
)

// GimmeController : Gives random meme(s) through /gimme endpoint
type GimmeController struct {
	R reddit.Reddit
}

// GetOneRandomMeme : Returns a single meme from a random subreddit
func (g GimmeController) GetOneRandomMeme(c *gin.Context) {

	// Choose Random Meme Subreddit
	sub := data.MemeSubreddits[GetRandomN(len(data.MemeSubreddits))]

	// Get 50 posts from that Subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Choose one post from the list
	meme := memes[GetRandomN(len(memes))]

	response := response.OneMemeResponse{
		PostLink:  meme.PostLink,
		Subreddit: meme.SubReddit,
		Title:     meme.Title,
		URL:       meme.URL,
	}

	c.JSON(http.StatusOK, response)
	return
}

// GetOnePostFromSub : Get one post from a specific subreddit
func (g GimmeController) GetOnePostFromSub(c *gin.Context) {

	sub := c.Param("interface")

	// Get 50 posts from that Subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Choose one post from the list
	meme := memes[GetRandomN(len(memes))]

	// TODO: Create custom response and error handling

	response := response.OneMemeResponse{
		PostLink:  meme.PostLink,
		Subreddit: meme.SubReddit,
		Title:     meme.Title,
		URL:       meme.URL,
	}
	c.JSON(http.StatusOK, response)
	return
}

// GetNRandomMemes : Returns N no. of memes from a random subreddit
func (g GimmeController) GetNRandomMemes(c *gin.Context) {

	count, _ := strconv.Atoi(c.Param("interface"))

	// Check if the count is less than 50
	if count > 50 {
		count = 50
	}

	// Choose a random meme subreddit
	sub := data.MemeSubreddits[GetRandomN(len(data.MemeSubreddits))]

	// Get 50 posts from that subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Get N no. of posts from that list
	memes = GetNRandomMemes(memes, count)

	var memesResponse []response.OneMemeResponse

	for _, meme := range memes {
		memeResponse := response.OneMemeResponse{
			PostLink:  meme.PostLink,
			Subreddit: meme.SubReddit,
			Title:     meme.Title,
			URL:       meme.URL,
		}

		memesResponse = append(memesResponse, memeResponse)
	}

	response := response.MultipleMemesResponse{
		Count: len(memesResponse),
		Memes: memesResponse,
	}

	c.JSON(http.StatusOK, response)
	return
}

// GetNPostsFromSub : Get N no. of posts from a specific subreddit
func (g GimmeController) GetNPostsFromSub(c *gin.Context) {

	sub := c.Param("interface")
	count, err := strconv.Atoi(c.Param("count"))

	if err != nil {
		// TODO: Handle Error properly
		c.JSON(http.StatusBadRequest, "Give proper count value")
		return
	}

	// Check if the count is less than 50
	if count > 50 {
		count = 50
	}

	// Get 50 posts from that subreddit
	memes := g.R.GetNPosts(sub, 50)

	// Get N no. of posts from that list
	memes = GetNRandomMemes(memes, count)

	var memesResponse []response.OneMemeResponse

	for _, meme := range memes {
		memeResponse := response.OneMemeResponse{
			PostLink:  meme.PostLink,
			Subreddit: meme.SubReddit,
			Title:     meme.Title,
			URL:       meme.URL,
		}

		memesResponse = append(memesResponse, memeResponse)
	}

	response := response.MultipleMemesResponse{
		Count: len(memesResponse),
		Memes: memesResponse,
	}

	c.JSON(http.StatusOK, response)
	return
}

// SubredditOrCount : Find if the route is /<subreddit>/ or /<count>/
func (g GimmeController) SubredditOrCount(c *gin.Context) {
	route := c.Param("interface")

	if _, err := strconv.Atoi(route); err == nil {
		g.GetNRandomMemes(c)
	} else {
		g.GetOnePostFromSub(c)
	}

	return
}
