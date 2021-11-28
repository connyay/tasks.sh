load("twitter", "client")

def run():
    token = env("TWITTER_BEARER_TOKEN")
    tweets = client(token).Tweets(parameters.get("user", "ryancohen"))
    if parameters.get("dump"):
        dump(tweets)
    for tweet in tweets:
        printf("Ryan tweeted: %q at %s\n", tweet.Text, tweet.CreatedAt)

run()