package go_cos

import (
	"fmt"
	"testing"
)

func TestGetBlock(t *testing.T) {
	cli := Initialize("/tmp/cosmos.toml")
	block, err := cli.GetBlockByNumber(384416)
	if err != nil {
		t.Fatal(err)
	}
	txs, err := cli.ParseTransaction(block)
	if err != nil {
		t.Fatal(err)
	}
	for _, tx := range txs {
		for _, m := range tx.Tx.Msgs {
			fmt.Println("tx msg type:", m.Type())
		}
	}
}
