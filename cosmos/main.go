package main

import (
	"fmt"
	"github.com/eager7/one_chain/cosmos/go_cos"
	"sync/atomic"
)

var height int64 = 12905

func main() {
	fmt.Println("start go cosmos program...")

	cli := go_cos.Initialize("/tmp/config.toml")
	for i := 0; i < 10; i++ {
		go CosMosBlockThread(cli)
	}
	select {}
}

func CosMosBlockThread(cli *go_cos.CosCli) {
	for {
		num := atomic.AddInt64(&height, 1)
		block, err := cli.GetBlockByNumber(num)
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
	}
}
