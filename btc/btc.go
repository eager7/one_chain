package btc

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

type Btc struct {
	client      *rpcclient.Client
	chainParams *chaincfg.Params
}

func Initialize(url, usr, pass string) (*Btc, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         url,
		Endpoint:     "ws",
		User:         usr,
		Pass:         pass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	params := chaincfg.MainNetParams
	return &Btc{client: client, chainParams: &params}, nil
}

func (b *Btc) CheckTransactionStandard(raw string) (string, error) {
	data, err := hex.DecodeString(raw)
	if err != nil {
		return "", errors.New("decode raw err:" + err.Error())
	}
	ret, err := b.client.DecodeRawTransaction(data)
	if err != nil {
		return "", errors.New("decode tx err:" + err.Error())
	}
	s, _ := json.Marshal(ret)
	return string(s), nil
}
