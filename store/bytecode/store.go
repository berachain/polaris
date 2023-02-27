package bytecode

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
import (
	"path/filepath"

	cachekv "cosmossdk.io/store/cachekv"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

// Compile-time interface assertion.
var _ storetypes.KVStore = (*Store)(nil)

// `Store` is a wrapper around `cachekv.Store` that implements the `storetypes.KVStore`
// interface. It is used to store the bytecode of contracts offchain. This is done to
// reduce the size of the state and to avoid the need to store the bytecode in the
// state.
type Store struct {
	*cachekv.Store
}

// `NewByteCodeStore` creates a new store and connects it to an file system based database
// located at: <flags.FlagHome/data/bytecode>.
func NewByteCodeStore(appOpts types.AppOptions) *Store {
	dbDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data")
	db, err := dbm.NewDB("bytecode", server.GetAppDBBackend(appOpts), dbDir)
	if err != nil {
		panic(err)
	}

	return &Store{
		cachekv.NewStore(&dbadapter.Store{DB: db}),
	}
}

// `NewFromDB` creates a new store from a given database.
func NewFromDB(db dbm.DB) *Store {
	return &Store{
		cachekv.NewStore(&dbadapter.Store{DB: db}),
	}
}
