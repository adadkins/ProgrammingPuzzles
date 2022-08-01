import praw
import os
from dotenv import load_dotenv

load_dotenv()

subreddits=os.environ["subreddits"]
search_phrase=os.environ["searched_phrase"]
response=os.environ["response"]
username=os.environ["username"]

def main():
    reddit: praw.Reddit = praw.Reddit(
        client_id=os.environ["client_id"],
        client_secret=os.environ["client_secret"],
        user_agent=os.environ["user_agent"],
        username=os.environ["username"],
        password=os.environ["password"]
    )
    while True:
        print("Listening...")
        listen(reddit)
        print("Exiting...")

def listen(reddit):
    for comment in reddit.subreddit("all").stream.comments(skip_existing=True):
        if search_phrase.lower() in comment.body.lower() and comment.author.name.lower()!=username.lower():
            print("found a comment:")
            print(comment.body)
            comment.reply(body=response) 

if __name__ == "__main__":
    main()