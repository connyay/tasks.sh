load("reddit", reddit_client="client")
load("environment", "REDDIT_BOT_CLIENT_ID", "REDDIT_BOT_CLIENT_SECRET", "REDDIT_BOT_USERNAME", "REDDIT_BOT_PASSWORD")
load("database", "db_query", "db_exec", "db_migrate")

def main(args):
    db_migrate([
'''
CREATE TABLE IF NOT EXISTS posts (
    thing_id TEXT PRIMARY KEY,
    title TEXT
)
'''
    ])

    subreddit = args.get("subreddit", "earthporn")
    sort = args.get("sort", "new")
    count = int(args.get("count", "3"))
    msg = sprintf("%s %s posts:\n", sort, subreddit)

    posts = reddit_client(
        REDDIT_BOT_CLIENT_ID,
        REDDIT_BOT_CLIENT_SECRET,
        REDDIT_BOT_USERNAME,
        REDDIT_BOT_PASSWORD,
    ).Posts(subreddit, sort)[0:count]

    for post in posts:
        rows_affected, err = db_exec('INSERT INTO posts("thing_id", "title") VALUES(?, ?)', post.Data.Name, post.Data.Title)
        if rows_affected > 0:
            logf("%s was new", post.Data.Title)
        if err:
            logf("Failed inserting %s %v", post.Data.Title, err)

    if args.get("dump"):
        rows = db_query("select * from posts")
        dump(rows)