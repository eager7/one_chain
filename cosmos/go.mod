module github.com/eager7/one_chain/cosmos

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.34.4
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.5
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
