package main

import (
	"fmt"
	"github.com/eager7/one_chain/cosmos/go_cos"
)

func main() {
	fmt.Println("start go cosmos program...")

	cli := go_cos.Initialize("/tmp/config.toml")
	go CosMosBlockThread(cli)
	select {}
}

func CosMosBlockThread(cli *go_cos.CosCli) {
	var height int64 = 12905
	for {
		block, err := cli.GetBlockByNumber(height)
		if err != nil {
			panic(err)
		}
		txs, err := cli.ParseTransaction(block)
		if err != nil {
			panic(err)
		}
		if err := cli.HandleTransactions(txs); err != nil {
			panic(err)
		}
		height++
	}
}
