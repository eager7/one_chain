package erc721_test

import (
	"context"
	"fmt"
	"github.com/eager7/one_chain/go_ethereum/contract/erc721"
	"github.com/eager7/one_chain/go_ethereum/go_eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"testing"
)

func TestTokenInfo(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc721.ReadTokenInfo("0xB8c77482e45F1F44dE1745F52C74426C631bDD52", eth.Client()))
}

func TestListenTransferEvent(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc721.ListenTransferEvent(context.Background(), eth.Client()), "0xB8c77482e45F1F44dE1745F52C74426C631bDD52")
}

func TestAbi(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(erc721.Erc721ABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(contractAbi.Events)
	fmt.Println(contractAbi.Methods)
}

func TestSupportInterface(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc721.SupportERC721Interface("0x8c9b261Faef3b3C2e64ab5E58e04615F8c788099", eth.Client()))
}
