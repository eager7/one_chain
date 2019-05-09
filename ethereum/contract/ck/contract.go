package ck

import (
	"context"
	"errors"
	"fmt"
	"github.com/eager7/one_chain/common/utils"
	"github.com/eager7/one_chain/go_ethereum/evm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func BalanceAt(address, contract string, client *ethclient.Client) (*big.Int, error) {
	instance, err := NewCk(common.HexToAddress(contract), client)
	if err != nil {
		return nil, errors.New("new token err:" + err.Error())
	}
	return instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func TokensOfOwner(address, contract string, client *ethclient.Client) ([]*big.Int, error) {
	instance, err := NewCk(common.HexToAddress(contract), client)
	if err != nil {
		return nil, errors.New("new token err:" + err.Error())
	}
	list, err := instance.TokensOfOwner(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil && !strings.Contains(err.Error(), "empty output") {
		return nil, err
	}
	return list, nil
}

func TokenOwnerOf(contract string, tokenId *big.Int, client *ethclient.Client) (string, error) {
	instance, err := NewCk(common.HexToAddress(contract), client)
	if err != nil {
		return "", err
	}
	owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

func PackMessage(name string, args ...interface{}) string {
	cAbi, err := abi.JSON(strings.NewReader(string(CkABI)))
	if err != nil {
		return ""
	}
	data, err := cAbi.Pack(name, args...)
	if err != nil {
		fmt.Println(err)
	}
	return utils.ByteToHex(data)
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

func ListenTransferEvent(ctx context.Context, client *ethclient.Client, address ...common.Address) error {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(7537735),
		ToBlock:   nil,
		Addresses: address,
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(CkABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.EIP165Event("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch utils.HexFormat(l.Topics[0].Hex()) {
		case logTransferSigHash:
			fmt.Println("Log Name: Transfer\n", utils.JsonString(l))
			var transferEvent LogTransfer
			err := contractAbi.Unpack(&transferEvent, "Transfer", l.Data)
			if err != nil {
				return err
			}
			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.TokenId.String())
		default:
			fmt.Println("unknown topic:", l.Topics[0].Hex())
		}
	}
	return nil
}

func AnalysisLogs(logs ...types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(CkABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.KnownEvent("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch utils.HexFormat(l.Topics[0].Hex()) {
		case logTransferSigHash:
			var transferEvent LogTransfer
			err := contractAbi.Unpack(&transferEvent, "Transfer", l.Data)
			if err != nil {
				return errors.New(err.Error())
			}
			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.TokenId)
		default:
			fmt.Println("unknown topic:", l.Topics[0].Hex())
		}
	}
	return nil
}

func AnalysisTransferLog(log *types.Log) (from string, to string, value string, err error) {
	contractAbi, err := abi.JSON(strings.NewReader(string(CkABI)))
	if err != nil {
		return "", "", "", errors.New(err.Error())
	}
	var transferEvent LogTransfer
	if err := contractAbi.Unpack(&transferEvent, "Transfer", log.Data); err != nil {
		return "", "", "", errors.New(err.Error())
	}
	return utils.HexFormat(transferEvent.From.Hex()), utils.HexFormat(transferEvent.To.Hex()), utils.BigIntToHex(transferEvent.TokenId), nil
}
