package main

import (
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/eager7/elog"
	"github.com/eager7/one_chain/bnb/sdk"
)

var log = elog.NewLogger("bnb", elog.DebugLevel)

func main() {
	bnb, err := sdk.Initialize("testnet-dex.binance.org", "tcp://39.104.144.212:27147", types.ProdNetwork)
	if err != nil {
		panic(err)
	}
	number, err := bnb.CurrentNumber()
	if err != nil {
		panic(err)
	}
	for i := number - 100; i < number; i++ {
		log.Notice("request block:", i)
		block, err := bnb.GetBlockByNumber(int64(i))
		if err != nil {
			panic(err)
		}
		if err := bnb.ParseTransactions(block); err != nil {
			panic(err)
		}
	}
}
