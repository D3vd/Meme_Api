from flask import Flask, render_template, jsonify
from reddit_handler import *

app = Flask(__name__)

meme_subreddits = ['memes', 'dankmemes', 'meirl']


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/gimme')
def one_post():
    sub = random.choice(meme_subreddits)
    try:
        re = get_posts(sub, 100)

    except ResponseException:
        return jsonify({
            'status_code': 500,
            'message': 'Internal Server Error'
        })

    r = random.choice(re)

    while not is_img_link(r["url"]):
        r = random.choice(re)

    return jsonify({
        'title': r["title"],
        'url': r["url"],
        'postLink': r["link"],
        'subreddit': sub
    })


@app.route('/gimme/<subreddit>')
def one_post_from_sub(subreddit):
    try:
        re = get_posts(subreddit, 100)

    except Redirect:
        return jsonify({
            'status_code': 404,
            'message': 'Invalid Subreddit'
        })

    except ResponseException:
        return jsonify({
            'status_code': 500,
            'message': 'Internal Server Error'
        })

    r = random.choice(re)

    while not is_img_link(r["url"]):
        r = random.choice(re)

    return jsonify({
        'title': r["title"],
        'url': r["url"],
        'postLink': r["link"],
        'subreddit': subreddit
    })


@app.route('/sample')
def sample():
    re = get_posts(random.choice(meme_subreddits), 100)

    r = random.choice(re)

    while not is_img_link(r["url"]):
        r = random.choice(re)

    return render_template('sample.html', title=r["title"], img_url=r["url"], shortlink=r["link"])


@app.route('/test')
def test():
    re = get_posts(random.choice(meme_subreddits), 100)

    return render_template('test.html', re=re)


@app.route('/<something>')
def not_found(something):
    return render_template('not_found.html')
