package go_eth

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"regexp"
)

func (e *Eth) IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func (e *Eth) IsContractAddress(address string, number *big.Int) (bool, error) {
	code, err := e.client.CodeAt(e.ctx, common.HexToAddress(address), number)
	if err != nil {
		return false, errors.New("get address' code err:" + err.Error())
	}
	return len(code) > 0, nil
}
