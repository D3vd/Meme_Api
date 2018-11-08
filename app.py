from flask import Flask, render_template, jsonify
from reddit_handler import *

app = Flask(__name__)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/gimme')
def one_post():
    re = get_posts('memes', 20)

    r = random.choice(re)

    return jsonify({
        'title': r[0],
        'url': r[1],
        'postLink': r[2]
    })
