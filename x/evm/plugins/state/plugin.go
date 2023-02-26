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
	"errors"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core"
	ethstate "pkg.berachain.dev/stargazer/eth/core/state"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/rpc"
	"pkg.berachain.dev/stargazer/lib/snapshot"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
	"pkg.berachain.dev/stargazer/store/snapmulti"
	"pkg.berachain.dev/stargazer/x/evm/plugins"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state/events"
	types "pkg.berachain.dev/stargazer/x/evm/types"
)

const pluginRegistryKey = `statePlugin`

var (
	// EmptyCodeHash is the code hash of an empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash      = crypto.Keccak256Hash(nil)
	emptyCodeHashBytes = emptyCodeHash.Bytes()
)

// `Plugin` is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosStargazer
	core.StatePlugin
	// `SetQueryContextFn` sets the query context func for the plugin.
	SetQueryContextFn(fn func(height int64, prove bool) (sdk.Context, error))
	// `IterateState` iterates over the state of all accounts and calls the given callback function.
	IterateState(fn func(addr common.Address, key common.Hash, value common.Hash) bool)
	// `IterateCode` iterates over the code of all accounts and calls the given callback function.
	IterateCode(fn func(addr common.Address, code []byte) bool)
}

// The StatePlugin is a very fun and interesting part of the EVM implementation. But if you want to
// join circus you need to know the rules. So here thet are:
//
//  1. You must ensure that the StatePlugin is only ever used in a single thread, because the
//     StatePlugin is not thread safe. And there are a bunch of optimizations made that are only
//     safe to do in a single thread.
//  2. When accessing or mutating the Plugin, you must ensure that the underlying account exists.
//     In the AccountKeeper, for performance reasons, this implementation of the StateDB will not
//     create accounts that do not exist. Notably calling `SetState()` on an account that does not
//     exist is completely possible, and the StateDB will not prevent you doing so. This lazy
//     creation improves performance a ton, as it prevents calling into the ak on
//     every SSTORE. The only accounts that should ever have `SetState()` called on them are
//     accounts that represent smart contracts. Because of this assumption, the only place that we
//     explicitly create accounts is in `CreateAccount()`, since `CreateAccount()` is called when
//     deploying a smart contract.
//  3. Accounts that are sent `evmDenom` coins during an eth transaction, will have an account
//     created for them, automatically by the Bank Module. However, these accounts will have a
//     codeHash of 0x000... This is because the Bank Module does not know that the account is an
//     EVM account, and so it does not set the codeHash. This is totally fine, we just need to
//     check both for both the codeHash being zero (0x000...) as well as the codeHash being empty
//     (0x567...)
type plugin struct {
	libtypes.Controller[string, libtypes.Controllable[string]]

	// We maintain a context in the StateDB, so that we can pass it with the correctly
	// configured multi-store to the precompiled contracts.
	ctx sdk.Context

	// Store a reference to the multi-store, in `ctx` so that we can access it directly.
	cms ControllableMultiStore

	// Store a reference to the Precompile Log Factory, which builds Eth logs from Cosmos events
	plf events.PrecompileLogFactory

	// Store the evm store key for quick lookups to the evm store
	storeKey storetypes.StoreKey

	// keepers used for balance and account information.
	ak AccountKeeper
	bk BankKeeper

	// getQueryContext allows for querying state a historical height.
	getQueryContext func(height int64, prove bool) (sdk.Context, error)

	// we load the evm denom in the constructor, to prevent going to
	// the params to get it mid interpolation.
	evmDenom string // TODO: get from configuration plugin.
}

// `NewPlugin` returns a plugin with the given context and keepers.
func NewPlugin(
	ak AccountKeeper,
	bk BankKeeper,
	storeKey storetypes.StoreKey,
	evmDenom string,
	plf events.PrecompileLogFactory,
) Plugin {
	return &plugin{
		storeKey: storeKey,
		ak:       ak,
		bk:       bk,
		evmDenom: evmDenom,
		plf:      plf,
	}
}

