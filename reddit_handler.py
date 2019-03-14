import praw
import random

from prawcore.exceptions import Redirect
from prawcore.exceptions import ResponseException


class ClientInfo:
    id = 'TpEjWEc2SHTafQ'
    secret = 'kuojXoBkMZs_JWrCVEaDWgW8FhA'
    user_agent = 'Meme Api'


def is_img_link(link):
    ext = link[-4:]
    if ext == '.jpg' or ext == '.png':
        return True
    else:
        return False


def get_posts(sub, limit):

    r = praw.Reddit(client_id=ClientInfo.id, client_secret=ClientInfo.secret,
                    user_agent=ClientInfo.user_agent)

    submissions = r.subreddit(sub).hot(limit=limit)

    result = [{"title": submission.title,
               "url": submission.url,
               "link": submission.shortlink}
              for submission in submissions]

    return result
