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

package offchain

import (
	"path/filepath"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cast"

	cachekv "cosmossdk.io/store/cachekv"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
)

// Compile-time checks to ensure types are correct.
var _ storetypes.KVStore = (*Store)(nil)

// TODO: Upgrade this implementation to use a tree based structure to store the data off-chain.
// TODO: Replace TransientStoreType with a new type?

// `Store` represents a store used for storing persistent data off-chain while also
// utilizing the multistore. We must register this store with the multistore as a transient store,
// in order to ensure that we don't include the contents of this store in the chain's AppHash.
type Store struct {
	*cachekv.Store
}

// `NewOffChainKVStore` creates a new store and connects it to an file
// system based database located at: <flags.FlagHome/data/<name>.
func NewOffChainKVStore(name string, appOpts types.AppOptions) *Store {
	dbDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data")
	db, err := dbm.NewDB(name, server.GetAppDBBackend(appOpts), dbDir)
	if err != nil {
		panic(err)
	}

	return &Store{
		cachekv.NewStore(&dbadapter.Store{DB: db}),
	}
}

// `NewFromDB` creates a new store and connects it to the provided database.
func NewFromDB(db dbm.DB) *Store {
	return &Store{
		cachekv.NewStore(&dbadapter.Store{DB: db}),
	}
}