// `Reset` implements `core.StatePlugin`.
func (p *plugin) Reset(ctx context.Context) {
	// reset the Controllable MultiStore and EventManager and attach them to the context
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	p.cms = snapmulti.NewStoreFrom(sdkCtx.MultiStore())
	cem := events.NewManagerFrom(sdkCtx.EventManager(), p.plf)
	p.ctx = sdkCtx.WithMultiStore(p.cms).WithEventManager(cem)

	// setup the snapshot controller
	ctrl := snapshot.NewController[string, libtypes.Controllable[string]]()
	_ = ctrl.Register(p.cms)
	_ = ctrl.Register(cem)
	p.Controller = ctrl
}

// `RegistryKey` implements `libtypes.Registrable`.
func (p *plugin) RegistryKey() string {
	return pluginRegistryKey
}

// ===========================================================================
// Account
// ===========================================================================

// CreateAccount implements the `StatePlugin` interface by creating a new account
// in the account keeper. It will allow accounts to be overridden.
func (p *plugin) CreateAccount(addr common.Address) {
	acc := p.ak.NewAccountWithAddress(p.ctx, addr[:])

	// save the new account in the account keeper
	p.ak.SetAccount(p.ctx, acc)

	// initialize the code hash to empty
	p.cms.GetKVStore(p.storeKey).Set(CodeHashKeyFor(addr), emptyCodeHashBytes)
}

// `Exist` implements the `StatePlugin` interface by reporting whether the given account address
// exists in the state. Notably this also returns true for suicided accounts, which is accounted
// for since, `RemoveAccount()` is not called until Commit.
func (p *plugin) Exist(addr common.Address) bool {
	return p.ak.HasAccount(p.ctx, addr[:])
}

// =============================================================================
// Balance
// =============================================================================

// GetBalance implements `StatePlugin` interface.
func (p *plugin) GetBalance(addr common.Address) *big.Int {
	// Note: bank keeper will return 0 if account/state_object is not found
	return p.bk.GetBalance(p.ctx, addr[:], p.evmDenom).Amount.BigInt()
}

// SetBalance implements `StatePlugin` interface.
func (p *plugin) SetBalance(addr common.Address, amount *big.Int) {
	currBalance := p.GetBalance(addr)
	delta := new(big.Int).Sub(currBalance, amount)
	if delta.Sign() < 0 {
		p.AddBalance(addr, new(big.Int).Neg(delta))
	} else if delta.Sign() > 0 {
		p.SubBalance(addr, delta)
	}
}

// AddBalance implements the `StatePlugin` interface by adding the given amount
// from the account associated with addr. If the account does not exist, it will be
// created.
func (p *plugin) AddBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(p.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Mint the coins to the evm module account
	if err := p.bk.MintCoins(p.ctx, types.ModuleName, coins); err != nil {
		panic(err)
	}

	// Send the coins from the evm module account to the destination address.
	if err := p.bk.SendCoinsFromModuleToAccount(
		p.ctx, types.ModuleName, addr[:], coins,
	); err != nil {
		panic(err)
	}
}

// SubBalance implements the `StatePlugin` interface by subtracting the given amount
// from the account associated with addr.
func (p *plugin) SubBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(p.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the evm module account.
	if err := p.bk.SendCoinsFromAccountToModule(
		p.ctx, addr[:], types.ModuleName, coins,
	); err != nil {
		panic(err)
	}

	// Burn the coins from the evm module account.
	if err := p.bk.BurnCoins(p.ctx, types.ModuleName, coins); err != nil {
		panic(err)
	}
}

