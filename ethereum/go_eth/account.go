package go_eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (e *Eth) BalanceAt(address string, number *big.Int) (*big.Int, error) {
	return e.client.BalanceAt(e.ctx, common.HexToAddress(address), number)
}

func (e *Eth) PendingBalanceAt(address string) (*big.Int, error) {
	return e.client.PendingBalanceAt(e.ctx, common.HexToAddress(address))
}

func (e *Eth) PendingNonceAt(address string) (uint64, error) {
	return e.client.PendingNonceAt(e.ctx, common.HexToAddress(address))
}
