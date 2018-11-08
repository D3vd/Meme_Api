import praw

from prawcore.exceptions import Redirect
from prawcore.exceptions import ResponseException


class ClientInfo:
    id = 'TpEjWEc2SHTafQ'
    secret = 'kuojXoBkMZs_JWrCVEaDWgW8FhA'
    user_agent = 'Meme Api'


def get_image_links(sub, limit):

    try:
        r = praw.Reddit(client_id=ClientInfo.id, client_secret=ClientInfo.secret,
                        user_agent=ClientInfo.user_agent)

        submissions = r.subreddit(sub).hot(limit=limit)

    except Redirect:
        print("Invalid Subreddit!")
        return 0

    except ResponseException:
        print("Client info is wrong. Check again.")
        return 0

    result = [[submission.title, submission.url, submission.shortlink]
              for submission in submissions]

    return result


if __name__ == '__main__':
    r = get_image_links('memes', 10)

    for i in range(0, 10):
        print('{}. {}\n {}\n {}\n'.format(i+1, r[i][0], r[i][1], r[i][2]))
