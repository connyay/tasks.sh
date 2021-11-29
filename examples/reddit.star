load("reddit", "client")
load("environment", "REDDIT_BOT_CLIENT_ID", "REDDIT_BOT_CLIENT_SECRET", "REDDIT_BOT_USERNAME", "REDDIT_BOT_PASSWORD")

def timestamp(ts_utc):
    format = "Mon, 02 Jan 2006 15:04:05 MST" # RFC1123
    return time.from_timestamp(int(ts_utc)).format(format)

def main(args):
    subreddit = args.get("subreddit", "earthporn")
    sort = args.get("sort", "rising")
    count = int(args.get("count", "5"))
    posts = client(
        REDDIT_BOT_CLIENT_ID,
        REDDIT_BOT_CLIENT_SECRET,
        REDDIT_BOT_USERNAME,
        REDDIT_BOT_PASSWORD,
    ).Posts(subreddit, sort)[0:count]
    if args.get("dump"):
        dump(posts)
    for post in posts:
        printf("[%s]: %q (/u/%s at %s)\n", subreddit, post.Data.Title, post.Data.Author, timestamp(post.Data.CreatedUtc))