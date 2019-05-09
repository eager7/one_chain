package go_eth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/petermattis/goid"
	"math/big"
	"time"
)

type Eth struct {
	ctx        context.Context
	ws         string
	rpc        string
	chainId    *big.Int
	client     *ethclient.Client
	header     chan *types.Header
	channelEnd chan struct{}
	channelOut chan<- interface{}
}

func (e *Eth) Initialize(ctx context.Context, ws string) (*Eth, error) {
	e.ctx = ctx
	e.ws = ws
	client, err := ethclient.Dial(ws)
	if err != nil {
		return nil, errors.New("initialize eth client err:" + err.Error())
	}
	e.client = client
	e.header = make(chan *types.Header)
	if _, err := e.client.SubscribeNewHead(e.ctx, e.header); err != nil {
		return nil, errors.New(err.Error())
	}
	e.chainId, err = client.NetworkID(ctx)
	if err != nil {
		return nil, errors.New("get chain id err:" + err.Error())
	}
	return e, nil
}

func (e *Eth) Client() *ethclient.Client {
	return e.client
}

func (e *Eth) Header() chan *types.Header {
	return e.header
}

func (e *Eth) SetChannel(channelEnd chan struct{}, channelOut chan<- interface{}) {
	e.channelEnd = channelEnd
	e.channelOut = channelOut
}

func (e *Eth) Close() {
	close(e.header)
	e.client.Close()
}

func (e *Eth) Reset(ws string) error {
	if ws == e.ws {
		return nil
	}
	n, err := e.Initialize(e.ctx, ws)
	if err != nil {
		return errors.New("reset eth error:" + err.Error())
	}
	e.Close()
	e.ws = ws
	e.client = n.client
	return nil
}

func (e *Eth) SubscribeNewHeader() {
	headers := make(chan *types.Header)
	sub, err := e.client.SubscribeNewHead(e.ctx, headers)
	if err != nil {
		panic(err)
	}
	if e.channelEnd == nil || e.channelOut == nil {
		return
	}
	<-e.channelEnd //waiting for catching main net high
	fmt.Println("start subscribe block header")
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("receive err signal:", err)
			return
		case header := <-headers:
			time.Sleep(time.Microsecond * 100)
			fmt.Println("success get chain block:", goid.Get(), header.Number, header.Hash().Hex())
			if e.channelOut != nil {
				e.channelOut <- header
			}
		case <-e.ctx.Done():
			fmt.Println("stop header subscribe")
			return
		}
	}
}
