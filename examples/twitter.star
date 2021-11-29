load("twitter", "client")
load("environment", "TWITTER_BEARER_TOKEN")

def main(args):
    twitter = client(TWITTER_BEARER_TOKEN)
    user = twitter.User(args.get("user", "Funfacts"))
    count = int(args.get("count", "5"))
    tweets = user.Tweets()[:count]
    if args.get("dump"):
        dump(tweets)
    for tweet in tweets:
        printf("[%s]: %q at %s\n", user.Name, tweet.Text, tweet.CreatedAt)
