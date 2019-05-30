package go_cos

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
	"testing"
)

func TestGetBlock(t *testing.T) {
	cli := Initialize("/tmp/cosmos.toml")
	block, err := cli.GetBlockByNumber(384876)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("block hash:", block.Block.Hash().String())
	fmt.Println("miner address:", block.Block.ProposerAddress.String())
	fmt.Println("vote:", block.Block.LastCommit.Precommits[0].ValidatorAddress.String())
	//
	addr, err := sdk.ValAddressFromBech32(`cosmosvaloper15urq2dtp9qce4fyc85m6upwm9xul3049e02707`)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("addr:", addr.String())
	fmt.Println("addr hash:", crypto.AddressHash(addr))
	//fmt.Println("addr hash:", crypto.Address())
	txs, err := cli.ParseTransaction(block)
	if err != nil {
		t.Fatal(err)
	}
	for _, tx := range txs {
		for _, m := range tx.Tx.Msgs {
			fmt.Println("tx msg type:", m.Type())
			switch m.Type() {
			case "send":
				msg, ok := m.(bank.MsgSend)
				if !ok {
					t.Fatal("can't convert")
				}
				fmt.Println(msg.FromAddress, msg.ToAddress)
				fmt.Println(msg.FromAddress.String(), msg.ToAddress.String(), msg.Amount)
			case "begin_unbonding":
				msg, ok := m.(types.MsgUndelegate)
				if !ok {
					t.Fatal("can't convert")
				}
				fmt.Println(msg.DelegatorAddress.String(), msg.ValidatorAddress.String(), msg.Amount)
			}
		}
	}
}
