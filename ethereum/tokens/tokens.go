package tokens

import (
	"errors"
	"fmt"
	"github.com/eager7/one_chain/ethereum/database/tables"
	"github.com/jbenet/goprocess"
	"github.com/jbenet/goprocess/periodic"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

const (
	UrlCoinList  = `https://api.coingecko.com/api/v3/coins/list`
	UrlCoinPrice = `https://api.coingecko.com/api/v3/simple/price`
	UrlCoinsInfo = `https://api.coingecko.com/api/v3/coins/%s`
)

type Token struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Price   string `json:"usd_price"`
}

type TokenMap struct {
	lock   sync.RWMutex
	Tokens map[string]*Token
}

var tokens = TokenMap{lock: sync.RWMutex{}, Tokens: make(map[string]*Token)}

func TokenListTask(db *gorm.DB) {
	task := func(proc goprocess.Process) {
		ListAllEthTokens(db)
	}
	process := periodicproc.Tick(time.Hour*24, task)
	process.Go(task)
}

func TokenPriceTask(db *gorm.DB) {
	task := func(proc goprocess.Process) {
		if err := TokenPrice(db); err != nil {
			fmt.Println("token price error:", err)
		}
	}
	process := periodicproc.Tick(time.Minute*60, task)
	process.Go(task)
}

func TokenPrice(db *gorm.DB) error {
	var coins []tables.TableCoinInfo
	if err := db.Order("name desc").Select("`name_id`").Find(&coins).Error; err != nil {
		return errors.New("select coins error:" + err.Error())
	}
	if len(coins) == 0 {
		return nil
	}
	var names []string
	for _, n := range coins {
		names = append(names, n.NameId)
	}
	prices, err := CoinsPrice(UrlCoinPrice, names)
	if err != nil {
		return err
	}
	for nameId, price := range prices {
		if err := db.LogMode(true).Model(&tables.TableCoinInfo{}).Where("name_id = ?", nameId).UpdateColumn("price", price.USD).Error; err != nil {
			fmt.Println("update", nameId, "error:", err)
		}
	}
	return nil
}

func ListAllEthTokens(db *gorm.DB) {
retry:
	list, err := CoinList(UrlCoinList)
	if err != nil {
		fmt.Println("CoinList err:", err)
		goto retry
	}
	fmt.Println(len(list))
	message := make(chan CoinBrief, len(list))
	for _, l := range list {
		message <- l
	}
	close(message)
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case in, ok := <-message:
					if ok {
						CoinInfoShard(db, in)
					} else {
						wg.Done()
						return
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("eth token len:", len(tokens.Tokens))
}
