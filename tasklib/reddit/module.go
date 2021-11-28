package reddit

import (
	"github.com/sourcegraph/starlight/convert"
	"github.com/thecsw/mira"
	"go.starlark.net/starlark"
)

var Module starlark.StringDict

func init() {
	m, err := convert.MakeStringDict(map[string]interface{}{
		"client": Client,
	})
	if err != nil {
		panic("converting reddit module")
	}
	Module = m
}

func Client(
	clientID,
	clientSecret,
	username,
	password string,
) (rc redditClient, err error) {
	creds := mira.Credentials{
		UserAgent:    "linux:github.com/connyay/tasks.sh:v0.0.0 (by /u/connyay)",
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	}

	reddit, err := mira.Init(creds)
	mira.ReadCredsFromEnv()
	return redditClient{reddit}, err
}

type redditClient struct {
	reddit *mira.Reddit
}

func (rc redditClient) Posts(subreddit string, sort string) ([]mira.PostListingChild, error) {
	return rc.reddit.GetSubredditPosts(subreddit, sort, "all", 10)
}
