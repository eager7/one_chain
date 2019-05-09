package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eager7/one_chain/cosmos/go_cos"
)

func main() {
	fmt.Println("start go cosmos program...")

	cli := go_cos.Initialize("/tmp/config.toml")
	block, err := cli.GetBlockByNumber(207503)
	if err != nil {
		panic(err)
	}
	fmt.Println(cli.JsonString(block))
	txs := cli.ParseTransaction(block)
	for _, tx := range txs {
		fmt.Println("hash:", hex.EncodeToString(tx.Hash()))
	}
}
