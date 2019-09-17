from flask import Flask, render_template, jsonify
from flask_cors import CORS, cross_origin

from reddit_handler import *

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

meme_subreddits = ['memes', 'dankmemes', 'meirl']


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/gimme')
@cross_origin()
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
@cross_origin()
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


@app.route('/gimme/<subreddit>/<int:count>')
@cross_origin()
def multiple_post_from_sub(subreddit, count):

    if count > 100:
        return jsonify({
            'status_code': 400,
            'message': 'Please ensure the count is less than 100'
        })

    try:
        re = get_posts(subreddit, count)

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

    memes = []

    for post in re:
        temp = {
            'title': post["title"],
            'url': post["url"],
            'postLink': post["link"]
        }

        memes.append(temp)

    return jsonify({
        'memes': memes,
        'count': count,
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


@app.errorhandler(404)
@app.route('/<something>')
def not_found(something):
    return render_template('not_found.html')
