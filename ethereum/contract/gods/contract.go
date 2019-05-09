package gods

import (
	"errors"
	"github.com/eager7/one_chain/common/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func BalanceAt(address, contract string, client *ethclient.Client) (*big.Int, error) {
	instance, err := NewGods(common.HexToAddress(contract), client)
	if err != nil {
		return nil, err
	}
	return instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

type LogTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func AnalysisTransferLog(log *types.Log) (from string, to string, value string, err error) {
	contractAbi, err := abi.JSON(strings.NewReader(string(GodsABI)))
	if err != nil {
		return "", "", "", errors.New(err.Error())
	}
	var transferEvent LogTransfer
	if err := contractAbi.Unpack(&transferEvent, "Transfer", log.Data); err != nil {
		return "", "", "", errors.New(err.Error())
	}
	return utils.HexFormat(transferEvent.From.Hex()), utils.HexFormat(transferEvent.To.Hex()), utils.BigIntToHex(transferEvent.TokenId), nil
}
