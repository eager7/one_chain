package erc721

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
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func BalanceAt(address, contract string, client *ethclient.Client) (*big.Int, error) {
	instance, err := NewErc721(common.HexToAddress(contract), client)
	if err != nil {
		return nil, errors.New("new token err:" + err.Error())
	}
	return instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func TokenOwnerOf(contract string, tokenId *big.Int, client *ethclient.Client) (string, error) {
	instance, err := NewErc721(common.HexToAddress(contract), client)
	if err != nil {
		return "", err
	}
	owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

func SupportERC721Interface(contract string, client *ethclient.Client) (bool, error) {
	instance, err := NewErc721(common.HexToAddress(contract), client)
	if err != nil {
		return false, err
	}

	//Check Eip165 Interface
	ret, err := instance.SupportsInterface(&bind.CallOpts{}, evm.In165Interface)
	if err != nil {
		return false, err
	}
	if !ret {
		return false, errors.New("can't implement eip165 interface by eip165 first")
	}
	ret, err = instance.SupportsInterface(&bind.CallOpts{}, evm.InNonInterface)
	if err != nil {
		return false, err
	}
	if ret {
		return false, errors.New("can't implement eip165 interface by none")
	}
	ret, err = instance.SupportsInterface(&bind.CallOpts{}, evm.In165Interface)
	if err != nil {
		return false, err
	}
	if !ret {
		return false, errors.New("can't implement eip165 interface by eip165 first")
	}
	//Check Eip721 Interface
	ret, err = instance.SupportsInterface(&bind.CallOpts{}, evm.In721Interface)
	if err != nil {
		return false, err
	}
	if !ret {
		return false, errors.New("can't implement eip721 interface")
	}
	return true, nil
}

func ReadTokenInfo(address string, client *ethclient.Client) (string, string, *big.Int, error) {
	instance, err := NewErc721(common.HexToAddress(address), client)
	if err != nil {
		return "", "", nil, errors.New("new erc721 err:" + err.Error())
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		fmt.Println("can't get name:", err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", nil, errors.New("symbol err:" + err.Error())
	}
	supply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return "", "", nil, errors.New("supply err:%v" + err.Error())
	}
	return name, symbol, supply, nil
}

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func ListenTransferEvent(ctx context.Context, client *ethclient.Client, address ...common.Address) error {
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: address,
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(Erc721ABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.EIP165Event("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch utils.HexFormat(l.Topics[0].Hex()) {
		case logTransferSigHash:
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			err := contractAbi.Unpack(&transferEvent, "Transfer", l.Data)
			if err != nil {
				return err
			}
			transferEvent.From = common.HexToAddress(l.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(l.Topics[2].Hex())
			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())
		default:
			fmt.Println("unknown topic:", l.Topics[0].Hex())
		}
	}
	return nil
}

func AnalysisLogs() {

}
