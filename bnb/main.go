package main

import (
	"fmt"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/eager7/one_chain/bnb/sdk"
)

func main() {
	bnb, err := sdk.Initialize("testnet-dex.binance.org", "tcp://127.0.0.1:27147", types.ProdNetwork)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		block, err := bnb.GetBlockByNumber(int64(i))
		if err != nil {
			panic(err)
		}
		fmt.Println(block)
	}
}
