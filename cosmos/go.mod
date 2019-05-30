module github.com/eager7/one_chain/cosmos

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.34.4
	github.com/eager7/elog v0.0.0-20190424070304-6500c8e87726
	github.com/robfig/cron v1.1.0 // indirect
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.5
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
