package block

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/core"
)

// InitGenesis performs genesis initialization for the evm module. It returns
// no validator updates.
func (p *plugin) InitGenesis(ctx sdk.Context, _ *core.Genesis) {
	// TODO: IMPLEMENT
	p.Prepare(ctx)
	// p.StoreHeader(ethGen.ToBlock().Header())
}

// ExportGenesis returns the exported genesis state as raw bytes for the evm
// module.
func (p *plugin) ExportGenesis(ctx sdk.Context, _ *core.Genesis) {
	// TODO: IMPLEMENT
	p.Prepare(ctx)
	// head, err := p.GetHeaderByNumber(0)
	// if err != nil {
	// 	panic(err)
	// }
}
