package main

import (
	"context"
	"fmt"
	"github.com/eager7/one_chain/go_ethereum/go_eth"
	"math/big"
)

func main() {
	fmt.Println("get ethereum block test...")
	ctx := context.Background()
	client, err := new(go_eth.Eth).Initialize(ctx, "wss://mainnet.infura.io/ws")
	if err != nil {
		panic(err)
	}
	block, err := client.Client().BlockByNumber(ctx, new(big.Int).SetUint64(700000))
	if err != nil {
		panic(err)
	}
	fmt.Println(block.Hash().Hex(), block.Transactions().Len())
	go client.SubscribeNewHeader()
	select{}
}
