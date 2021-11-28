load("twitter", "client")

def run():
    token = env("BEARER_TOKEN")
    tweets = client(token).Tweets("ryancohen")
    dump(tweets)
    for tweet in tweets:
        printf("Ryan tweeted: %q at %s\n", tweet.Text, tweet.CreatedAt)

run()