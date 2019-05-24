package go_cos

import "github.com/spf13/viper"

func LoadConfigFile(file string) {
	viper.SetConfigFile(file)
	viper.SetDefault("node", "tcp://39.104.144.212:26657")
	viper.SetDefault("trust-node", "true")
}
