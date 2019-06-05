module github.com/eager7/one_chain/bnb

go 1.12

require (
	github.com/binance-chain/go-sdk v1.0.4
	github.com/eager7/elog v0.0.0-20190424070304-6500c8e87726
	github.com/robfig/cron v1.1.0 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/tendermint/tendermint v0.31.2-rc0
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1
