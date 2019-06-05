package sdk

import (
	"errors"
	"github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Bnb struct {
	keyManager keys.KeyManager
	dexClient  client.DexClient
	rpcClient  rpc.Client
}

func Initialize(dex, rp string, net types.ChainNetwork) (*Bnb, error) {
	key, err := keys.NewKeyManager()
	if err != nil {
		return nil, errors.New("new key manager err:" + err.Error())
	}

	c, err := client.NewDexClient(dex, net, key)
	if err != nil {
		return nil, errors.New("new bnb client err:" + err.Error())
	}
	r := rpc.NewRPCClient(rp, net)
	return &Bnb{keyManager: key, dexClient: c, rpcClient: r}, nil
}

func (b *Bnb) GetBlockByNumber(height int64) (*ctypes.ResultBlock, error) {
	block, err := b.rpcClient.Block(&height)
	if err != nil {
		return nil, errors.New("rpc get block err:" + err.Error())
	}
	return block, nil
}

