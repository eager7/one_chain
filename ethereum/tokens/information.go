package tokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eager7/one_chain/ethereum/database/tables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CoinInfo struct {
	Id                  string  `json:"id"`
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	BlockTimeInMinutes  float64 `json:"block_time_in_minutes"`
	CountryOrigin       string  `json:"country_origin"`
	GenesisDate         string  `json:"genesis_date"`
	ContractAddress     string  `json:"contract_address"`
	MarketCapRank       float64 `json:"market_cap_rank"`
	CoingeckoRank       float64 `json:"coingecko_rank"`
	CoingeckoScore      float64 `json:"coingecko_score"`
	DeveloperScore      float64 `json:"developer_score"`
	CommunityScore      float64 `json:"community_score"`
	LiquidityScore      float64 `json:"liquidity_score"`
	PublicInterestScore float64 `json:"public_interest_score"`
	LastUpdated         string  `json:"last_updated"`
	Image               struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	Links struct {
		BlockChainSite []string `json:"blockchain_site"`
	} `json:"links"`
	MarketData struct {
		CurrentPrice struct {
			USD float64 `json:"usd"`
		} `json:"current_price"`
		MarketCap struct {
			USD float64 `json:"usd"`
		} `json:"market_cap"`
		MarketCapRank float64 `json:"market_cap_rank"`
		TotalVolume   struct {
			USD float64 `json:"usd"`
		} `json:"total_volume"`
		High24h struct {
			USD float64 `json:"usd"`
		} `json:"high_24h"`
		Low24h struct {
			USD float64 `json:"usd"`
		} `json:"low_24h"`
		TotalSupply        float64 `json:"total_supply"`
		CirculatingSupply  float64 `json:"circulating_supply"`
		LastUpdated        string  `json:"last_updated"`
		PriceChange24h     float64 `json:"price_change_24h"`
		MarketCapChange24h float64 `json:"market_cap_change_24h"`
	} `json:"market_data"`
}

func CoinInformation(u string, id string) (info *CoinInfo, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(u, id), nil)
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
	info = new(CoinInfo)
	if err := json.Unmarshal(respBody, info); err != nil {
		fmt.Println("CoinInformation:", id, string(respBody))
		return nil, errors.New("coin brief json unmarshal err:" + err.Error())
	}
	if info.ContractAddress == "" && len(info.Links.BlockChainSite) > 0 {
		for _, link := range info.Links.BlockChainSite {
			if strings.Contains(link, "https://etherscan.io/token/") {
				info.ContractAddress = strings.Replace(link, "https://etherscan.io/token/", "", -1)
			}
		}
	}
	return info, nil
}


func CoinInfoShard(db *gorm.DB, coin CoinBrief) {
retry:
	info, err := CoinInformation(UrlCoinsInfo, coin.Id)
	if err != nil {
		fmt.Println("CoinInformation err:", err)
		time.Sleep(time.Second * 1)
		goto retry
	}
	//fmt.Println(info.Symbol, info.ContractAddress, info.MarketData.CurrentPrice.USD)
	if info.ContractAddress == "" {
		return
	}
	tCoin := tables.TableCoinInfo{
		NameId:      coin.Id,
		Name:        coin.Name,
		Symbol:      coin.Symbol,
		Decimals:    0,
		Contract:    common.HexToAddress(info.ContractAddress).Hex(),
		Transaction: "",
		Supply:      fmt.Sprintf("%f", info.MarketData.TotalSupply),
		Price:       info.MarketData.CurrentPrice.USD,
		ICon:        info.Image.Large,
	}
	if err := db.LogMode(false).Exec(tCoin.SqlCommand()).Error; err != nil {
		fmt.Println("sql raw error:", err, tCoin.SqlCommand())
		goto retry
	}
	tokens.lock.Lock()
	tokens.Tokens[strings.ToLower(info.ContractAddress)] = &Token{
		Name:    info.Name,
		Symbol:  info.Symbol,
		Address: strings.ToLower(info.ContractAddress),
		Price:   fmt.Sprintf("%f", info.MarketData.CurrentPrice.USD),
	}
	tokens.lock.Unlock()
}
