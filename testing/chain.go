package testing

import "context"

// EthereumLikeChain is an interface for Ethereum-like blockchains.
type EthereumLikeChain interface {
	Start(context.Context) error
	Stop() error
	Reset(context.Context) error
	SetGenesis(genesisPath string)
	GetGenesis() string
}
