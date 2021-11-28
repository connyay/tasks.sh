load("twitter", "client")
load("environment", "TWITTER_BEARER_TOKEN")

def run():
    tweets = client(TWITTER_BEARER_TOKEN).Tweets(parameters.get("user", "ryancohen"))
    if parameters.get("dump"):
        dump(tweets)
    for tweet in tweets:
        printf("Ryan tweeted: %q at %s\n", tweet.Text, tweet.CreatedAt)

run()