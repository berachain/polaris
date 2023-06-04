package runtime

import (
	"encoding/json"
	"io"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// AppBuilder is a type that is injected into a container by the runtime module
// (as *AppBuilder) which can be used to create an app which is compatible with
// the existing app.go initialization conventions.
type AppBuilder struct {
	// Used to build the baseapp
	*runtime.AppBuilder

	//  used for export.
	polarisApp *PolarisApp
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (a *AppBuilder) DefaultGenesis() map[string]json.RawMessage {
	return a.AppBuilder.DefaultGenesis()
}

// Build builds an *App instance.
func (a *AppBuilder) Build(db dbm.DB, traceStore io.Writer, baseAppOptions ...func(*baseapp.BaseApp)) *PolarisApp {
	a.polarisApp = &PolarisApp{}

	// TODO: move this somewhere better, introduce non IAVL enforced module keys as a PR to the SDK
	// we ask @tac0turtle how 2 fix
	offchainKey := storetypes.NewKVStoreKey("offchain-evm")

	// Build the base runtime.App (and thus baseapp.BaseApp)
	a.polarisApp.App = a.AppBuilder.Build(db, traceStore, baseAppOptions...)

	// Mount our custom stores.
	a.polarisApp.MountCustomStores(offchainKey)

	// Return the app
	return a.polarisApp
}
