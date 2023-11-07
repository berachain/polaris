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

package state

import (
	"context"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events"
	libtypes "github.com/berachain/polaris/lib/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ControllableEventManager defines a cache EventManager that is controllable (snapshottable
// and registrable). It also supports precompile execution by allowing the caller to native events
// as Eth logs.
type ControllableEventManager interface {
	libtypes.Controllable[string]
	sdk.EventManagerI

	// BeginPrecompileExecution begins a precompile execution by setting the logs DB.
	BeginPrecompileExecution(events.LogsDB)
	// EndPrecompileExecution ends a precompile execution by resetting the logs DB to nil.
	EndPrecompileExecution()

	// IsReadOnly returns true if the EventManager is read-only.
	IsReadOnly() bool
	// SetReadOnly sets the EventManager to the given read-only mode.
	SetReadOnly(bool)
}

// ControllableMultiStore defines a cache MultiStore that is controllable (snapshottable and
// registrable). It also supports getting the committed KV store from the MultiStore.
type ControllableMultiStore interface {
	libtypes.Controllable[string]
	storetypes.MultiStore

	// GetCommittedKVStore returns the committed KV store from the MultiStore.
	GetCommittedKVStore(storetypes.StoreKey) storetypes.KVStore
}

// AccountKeeper defines the expected account keeper.
type AccountKeeper interface {
	AddressCodec() address.Codec
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetSequence(context.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	SetAccount(ctx context.Context, account sdk.AccountI)
	RemoveAccount(ctx context.Context, account sdk.AccountI)
	IterateAccounts(ctx context.Context, cb func(account sdk.AccountI) bool)
}
