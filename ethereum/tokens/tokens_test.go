package tokens

import (
	"fmt"
	"github.com/eager7/one_chain/ethereum/database"
	"testing"
)

func TestCoinList(t *testing.T) {
	list, err := CoinList(UrlCoinList)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}

func TestCoinInformation(t *testing.T) {
	info, err := CoinInformation(UrlCoinsInfo, "airtoken")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", info)
}

func TestEthListAll(t *testing.T) {
	db, err := database.Initialize("127.0.0.1:3306", "root", "plainchant", "one_chain_db", 100)
	if err != nil {
		t.Fatal(err)
	}
	ListAllEthTokens(db)
}

func TestTokenPrice(t *testing.T) {
	db, err := database.Initialize("127.0.0.1:3306", "root", "plainchant", "one_chain_db", 100)
	if err != nil {
		t.Fatal(err)
	}
	if err := TokenPrice(db); err != nil {
		t.Fatal(err)
	}
}
