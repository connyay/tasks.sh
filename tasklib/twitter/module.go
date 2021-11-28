package twitter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

var Module starlark.StringDict

func init() {
	m, err := convert.MakeStringDict(map[string]interface{}{
		"client": Client,
	})
	if err != nil {
		panic("converting twitter module")
	}
	Module = m
}

func Client(token string) twitterClient {
	return twitterClient{token}
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`

	twitterClient
}

func (u User) Tweets() ([]Tweet, error) {
	query := url.Values{}
	query.Add("tweet.fields", "created_at,source")
	url := "https://api.twitter.com/2/users/" + u.ID + "/tweets?" + query.Encode()
	res, err := u.sendGet(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var timelineResponse struct {
		Data []Tweet `json:"data"`
		Meta struct {
			OldestID    string `json:"oldest_id"`
			NewestID    string `json:"newest_id"`
			ResultCount int    `json:"result_count"`
			NextToken   string `json:"next_token"`
		} `json:"meta"`
	}
	err = json.NewDecoder(res.Body).Decode(&timelineResponse)
	if err != nil {
		return nil, err
	}
	return timelineResponse.Data, nil
}

type Tweet struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	Source    string    `json:"source"`
}

type twitterClient struct {
	token string
}

func (tc twitterClient) User(username string) (User, error) {
	query := url.Values{}
	query.Add("usernames", username)
	url := "https://api.twitter.com/2/users/by?" + query.Encode()
	res, err := tc.sendGet(url)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()
	var usersRes struct {
		Data []User `json:"data"`
	}
	err = json.NewDecoder(res.Body).Decode(&usersRes)
	if err != nil {
		return User{}, err
	}
	if len(usersRes.Data) == 0 {
		return User{}, errors.New("user not found")
	}
	u := usersRes.Data[0]
	u.twitterClient = tc
	return u, nil
}

func (tc twitterClient) sendGet(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tc.token)
	return http.DefaultClient.Do(req)
}
