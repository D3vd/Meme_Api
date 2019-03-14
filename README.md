# Meme API

JSON API for a random meme scraped from reddit.

To see a sample check out https://meme-api.herokuapp.com/sample

API Link : https://meme-api.herokuapp.com/gimme

### Example Response:

```
{
    'postLink': 'https://redd.it/9vqgv2',
    'subreddit': 'memes',
    'title': 'Good mor-ning Reddit!...',
    'url': 'https://i.redd.it/yykt3r9zsex11.png'
}
```



## Custom Endpoints

#### Specify Subreddit: 

By default the API grabs a random meme from '*memes*', '*dankmemes*', '*meirl*', '*pewdiepiesubmissions*' subreddits. To provide your own custom subreddit use the following endpoint.

Endpoint: [/gimme/{subreddit}](https://meme-api.herokuapp.com/gimme/dankmemes)

Example:  https://meme-api.herokuapp.com/gimme/dankmemes

