package simapp

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "pkg.berachain.dev/stargazer/crypto/codec"
	"pkg.berachain.dev/stargazer/x/evm/plugins/txpool/mempool"
)

// `StargazerAppOptions` is a list of `func(*baseapp.BaseApp)` that are used to configure the baseapp.
func StargazerAppOptions(
	interfaceRegistry types.InterfaceRegistry, baseAppOptions ...func(*baseapp.BaseApp),
) []func(*baseapp.BaseApp) {
	stargazerAppOptions := append(
		baseAppOptions,
		[]func(bApp *baseapp.BaseApp){
			baseapp.SetMempool(mempool.NewEthTxPool()),
			func(bApp *baseapp.BaseApp) {
				cryptocodec.RegisterInterfaces(interfaceRegistry)
			},
		}...,
	)
	return stargazerAppOptions
}
