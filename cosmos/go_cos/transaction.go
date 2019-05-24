package go_cos

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/rpc/core/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"time"
)

type Transaction struct {
	Tx      auth.StdTx
	Receipt sdk.TxResponse
}

func (c *CosCli) ParseTransaction(block *core_types.ResultBlock) (txs []Transaction, err error) {
	for _, t := range block.Block.Txs {
		if tx, err := queryTx(c.cdc, c.ctx, t.Hash(), block); err != nil {
			return nil, errors.New("ParseTransaction err:" + err.Error())
		} else {
			txs = append(txs, tx)
		}
	}
	return txs, nil
}

func queryTx(cdc *codec.Codec, cliCtx context.CLIContext, hash []byte, resBlock *ctypes.ResultBlock) (Transaction, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return Transaction{}, err
	}
	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return Transaction{}, err
	}
	out, tx, err := formatTxResult(cdc, resTx, resBlock)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{Tx: tx, Receipt: out}, nil
}

func formatTxResult(cdc *codec.Codec, resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (sdk.TxResponse, auth.StdTx, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return sdk.TxResponse{}, auth.StdTx{}, err
	}
	return sdk.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), tx, nil
}

func parseTx(cdc *codec.Codec, txBytes []byte) (auth.StdTx, error) {
	var tx auth.StdTx
	err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
	if err != nil {
		return auth.StdTx{}, err
	}
	return tx, nil
}
