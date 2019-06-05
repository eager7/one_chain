package tokens

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CoinBrief struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func CoinList(u string) (list []CoinBrief, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, errors.New("http new request err:" + err.Error())
	}

	q := url.Values{}
	//q.Add("limit", "5000")
	req.Header.Set("accept", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("client do err:" + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request error:" + resp.Status)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("read body err:" + err.Error())
	}
	if err := json.Unmarshal(respBody, &list); err != nil {
		return nil, errors.New("coin brief json unmarshal err:" + err.Error())
	}
	return list, nil
}
