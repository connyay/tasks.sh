load("reddit", "client", "SORT_NEW")

def run():
    posts = client(
        env("REDDIT_BOT_CLIENT_ID"),
        env("REDDIT_BOT_CLIENT_SECRET"),
        env("REDDIT_BOT_USERNAME"),
        env("REDDIT_BOT_PASSWORD"),
    ).Posts("earthporn", SORT_NEW)
    # dump(posts)
    for post in posts:
        printf("[earthporn]: %q (/u/%s at %f)\n", post.Data.Title, post.Data.Author, post.Data.CreatedUtc)

run()