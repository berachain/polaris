package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// InitGenesis is called during the InitGenesis.
func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) error {
	// We configure the logger here because we want to get the logger off the context opposed to allocating a new one.
	k.ConfigureGethLogger(ctx)

	// TODO: remove InitGenesis from the interfaces, do check and run instead
	// Initialize all the plugins.
	for _, plugin := range k.host.GetAllPlugins() {
		plugin.InitGenesis(ctx, &genState)
	}

	// Start the polaris "Node" in order to spin up things like the JSON-RPC server.
	if err := k.polaris.StartServices(); err != nil {
		return err
	}
	return nil
}
