from flask import Flask, render_template, jsonify
from reddit_handler import *

app = Flask(__name__)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/gimme')
def one_post():
    re = get_posts('memes', 10)

    r = random.choice(re)

    return jsonify({
        'title': r[0],
        'url': r[1],
        'postLink': r[2]
    })


@app.route('/sample')
def sample():
    re = get_posts('memes', 20)

    r = random.choice(re)

    return render_template('sample.html', title=r[0], img_url=r[1], shortlink=r[2])
