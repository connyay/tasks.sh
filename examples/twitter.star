load("twitter", "client")
load("environment", "TWITTER_BEARER_TOKEN")

def run():
    twitter = client(TWITTER_BEARER_TOKEN)
    user = twitter.User(parameters.get("user", "Funfacts"))
    count = int(parameters.get("count", "5"))
    tweets = user.Tweets()[:count]
    if parameters.get("dump"):
        dump(tweets)
    for tweet in tweets:
        printf("[%s]: %q at %s\n", user.Name, tweet.Text, tweet.CreatedAt)

run()