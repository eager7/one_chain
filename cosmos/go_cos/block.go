package go_cos

import (
	"fmt"
	"github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

func (c *CosCli) GetBlockByNumber(height int64) (*core_types.ResultBlock, error) {
	node, err := c.cos.GetNode()
	if err != nil {
		return nil, err
	}
	return node.Block(&height)
}

func (c *CosCli) ParseTransaction(block *core_types.ResultBlock) []types.Tx {
	return block.Block.Data.Txs
}

func (c *CosCli) JsonString(res interface{}) string {
	data, err := c.cdc.MarshalJSON(res)
	if err != nil {
		fmt.Println("marsh json err:", err)
		return ""
	}
	return string(data)
}
