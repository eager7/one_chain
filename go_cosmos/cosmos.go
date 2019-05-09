package go_cosmos

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/rpc/core/types"
)

type CosCli struct {
	cos *context.CLIContext
	cdc *amino.Codec
}

func Initialize(file string) *CosCli {
	LoadConfigFile(file)
	ctx := context.NewCLIContext()
	fmt.Println("dir:", viper.ConfigFileUsed())
	fmt.Println("url:", ctx.NodeURI)

	var cdc = amino.NewCodec()
	core_types.RegisterAmino(cdc)

	return &CosCli{cos: &ctx, cdc: cdc}
}
