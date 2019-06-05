package sdk

import (
	"errors"
	"github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/eager7/elog"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var log = elog.NewLogger("sdk", elog.DebugLevel)

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

func (b *Bnb) CurrentNumber() (int64, error) {
	ret, err := b.rpcClient.Status()
	if err != nil {
		return 0, errors.New("rpc get status err:" + err.Error())
	}
	log.Info("block chain status:", ret.SyncInfo)
	info, err := b.rpcClient.ABCIInfo()
	if err != nil {
		return 0, errors.New("rpc get info err:" + err.Error())
	}
	log.Info("block number:", info.Response.LastBlockHeight)
	return info.Response.LastBlockHeight, nil
}

func (b *Bnb) GetBlockByNumber(height int64) (*ctypes.ResultBlock, error) {
	block, err := b.rpcClient.Block(&height)
	if err != nil {
		return nil, errors.New("rpc get block err:" + err.Error())
	}
	return block, nil
}

func (b *Bnb) ParseTransactions(block *ctypes.ResultBlock) error {
	for _, t := range block.Block.Txs {
		tx, err := b.GetTransaction(t.Hash())
		if err != nil {
			return errors.New("ParseTransactions error:" + err.Error())
		}
		log.Debug(tx.Hash.String(), tx)
	}
	return nil
}

func (b *Bnb) GetTransaction(hash []byte) (*ctypes.ResultTx, error) {
	ret, err := b.rpcClient.Tx(hash, true)
	if err != nil {
		return nil, errors.New("rpc get tx err:" + err.Error())
	}
	return ret, nil
}
