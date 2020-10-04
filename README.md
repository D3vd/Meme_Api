# Meme API

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/990c1b4d07e74fe48ac6cefb9238f601)](https://app.codacy.com/gh/D3vd/Meme_Api?utm_source=github.com&utm_medium=referral&utm_content=D3vd/Meme_Api&utm_campaign=Badge_Grade)
[![CodeFactor](https://www.codefactor.io/repository/github/r3l3ntl3ss/meme_api/badge)](https://www.codefactor.io/repository/github/r3l3ntl3ss/meme_api)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/97620b169e9c49cba3df99699afbad26)](https://app.codacy.com/manual/dev.91096/Meme_Api?utm_source=github.com&utm_medium=referral&utm_content=R3l3ntl3ss/Meme_Api&utm_campaign=Badge_Grade_Dashboard)

JSON API for a random meme scraped from reddit.

API Link : [https://meme-api.herokuapp.com/gimme](https://meme-api.herokuapp.com/gimme)

**Example Response:**

```json
{
  "postLink": "https://redd.it/9vqgv2",
  "subreddit": "memes",
  "title": "Good mor-ning Reddit!...",
  "url": "https://i.redd.it/yykt3r9zsex11.png",
  "nsfw": false,
  "spoiler": false
}
```

## Custom Endpoints

### Specify count (MAX 50)

In order to get multiple memes in a single request specify the count with the following endpoint.

Endpoint: [/gimme/{count}](https://meme-api.herokuapp.com/gimme/2)

Example: [https://meme-api.herokuapp.com/gimme/2](https://meme-api.herokuapp.com/gimme/2)

Response:

```json
{
  "count": 2,
  "memes": [
    {
      "postLink": "https://redd.it/d5bn24",
      "subreddit": "me_irl",
      "title": "me_irl",
      "url": "https://i.redd.it/6wjb8gibu2n31.jpg",
      "nsfw": false,
      "spoiler": false
    },
    {
      "postLink": "https://redd.it/d4zipy",
      "subreddit": "me_irl",
      "title": "Me🚐_irl",
      "url": "https://i.redd.it/hyl6fgweswm31.jpg",
      "nsfw": false,
      "spoiler": false
    }
  ]
}
```

### Specify Subreddit

By default the API grabs a random meme from '_memes_', '_dankmemes_', '_me_irl_' subreddits. To provide your own custom subreddit use the following endpoint.

Endpoint: [/gimme/{subreddit}](https://meme-api.herokuapp.com/gimme/dankmemes)

Example: [https://meme-api.herokuapp.com/gimme/dankmemes](https://meme-api.herokuapp.com/gimme/dankmemes)

### Specify Subreddit Count (MAX 50)

In order to get a custom number of memes from a specific subreddit provide the name of the subreddit and the count in the following endpoint.

Endpoint: [/gimme/{subreddit}/{count}](https://meme-api.herokuapp.com/gimme/dankmemes/2)

Example: [https://meme-api.herokuapp.com/gimme/dankmemes/2](https://meme-api.herokuapp.com/gimme/dankmemes/2)

Response:

```json
{
  "count": 2,
  "memes": [
    {
      "postLink": "https://redd.it/d5e119",
      "title": "Mph and km/h so everyone can understand",
      "url": "https://i.redd.it/imb4r0se74n31.jpg",
      "nsfw": false,
      "spoiler": false
    },
    {
      "postLink": "https://redd.it/d5e5ns",
      "title": "Funnier as a team.",
      "url": "https://i.redd.it/ixb7absfa4n31.jpg",
      "nsfw": false,
      "spoiler": false
    }
  ],
  "subreddit": "dankmemes"
}
```
