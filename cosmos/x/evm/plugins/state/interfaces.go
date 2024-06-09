// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
