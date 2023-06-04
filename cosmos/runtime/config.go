package runtime

// import (
// 	"cosmossdk.io/log"

// 	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
// 	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
// 	"cosmossdk.io/core/appconfig"
// 	"cosmossdk.io/depinject"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/runtime"

// 	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/txpool/mempool"
// )

// func MakePolarisConfig(logger log.Logger, precompiles ) depinject.Config {
// 	ethTxMempool := evmmempool.NewPolarisEthereumTxPool()
// 	var mempoolOpt runtime.BaseAppOption = baseapp.SetMempool(ethTxMempool)
// 	return depinject.Configs(
// 		depinject.Supply(mempoolOpt),
// 		appconfig.Compose(&appv1alpha1.Config{
// 			Modules: []*appv1alpha1.ModuleConfig{
// 				{
// 					Name: "runtime",
// 					Config: appconfig.WrapAny(&runtimev1alpha1.Module{
// 						AppName: "BaseAppApp",
// 					}),
// 				},
// 			},
// 		}))
// }
