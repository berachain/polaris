package config

import (
	"github.com/ethereum/go-ethereum/node"
	"github.com/spf13/viper"
	"pkg.berachain.dev/polaris/eth/polar"
)

const ()

// Config is the base config for both out-of-process and in-process oracles. If the oracle is to be configured out-of-process in base-app, a
// grpc-client of the grpc-server running at RemoteAddress is instantiated, otherwise, an in-process LocalClient oracle is instantiated.
type Config struct {
	Polar polar.Config
	Node  node.Config
}

// ReadConfigFromFile reads a config from a file and returns the config.
func ReadConfigFromFile(path string) (*Config, error) {
	// read in config file
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// unmarshal config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
