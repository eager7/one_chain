package ck_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eager7/one_chain/go_ethereum/contract/ck"
	"github.com/eager7/one_chain/go_ethereum/go_eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"testing"
)

func TestAnalysisLogs(t *testing.T) {
	s := `  {
                "address": "0x06012c8cf97bead5deae237070f9587f8e7a266d",
                "blockHash": "0x2d91c21669f3efbe17e6fdf3b537fb7601821b2c53da8a59c6cadaf7d51da00d",
                "blockNumber": "0x7304a0",
                "data": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000007c00c9f0e7aed440c0c730a9bd9ee4f49de20d5c0000000000000000000000000000000000000000000000000000000000171b18",
                "logIndex": "0xd",
                "removed": false,
                "topics": [
                    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
                ],
                "transactionHash": "0xf649a6e13e6cb225b11d5336a4149ac44f709bf575bce67efad9770428dd4705",
                "transactionIndex": "0x18"
            }`
	logCat := types.Log{}
	if err := json.Unmarshal([]byte(s), &logCat); err != nil {
		t.Fatal(err)
	}
	if err := ck.AnalysisLogs(logCat); err != nil {
		t.Fatal(err)
	}
}

func TestTokenOfOwner(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws")
	if err != nil {
		t.Fatal(err)
	}
	list, err := ck.TokensOfOwner("0x48bf0f9c84d3E449c7b1cc09C95FcbA1F8bb948c", "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d", eth.Client())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("lists:", list)

	balance, err := ck.BalanceAt("0x48bf0f9c84d3E449c7b1cc09C95FcbA1F8bb948c", "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d", eth.Client())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance:", balance)

	owner, err := ck.TokenOwnerOf("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d", new(big.Int).SetUint64(195297), eth.Client())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(owner)
}

func TestAbi(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(ck.CkABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(contractAbi.Events)
	fmt.Println(contractAbi.Methods)
}

func TestPacket(t *testing.T) {
	fmt.Println("owner:", "0x"+ck.PackMessage("tokensOfOwner", common.HexToAddress("0xa096b47ebf7727d01ff4f09c34fc6591f2c375f0")))
	fmt.Println("balanceOf", "0x"+ck.PackMessage("balanceOf", common.HexToAddress("0x48bf0f9c84d3E449c7b1cc09C95FcbA1F8bb948c")))
}
