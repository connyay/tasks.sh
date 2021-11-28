package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

var Module starlark.StringDict

func init() {
	m, err := convert.MakeStringDict(map[string]interface{}{
		"client": Client,
	})
	if err != nil {
		panic("converting twilio module")
	}
	Module = m
}

func Client(
	accountSID,
	authToken,
	numberFrom string,
) (rc twilioClient, err error) {
	return twilioClient{accountSID, authToken, numberFrom}, err
}

type twilioClient struct {
	accountSID, authToken, numberFrom string
}

func (tc twilioClient) SendSMS(number, body string) (string, error) {
	if number == "" {
		return "", errors.New("number is required")
	}
	if body == "" {
		return "", errors.New("body is required")
	}
	parameters := url.Values{}
	parameters.Set("To", number)
	parameters.Set("From", tc.numberFrom)
	parameters.Set("Body", body)

	url := "https://api.twilio.com/2010-04-01/Accounts/" + tc.accountSID + "/Messages.json"
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(parameters.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(tc.accountSID, tc.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("incorrect status code %d %s", res.StatusCode, res.Status)
	}
	defer res.Body.Close()
	var data struct {
		SID string `json:"sid"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	return data.SID, err
}
