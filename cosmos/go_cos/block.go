package go_cos

import (
	"fmt"
	"github.com/tendermint/tendermint/rpc/core/types"
)

func (c *CosCli) GetBlockByNumber(height int64) (*core_types.ResultBlock, error) {
	node, err := c.ctx.GetNode()
	if err != nil {
		return nil, err
	}
	return node.Block(&height)
}

func (c *CosCli) JsonString(res interface{}) string {
	data, err := c.cdc.MarshalJSON(res)
	if err != nil {
		fmt.Println("marsh json err:", err)
		return ""
	}
	return string(data)
}
