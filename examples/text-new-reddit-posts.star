load("reddit", reddit_client="client", "SORT_NEW")
load("twilio", twilio_client="client")

def get_posts(subreddit):
    posts = reddit_client(
        env("REDDIT_BOT_CLIENT_ID"),
        env("REDDIT_BOT_CLIENT_SECRET"),
        env("REDDIT_BOT_USERNAME"),
        env("REDDIT_BOT_PASSWORD"),
    ).Posts(subreddit, SORT_NEW)
    if parameters.get("dump"):
        dump(posts)
    return posts

def send_msg(number, msg):
    accountSID, authToken, numberFrom = env("TWILIO_ACCOUNT_SID"), env("TWILIO_AUTH_TOKEN"), env("TWILIO_NUMBER_FROM")
    id = twilio_client(accountSID, authToken, numberFrom).SendSMS(number, msg)
    printf("Message sent - ID: %s", id)

def run():
    subreddit = parameters.get("subreddit", "earthporn")
    msg = sprintf("Three latest %s posts:\n", subreddit)
    for post in get_posts(subreddit)[0:3]:
        msg += sprintf("%s (/u/%s @ %v\n", post.Data.Title, post.Data.Author, post.Data.CreatedUtc)

    logf(msg)
    if parameters.get("send"):
        number = parameters.get("number", "")
        send_msg(number, msg)

run()