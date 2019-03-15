package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"weibo/app"
	"weibo/code"
	"weibo/entity"
)

func GetWeibo(user, token string) (person entity.Personreq, err error) {
	URL, err := url.Parse("https://api.weibo.com/2/users/show.json")
	if err != nil {
		return
	}
	query := URL.Query()
	if user != "" {
		query.Add("screen_name", user)
	}
	if token != "" {
		query.Add("access_token", token)
	}
	URL.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return
	}
	resp, err := Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		app.Logger.Warn().Str("response", string(b)).Msg("received error response.")
		var result code.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &person)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
