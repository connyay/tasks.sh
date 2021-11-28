load("reddit", reddit_client="client")
load("twilio", twilio_client="client")
load("environment", "REDDIT_BOT_CLIENT_ID", "REDDIT_BOT_CLIENT_SECRET", "REDDIT_BOT_USERNAME", "REDDIT_BOT_PASSWORD")
load("environment", "TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN", "TWILIO_NUMBER_FROM")

def get_posts(subreddit, sort):
    posts = reddit_client(
        REDDIT_BOT_CLIENT_ID,
        REDDIT_BOT_CLIENT_SECRET,
        REDDIT_BOT_USERNAME,
        REDDIT_BOT_PASSWORD,
    ).Posts(subreddit, sort)
    if parameters.get("dump"):
        dump(posts)
    return posts

def send_msg(number, msg):
    id = twilio_client(
        TWILIO_ACCOUNT_SID,
        TWILIO_AUTH_TOKEN,
        TWILIO_NUMBER_FROM,
    ).SendSMS(number, msg)
    printf("Message sent - ID: %s", id)

def run():
    subreddit = parameters.get("subreddit", "earthporn")
    sort = parameters.get("sort", "new")
    count = int(parameters.get("count", "3"))
    msg = sprintf("%s %s posts:\n", sort, subreddit)
    for post in get_posts(subreddit, sort)[0:count]:
        msg += sprintf("%s (/u/%s @ %v\n", post.Data.Title, post.Data.Author, post.Data.CreatedUtc)

    logf(msg)
    if parameters.get("send"):
        number = parameters.get("number", "")
        send_msg(number, msg)

run()