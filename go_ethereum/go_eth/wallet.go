package go_eth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
)

func (e *Eth) CreatePairKey() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", errors.New("generate key err:" + err.Error())
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return hexutil.Encode(privateKeyBytes)[2:], hexutil.Encode(publicKeyBytes)[4:], address, nil
}

func (e *Eth) CreateKeyStore(dir, password string) error {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return errors.New("create key store err:" + err.Error())
	}
	fmt.Println(account.Address.Hex())
	return nil
}

func (e *Eth) ImportKeyStore(dir, file, password string) error {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("import key store err:" + err.Error())
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		return errors.New("import key store err:" + err.Error())
	}
	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil {
		return errors.New("import key store err:" + err.Error())
	}
	return nil
}