// `TransferBalance` sends the given amount from one account to another. It will
// error if the sender does not have enough funds to send.
func (p *plugin) TransferBalance(from, to common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(p.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the destination address.
	if err := p.bk.SendCoins(p.ctx, from[:], to[:], coins); err != nil {
		// This is safe to panic as the error is only returned if the sender does
		// not have enough funds to send, which should be guarded by `CanTransfer`.
		panic(err)
	}
}

// =============================================================================
// Nonce
// =============================================================================

// GetNonce implements the `StatePlugin` interface by returning the nonce
// of an account.
func (p *plugin) GetNonce(addr common.Address) uint64 {
	acc := p.ak.GetAccount(p.ctx, addr[:])
	if acc == nil {
		return 0
	}
	return acc.GetSequence()
}

// SetNonce implements the `StatePlugin` interface by setting the nonce
// of an account.
func (p *plugin) SetNonce(addr common.Address, nonce uint64) {
	// get the account or create a new one if doesn't exist
	acc := p.ak.GetAccount(p.ctx, addr[:])
	if acc == nil {
		acc = p.ak.NewAccountWithAddress(p.ctx, addr[:])
	}

	if err := acc.SetSequence(nonce); err != nil {
		panic(err)
	}

	p.ak.SetAccount(p.ctx, acc)
}

// =============================================================================
// Code
// =============================================================================

// GetCodeHash implements the `StatePlugin` interface by returning
// the code hash of account.
func (p *plugin) GetCodeHash(addr common.Address) common.Hash {
	if !p.ak.HasAccount(p.ctx, addr[:]) {
		// if account at addr does not exist, return zeros
		return common.Hash{}
	}

	ch := p.cms.GetKVStore(p.storeKey).Get(CodeHashKeyFor(addr))
	if ch == nil {
		// account exists but does not have a codehash, return empty
		return emptyCodeHash
	}

	return common.BytesToHash(ch)
}

// GetCode implements the `StatePlugin` interface by returning
// the code of account (nil if not exists).
func (p *plugin) GetCode(addr common.Address) []byte {
	codeHash := p.GetCodeHash(addr)
	if (codeHash == common.Hash{}) || codeHash == emptyCodeHash {
		// if account at addr does not exist or the account  does not have a codehash, return nil
		return nil
	}
	return p.cms.GetKVStore(p.storeKey).Get(CodeKeyFor(codeHash))
}

// SetCode implements the `StatePlugin` interface by setting the code hash and
// code for the given account.
func (p *plugin) SetCode(addr common.Address, code []byte) {
	codeHash := crypto.Keccak256Hash(code)
	ethStore := p.cms.GetKVStore(p.storeKey)
	ethStore.Set(CodeHashKeyFor(addr), codeHash[:])

	// store or delete code
	if len(code) == 0 {
		ethStore.Delete(CodeKeyFor(codeHash))
	} else {
		ethStore.Set(CodeKeyFor(codeHash), code)
	}
}

// `IterateCode` iterates over all the addresses with code and calls the given method.
func (p *plugin) IterateCode(fn func(address common.Address, code []byte) bool) {
	it := storetypes.KVStorePrefixIterator(
		p.cms.GetKVStore(p.storeKey),
		[]byte{keyPrefixCodeHash},
	)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		addr := AddressFromCodeHashKey(it.Key())
		if fn(addr, p.GetCode(addr)) {
			break
		}
	}
}

// GetCodeSize implements the `StatePlugin` interface by returning the size of the
// code associated with the given `StatePlugin`.
func (p *plugin) GetCodeSize(addr common.Address) int {
	return len(p.GetCode(addr))
}

// =============================================================================
// State
// =============================================================================

// `GetCommittedState` implements the `StatePlugin` interface by returning the
// committed state of slot in the given address.
func (p *plugin) GetCommittedState(
	addr common.Address,
	slot common.Hash,
) common.Hash {
	return getStateFromStore(p.cms.GetCommittedKVStore(p.storeKey), addr, slot)
}

// `GetState` implements the `StatePlugin` interface by returning the current state
// of slot in the given address.
func (p *plugin) GetState(addr common.Address, slot common.Hash) common.Hash {
	return getStateFromStore(p.cms.GetKVStore(p.storeKey), addr, slot)
}

// `SetState` sets the state of an address.
func (p *plugin) SetState(addr common.Address, key, value common.Hash) {
	// For performance reasons, we don't check to ensure the account exists before we execute.
	// This is reasonably safe since under normal operation, SetState is only ever called by the
	// SSTORE opcode in the EVM, which will only ever be called on an account that exists, since
	// it would with 100% certainty have been created by a prior Create, thus setting its code
	// hash.
	//
	// CONTRACT: never manually call SetState outside of `opSstore`, or InitGenesis.

	// If empty value is given, delete the state entry.
	if len(value) == 0 || (value == common.Hash{}) {
		p.cms.GetKVStore(p.storeKey).Delete(SlotKeyFor(addr, key))
		return
	}

	// Set the state entry.
	p.cms.GetKVStore(p.storeKey).Set(SlotKeyFor(addr, key), value[:])
}

