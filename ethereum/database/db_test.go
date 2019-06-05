package database

import (
	"github.com/eager7/one_chain/ethereum/database/tables"
	"testing"
)

func TestCoin(t *testing.T) {
	db, err := Initialize("127.0.0.1:3306", "root", "plainchant", "one_chain_db", 100)
	if err != nil {
		t.Fatal(err)
	}
	coin := tables.TableCoinInfo{
		NameId:      "name id",
		Name:        "name",
		Symbol:      "symbol",
		Decimals:    10,
		Contract:    "contract",
		Transaction: "tx",
		Supply:      "supply",
		Price:       120,
		ICon:        "bai du.com",
	}
	cmd := coin.SqlCommand()
	if err := db.LogMode(true).Exec(cmd).Error; err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	db, err := Initialize("127.0.0.1:3306", "root", "plainchant", "one_chain_db", 100)
	if err != nil {
		t.Fatal(err)
	}
	db = db.LogMode(true)
	if err := db.Model(&tables.TableCoinInfo{}).Where("contract = ?", "pct").UpdateColumn("price", 121).Error; err != nil {
		t.Fatal(err)
	}
}

