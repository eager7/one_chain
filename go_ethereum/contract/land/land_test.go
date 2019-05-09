package land

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
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
	if err := AnalysisLogs(logCat); err != nil {
		t.Fatal(err)
	}
}

func TestEventList(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(LandABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(contractAbi.Events)
}
