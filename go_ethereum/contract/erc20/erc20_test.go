package erc20_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eager7/one_chain/common/utils"
	"github.com/eager7/one_chain/go_ethereum/contract/erc20"
	"github.com/eager7/one_chain/go_ethereum/go_eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
	"testing"
)

func TestTokenInfo(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://ropsten.infura.io/ws")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc20.ReadTokenInfo("0x19062dB4873E8d3752eA10d9038fe7C4ED6B1d4c", eth.Client(), "wss://ropsten.infura.io/ws"))
}

func TestBalanceAt(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc20.BalanceAt("0x93E4c75ff6e0C7BC4Bcfc7CA2F19e6733d6009Eb", "0xE2aCe83f9aCFC236bD95481045990bfa6B711E16", eth.Client()))
}

func TestListenTransferEvent(t *testing.T) {
	eth, err := new(go_eth.Eth).Initialize(context.Background(), "wss://mainnet.infura.io/ws")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(erc20.ListenTransferEvent(context.Background(), eth.Client()), "0xB8c77482e45F1F44dE1745F52C74426C631bDD52")
}

func TestAbi(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(contractAbi.Events)
	fmt.Println(contractAbi.Methods)
}

func TestAnalysisLogs(t *testing.T) {
	s := `{"address":"0x8c9b261faef3b3c2e64ab5e58e04615f8c788099","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007c00c9f0e7aed440c0c730a9bd9ee4f49de20d5c","0x000000000000000000000000000000000000000000000000000000000002a8c2"],"data":"0x","blockNumber":"0x730447","transactionHash":"0xf0603a6ae754eefaa90ad4c32dfe20a41afa419153f399cd896906bdc3ba7f3f","transactionIndex":"0xe","blockHash":"0xd7b52a24d5022166b09afb6d6d6795d0651b2eec7bab3c547d9c039fe4267e2c","logIndex":"0x12","removed":false}`
	log721 := types.Log{}
	if err := json.Unmarshal([]byte(s), &log721); err != nil {
		t.Fatal(err)
	}
	s = `{"address":"0x846c66cf71c43f80403b51fe3906b3599d63336f","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000007ed1e469fcb3ee19c0366d829e291451be638e59","0x00000000000000000000000005aa9df94c0cbf9532c0be0c3aec3782ab7dd344"],"data":"0x0000000000000000000000000000000000000000000013a5edf00e4709891200","blockNumber":"0x730447","transactionHash":"0x2bdc663f29203b81f034f174d3327da944c5e9e8c9025391f2bd7ec511e7e655","transactionIndex":"0xd","blockHash":"0xd7b52a24d5022166b09afb6d6d6795d0651b2eec7bab3c547d9c039fe4267e2c","logIndex":"0x11","removed":false}`
	log20 := types.Log{}
	if err := json.Unmarshal([]byte(s), &log20); err != nil {
		t.Fatal(err)
	}

	s = `  {
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
	if err := erc20.AnalysisLogs(log20, log721, logCat); err != nil {
		t.Fatal(err)
	}
}

func TestPackMessage(t *testing.T) {
	fmt.Println("balanceOf", "0x"+utils.HexToString(utils.ByteToHex(erc20.PackMessage("balanceOf", common.HexToAddress("0x93E4c75ff6e0C7BC4Bcfc7CA2F19e6733d6009Eb")))))
	fmt.Println("name", "0x"+utils.HexToString(utils.ByteToHex(erc20.PackMessage("name"))))
	fmt.Println("symbol", "0x"+utils.HexToString(utils.ByteToHex(erc20.PackMessage("symbol"))))
}
