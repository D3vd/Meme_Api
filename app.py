from flask import Flask, render_template, jsonify
from reddit_handler import *

app = Flask(__name__)

meme_subreddits = ['memes', 'dankmemes', 'meirl','pewdiepiesubmissions']


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/gimme')
def one_post():
    sub = random.choice(meme_subreddits)
    re = get_posts(sub, 100)

    r = random.choice(re)

    while not is_img_link(r[1]):
        r = random.choice(re)

    return jsonify({
        'title': r[0],
        'url': r[1],
        'postLink': r[2],
        'subreddit': sub
    })


@app.route('/sample')
def sample():
    re = get_posts(random.choice(meme_subreddits), 100)

    r = random.choice(re)

    while not is_img_link(r[1]):
        r = random.choice(re)

    return render_template('sample.html', title=r[0], img_url=r[1], shortlink=r[2])


@app.route('/test')
def test():
    re = get_posts(random.choice(meme_subreddits), 100)

    return render_template('test.html', re=re)


@app.route('/<something>')
def not_found(something):
    return render_template('not_found.html')
