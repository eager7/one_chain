package go_eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (e *Eth) HeaderByNumber(number *big.Int) (*types.Header, error) {
	return e.client.HeaderByNumber(e.ctx, number)
}

type Transaction struct {
	*types.Transaction
	Time    uint64
	From    string
	GasUsed uint64
	Index   uint16
}

func (e *Eth) BlockByNumber(number *big.Int) (*types.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, number)
	if err != nil {
		if err == context.Canceled {
			return nil, err
		}
		return nil, errors.New("get block by number err:" + err.Error())
	}
	return block, nil
}

func (e *Eth) BlockByHash(hash common.Hash) (*types.Block, error) {
	block, err := e.client.BlockByHash(e.ctx, hash)
	if err != nil {
		return nil, errors.New("get block by number err:" + err.Error())
	}
	return block, nil
}

func (e *Eth) AnalysisTransactions(block *types.Block) (map[common.Hash]*Transaction, error) {
	txs := make(map[common.Hash]*Transaction)
	for index, tx := range block.Transactions() {
		t := Transaction{
			Transaction: block.Transaction(tx.Hash()),
			Time:        block.Time(),
			Index:       uint16(index),
		}
		addr, err := e.client.TransactionSender(e.ctx, tx, block.Hash(), uint(index))
		if err != nil {
			return nil, err
		}
		t.From = addr.Hex()
		txs[tx.Hash()] = &t
	}
	return txs, nil
}

func (e *Eth) TransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return e.client.TransactionReceipt(e.ctx, hash)
}

func (e *Eth) AnalysisReceipts(block *types.Block) (map[common.Hash]*types.Receipt, error) {
	receipts := make(map[common.Hash]*types.Receipt, 1)
	for _, tx := range block.Transactions() {
		receipt, err := e.TransactionReceipt(tx.Hash())
		if err != nil {
			return nil, errors.New("analysis receipts err:" + err.Error())
		}
		receipts[tx.Hash()] = receipt
	}
	return receipts, nil
}
