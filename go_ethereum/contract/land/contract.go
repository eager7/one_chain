package land

import (
	"context"
	"errors"
	"fmt"
	"github.com/eager7/one_chain/common/utils"
	"github.com/eager7/one_chain/go_ethereum/evm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

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
	contractAbi, err := abi.JSON(strings.NewReader(string(LandABI)))
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
	contractAbi, err := abi.JSON(strings.NewReader(string(LandABI)))
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
	contractAbi, err := abi.JSON(strings.NewReader(string(LandABI)))
	if err != nil {
		return "", "", "", errors.New(err.Error())
	}
	var transferEvent LogTransfer
	if err := contractAbi.Unpack(&transferEvent, "Transfer", log.Data); err != nil {
		return "", "", "", errors.New(err.Error())
	}
	return utils.HexFormat(transferEvent.From.Hex()), utils.HexFormat(transferEvent.To.Hex()), utils.BigIntToHex(transferEvent.TokenId), nil
}
