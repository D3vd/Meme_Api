# Meme API

JSON API for a random meme scraped from reddit.

To see a sample check out https://meme-api.herokuapp.com/sample

API Link : https://meme-api.herokuapp.com/gimme

### Example Response:

```json
{
  "postLink": "https://redd.it/9vqgv2",
  "subreddit": "memes",
  "title": "Good mor-ning Reddit!...",
  "url": "https://i.redd.it/yykt3r9zsex11.png"
}
```

## Custom Endpoints

### Specify count (MAX 50)

In order to get multiple memes in a single request specify the count with the following endpoint.

Endpoint: [/gimme/{count}](https://meme-api.herokuapp.com/gimme/2)

Example: https://meme-api.herokuapp.com/gimme/2

Response:

```json
{
  "count": 2,
  "memes": [
    {
      "postLink": "https://redd.it/d5bn24",
      "subreddit": "meirl",
      "title": "meirl",
      "url": "https://i.redd.it/6wjb8gibu2n31.jpg"
    },
    {
      "postLink": "https://redd.it/d4zipy",
      "subreddit": "meirl",
      "title": "Me🚐irl",
      "url": "https://i.redd.it/hyl6fgweswm31.jpg"
    }
  ]
}
```

### Specify Subreddit

By default the API grabs a random meme from '_memes_', '_dankmemes_', '_meirl_' subreddits. To provide your own custom subreddit use the following endpoint.

Endpoint: [/gimme/{subreddit}](https://meme-api.herokuapp.com/gimme/dankmemes)

Example: https://meme-api.herokuapp.com/gimme/dankmemes

### Specify Subreddit Count (MAX 50)

In order to get a custom number of memes from a specific subreddit provide the name of the subreddit and the count in the following endpoint.

Endpoint: [/gimme/{subreddit}/{count}](https://meme-api.herokuapp.com/gimme/dankmemes/2)

Example: https://meme-api.herokuapp.com/gimme/dankmemes/2

Response:

```json
{
  "count": 2,
  "memes": [
    {
      "postLink": "https://redd.it/d5e119",
      "title": "Mph and km/h so everyone can understand",
      "url": "https://i.redd.it/imb4r0se74n31.jpg"
    },
    {
      "postLink": "https://redd.it/d5e5ns",
      "title": "Funnier as a team.",
      "url": "https://i.redd.it/ixb7absfa4n31.jpg"
    }
  ],
  "subreddit": "dankmemes"
}
```
