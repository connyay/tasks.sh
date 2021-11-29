load("reddit", reddit_client="client")
load("twilio", twilio_client="client")
load("environment", "REDDIT_BOT_CLIENT_ID", "REDDIT_BOT_CLIENT_SECRET", "REDDIT_BOT_USERNAME", "REDDIT_BOT_PASSWORD")
load("environment", "TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN", "TWILIO_NUMBER_FROM")

def main(args):
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
    if args.get("dump"):
        dump(posts)

    for post in posts:
        msg += sprintf("%s (/u/%s @ %v\n", post.Data.Title, post.Data.Author, post.Data.CreatedUtc)

    logf(msg)
    if args.get("send"):
        number = args.get("number", "")
        id = twilio_client(
            TWILIO_ACCOUNT_SID,
            TWILIO_AUTH_TOKEN,
            TWILIO_NUMBER_FROM,
        ).SendSMS(number, msg)
        printf("Message sent - ID: %s", id)