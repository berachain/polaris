// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
