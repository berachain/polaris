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

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/lib/utils"
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PolarisApp struct {
	*runtime.App

	auxStoreKeys []storetypes.StoreKey

	wrappedTxPool *evmmempool.WrappedGethTxPool
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *PolarisApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	// Aux keys first
	for _, k := range app.auxStoreKeys {
		if kv, ok := utils.GetAs[*storetypes.KVStoreKey](k); ok && kv.Name() == storeKey {
			return kv
		}
	}

	// Then base keys
	kvStoreKey, ok := utils.GetAs[*storetypes.KVStoreKey](app.UnsafeFindStoreKey(storeKey))
	if !ok {
		return nil
	}
	return kvStoreKey
}

// KVStoreKeys returns all the registered KVStoreKeys.
func (app *PolarisApp) KVStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)

	// Aux keys first
	for _, k := range app.auxStoreKeys {
		if kv, ok := utils.GetAs[*storetypes.KVStoreKey](k); ok {
			keys[kv.Name()] = kv
		}
	}

	// Then base keys
	for _, k := range app.GetStoreKeys() {
		if kv, ok := utils.GetAs[*storetypes.KVStoreKey](k); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// // RegisterEthSecp256k1SignatureType registers the eth_secp256k1 signature type.
// func (app *App) RegisterEthSecp256k1SignatureType() {
// 	ethcryptocodec.RegisterInterfaces(app.CodecInterfaceRegistry)
// }

// MountCustomStore mounts a custom store to the baseapp.
// TODO: GET UPSTREAMED
func (app *PolarisApp) MountCustomStores(keys ...storetypes.StoreKey) {
	for _, key := range keys {
		// StoreTypeDB doesn't do anything upon commit, so its blessed for the offchain stuff.
		app.MountStore(key, storetypes.StoreTypeDB)
		app.auxStoreKeys = append(app.auxStoreKeys, key)
	}
}
