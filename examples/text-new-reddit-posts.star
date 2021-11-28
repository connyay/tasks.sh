load("reddit", reddit_client="client")
load("twilio", twilio_client="client")

def get_posts(subreddit, sort):
    posts = reddit_client(
        env("REDDIT_BOT_CLIENT_ID"),
        env("REDDIT_BOT_CLIENT_SECRET"),
        env("REDDIT_BOT_USERNAME"),
        env("REDDIT_BOT_PASSWORD"),
    ).Posts(subreddit, sort)
    if parameters.get("dump"):
        dump(posts)
    return posts

def send_msg(number, msg):
    accountSID, authToken, numberFrom = env("TWILIO_ACCOUNT_SID"), env("TWILIO_AUTH_TOKEN"), env("TWILIO_NUMBER_FROM")
    id = twilio_client(accountSID, authToken, numberFrom).SendSMS(number, msg)
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