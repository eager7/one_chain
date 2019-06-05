package tokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type CoinPrice struct {
	USD float64 `json:"usd"`
}

func CoinsPrice(u string, coins []string) (list map[string]CoinPrice, err error) {
	list = make(map[string]CoinPrice)
	lists := SplitSliceLen(coins, 100)
	for _, l := range lists {
		if err := CoinsPriceShard(u, l, list); err != nil {
			fmt.Println("CoinsPriceShard err:", err.Error())
		}
		fmt.Println("CoinsPrice:", len(list))
	}
	return list, nil
}

func SplitSliceLen(slice []string, length int) [][]string {
	var sTemp [][]string
	n := len(slice) / length
	for i := 0; i < n; i++ {
		ll := slice[i*length : (i+1)*length]
		sTemp = append(sTemp, ll)
	}
	sTemp = append(sTemp, slice[length*n:])
	return sTemp
}

func CoinsPriceShard(u string, coins []string, list map[string]CoinPrice) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return errors.New("http new request err:" + err.Error())
	}

	q := url.Values{}
	q.Add("ids", strings.Join(coins, ","))
	q.Add("vs_currencies", "usd")
	req.Header.Set("accept", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return errors.New("client do err:" + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("request error:" + resp.Status)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("read body err:" + err.Error())
	}
	if err := json.Unmarshal(respBody, &list); err != nil {
		return errors.New("coin brief json unmarshal err:" + err.Error())
	}
	return nil
}