// `SetStorage` sets the storage of an address.
func (p *plugin) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	for key, value := range storage {
		p.SetState(addr, key, value)
	}
}

// `IterateState` iterates over all the contract state, and calls the given function.
func (p *plugin) IterateState(cb func(addr common.Address, key, value common.Hash) bool) {
	it := storetypes.KVStorePrefixIterator(
		p.cms.GetCommittedKVStore(p.storeKey),
		[]byte{keyPrefixStorage},
	)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		k, v := it.Key(), it.Value()
		addr := AddressFromSlotKey(k)
		slot := SlotFromSlotKey(k)
		if cb(addr, slot, common.BytesToHash(v)) {
			break
		}
	}
}

// =============================================================================
// ForEachStorage
// =============================================================================

// `ForEachStorage` implements the `StatePlugin` interface by iterating through the contract state
// contract storage, the iteration order is not defined.
//
// Note: We do not support iterating through any storage that is modified before calling
// `ForEachStorage`; only committed state is iterated through.
func (p *plugin) ForEachStorage(
	addr common.Address,
	cb func(key, value common.Hash) bool,
) error {
	it := storetypes.KVStorePrefixIterator(
		p.cms.GetKVStore(p.storeKey),
		StorageKeyFor(addr),
	)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		committedValue := it.Value()
		if len(committedValue) > 0 {
			if !cb(common.BytesToHash(it.Key()), common.BytesToHash(committedValue)) {
				return nil // stop iteration
			}
		}
	}

	return nil
}

// `DeleteSuicides` manually deletes the given suicidal accounts.
func (p *plugin) DeleteSuicides(suicides []common.Address) {
	for _, suicidalAddr := range suicides {
		acct := p.ak.GetAccount(p.ctx, suicidalAddr[:])
		if acct == nil {
			// handles the double suicide case
			continue
		}

		// clear storage
		_ = p.ForEachStorage(suicidalAddr,
			func(key, _ common.Hash) bool {
				p.SetState(suicidalAddr, key, common.Hash{})
				return true
			})

		// clear the codehash from this account
		p.cms.GetKVStore(p.storeKey).Delete(CodeHashKeyFor(suicidalAddr))

		// remove auth account
		p.ak.RemoveAccount(p.ctx, acct)
	}
}

// =============================================================================
// Historical State
// =============================================================================

// `SetQueryContextFn` sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// `GetStateByNumber` implements `core.StatePlugin`.
func (p *plugin) GetStateByNumber(number int64) (vm.GethStateDB, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("no query context function set in host chain")
	}
	// Handle rpc.BlockNumber negative numbers.
	var iavlHeight int64
	//nolint:exhaustive // this has to be a golangci-lint bug.
	switch rpc.BlockNumber(number) {
	case rpc.SafeBlockNumber:
	case rpc.FinalizedBlockNumber:
		iavlHeight = p.ctx.BlockHeight() - 1
	case rpc.PendingBlockNumber:
	case rpc.LatestBlockNumber:
		iavlHeight = p.ctx.BlockHeight()
	case rpc.EarliestBlockNumber:
		iavlHeight = 0
	}

	// Get the query context at the given height.
	ctx, err := p.getQueryContext(iavlHeight, false)
	if err != nil {
		return nil, err
	}

	// Create a StateDB with the requested chain height.
	sp := NewPlugin(p.ak, p.bk, p.storeKey, p.evmDenom, p.plf)
	sp.Reset(ctx)
	return ethstate.NewStateDB(sp), nil
}

// `getStateFromStore` returns the current state of the slot in the given address.
func getStateFromStore(
	store storetypes.KVStore,
	addr common.Address, slot common.Hash,
) common.Hash {
	if value := store.Get(SlotKeyFor(addr, slot)); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}
