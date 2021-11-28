load("reddit", "client", "SORT_NEW")

def run():
    subreddit = parameters.get("subreddit", "earthporn")
    posts = client(
        env("REDDIT_BOT_CLIENT_ID"),
        env("REDDIT_BOT_CLIENT_SECRET"),
        env("REDDIT_BOT_USERNAME"),
        env("REDDIT_BOT_PASSWORD"),
    ).Posts(subreddit, SORT_NEW)
    if parameters.get("dump"):
        dump(posts)
    for post in posts:
        printf("[%s]: %q (/u/%s at %f)\n", subreddit, post.Data.Title, post.Data.Author, post.Data.CreatedUtc)

run()