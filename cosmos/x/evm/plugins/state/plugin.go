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

	"pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/store/snapmulti"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/rpc"
	"pkg.berachain.dev/polaris/lib/snapshot"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

const pluginRegistryKey = `statePlugin`

var (
	// EmptyCodeHash is the code hash of an empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash      = crypto.Keccak256Hash(nil)
	emptyCodeHashBytes = emptyCodeHash.Bytes()
)

// Plugin is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosPolaris
	core.StatePlugin
	// SetQueryContextFn sets the query context func for the plugin.
	SetQueryContextFn(fn func(height int64, prove bool) (sdk.Context, error))
	// IterateState iterates over the state of all accounts and calls the given callback function.
	IterateState(fn func(addr common.Address, key common.Hash, value common.Hash) bool)
	// IterateCode iterates over the code of all accounts and calls the given callback function.
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

	// Store a precompile log factory that builds Eth logs from Cosmos events
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
	cp ConfigurationPlugin
}

// NewPlugin returns a plugin with the given context and keepers.
func NewPlugin(
	ak AccountKeeper,
	bk BankKeeper,
	storeKey storetypes.StoreKey,
	cp ConfigurationPlugin,
	plf events.PrecompileLogFactory,
) Plugin {
	return &plugin{
		storeKey: storeKey,
		ak:       ak,
		bk:       bk,
		cp:       cp,
		plf:      plf,
	}
}

// Reset implements `core.StatePlugin`.
func (p *plugin) Reset(ctx context.Context) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// We have to build a custom `SnapMulti` to use with the StateDB. This is because the
	// ethereum utilizes the concept of snapshots, whereas the current implementation of the
	// Cosmos-SDK `CacheKV` uses "wraps".
	p.cms = snapmulti.NewStoreFrom(sdkCtx.MultiStore())

	// We have to build a custom event manager to use with the StateDB. This is because the we want
	// a way to handle converting Cosmos events from precompiles into Ethereum logs.
	cem := events.NewManagerFrom(sdkCtx.EventManager(), p.plf)

	// We need to build a custom configuration for the context in order to handle precompile event logs
	// and proper gas consumption.
	p.ctx = sdkCtx.WithMultiStore(p.cms).WithEventManager(cem)

	// We  also remove the KVStore gas metering from the context prior to entering the EVM
	// state transition. This is because the EVM is not aware of the Cosmos SDK's gas metering
	// and is designed to be used in a standalone manner, as each of the EVM's opcodes are priced individually.
	// By setting the gas configs to empty structs, we ensure that SLOADS and SSTORES in the EVM
	// are not being charged additional gas unknowingly.
	p.ctx = p.ctx.WithKVGasConfig(storetypes.GasConfig{}).WithTransientKVGasConfig(storetypes.GasConfig{})

	// We setup a snapshot controller in order to properly handle reverts.
	ctrl := snapshot.NewController[string, libtypes.Controllable[string]]()

	// We register the Controllable MultiStore with the snapshot controller.
	if err := ctrl.Register(p.cms); err != nil {
		panic(err)
	}

	// We also register the Controllable EventManager with the snapshot controller.
	if err := ctrl.Register(cem); err != nil {
		panic(err)
	}
	p.Controller = ctrl
}

// GetContext implements `core.StatePlugin`.
func (p *plugin) GetContext() context.Context {
	return p.ctx
}

// RegistryKey implements `libtypes.Registrable`.
func (p *plugin) RegistryKey() string {
	return pluginRegistryKey
}

// ===========================================================================
// Accounts
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

// Exist implements the `StatePlugin` interface by reporting whether the given account address
// exists in the state. Notably this also returns true for suicided accounts, which is accounted
// for since, `RemoveAccount()` is not called until Commit.
func (p *plugin) Exist(addr common.Address) bool {
	return p.ak.HasAccount(p.ctx, addr[:])
}

// Empty implements the `PolarisStateDB` interface by returning whether the state object
// is either non-existent or empty according to the EIP161 epecification
// (balance = nonce = code = 0)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
func (p *plugin) Empty(addr common.Address) bool {
	ch := p.GetCodeHash(addr)
	return p.GetNonce(addr) == 0 &&
		(ch == emptyCodeHash || ch == common.Hash{}) &&
		p.GetBalance(addr).Sign() == 0
}

// `DeleteAccounts` manually deletes the given accounts.
func (p *plugin) DeleteAccounts(accounts []common.Address) {
	for _, account := range accounts {
		acct := p.ak.GetAccount(p.ctx, account[:])
		if acct == nil {
			// handles the double suicide case
			continue
		}

		// clear storage
		_ = p.ForEachStorage(account,
			func(key, _ common.Hash) bool {
				p.SetState(account, key, common.Hash{})
				return true
			})

		// clear the codehash from this account
		p.cms.GetKVStore(p.storeKey).Delete(CodeHashKeyFor(account))

		// remove auth account
		p.ak.RemoveAccount(p.ctx, acct)
	}
}

// =============================================================================
// Balance
// =============================================================================

// GetBalance implements `StatePlugin` interface.
func (p *plugin) GetBalance(addr common.Address) *big.Int {
	// Note: bank keeper will return 0 if account/state_object is not found
	return p.bk.GetBalance(p.ctx, addr[:], p.cp.GetEvmDenom()).Amount.BigInt()
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
// from thew account associated with addr. If the account does not exist, it will be
// created.
func (p *plugin) AddBalance(addr common.Address, amount *big.Int) {
	err := lib.MintCoinsToAddress(p.ctx, p.bk, types.ModuleName, addr, p.cp.GetEvmDenom(), amount)
	if err != nil {
		panic(err)
	}
}

// SubBalance implements the `StatePlugin` interface by subtracting the given amount
// from the account associated with addr.
func (p *plugin) SubBalance(addr common.Address, amount *big.Int) {
	err := lib.BurnCoinsFromAddress(p.ctx, p.bk, types.ModuleName, addr, p.cp.GetEvmDenom(), amount)
	if err != nil {
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

// IterateCode iterates over all the addresses with code and calls the given method.
func (p *plugin) IterateCode(fn func(address common.Address, code []byte) bool) {
	it := storetypes.KVStorePrefixIterator(
		p.cms.GetKVStore(p.storeKey),
		[]byte{types.CodeHashKeyPrefix},
	)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		addr := AddressFromCodeHashKey(it.Key())
		if fn(addr, p.GetCode(addr)) {
			break
		}
	}
}

// =============================================================================
// Storage
// =============================================================================

// GetCommittedState implements the `StatePlugin` interface by returning the
// committed state of slot in the given address.
func (p *plugin) GetCommittedState(
	addr common.Address,
	slot common.Hash,
) common.Hash {
	return getStateFromStore(p.cms.GetCommittedKVStore(p.storeKey), addr, slot)
}

// GetState implements the `StatePlugin` interface by returning the current state
// of slot in the given address.
func (p *plugin) GetState(addr common.Address, slot common.Hash) common.Hash {
	return getStateFromStore(p.cms.GetKVStore(p.storeKey), addr, slot)
}

// SetState sets the state of an address.
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

// SetStorage sets the storage of an address.
func (p *plugin) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	for key, value := range storage {
		p.SetState(addr, key, value)
	}
}

// IterateState iterates over all the contract state, and calls the given function.
func (p *plugin) IterateState(cb func(addr common.Address, key, value common.Hash) bool) {
	it := storetypes.KVStorePrefixIterator(
		p.cms.GetCommittedKVStore(p.storeKey),
		[]byte{types.StorageKeyPrefix},
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

// ForEachStorage implements the `StatePlugin` interface by iterating through the contract state
// contract storage, the iteration order is not defined.
//
// Note: We do not support iterating through any storage that is modified before calling
// ForEachStorage; only committed state is iterated through.
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

// getStateFromStore returns the current state of the slot in the given address.
func getStateFromStore(
	store storetypes.KVStore,
	addr common.Address, slot common.Hash,
) common.Hash {
	if value := store.Get(SlotKeyFor(addr, slot)); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// =============================================================================
// Historical State
// =============================================================================

// SetQueryContextFn sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// GetStateByNumber implements `core.StatePlugin`.
func (p *plugin) GetStateByNumber(number int64) (core.StatePlugin, error) {
	if p.getQueryContext == nil {
		return nil, errors.New("no query context function set in host chain")
	}
	// Handle rpc.BlockNumber negative numbers.
	var iavlHeight int64
	switch rpc.BlockNumber(number) { //nolint:nolintlint,exhaustive // golangci-lint bug?
	case rpc.SafeBlockNumber, rpc.FinalizedBlockNumber:
		iavlHeight = p.ctx.BlockHeight() - 1
	case rpc.PendingBlockNumber, rpc.LatestBlockNumber:
		iavlHeight = p.ctx.BlockHeight()
	case rpc.EarliestBlockNumber:
		iavlHeight = 1
	default:
		iavlHeight = number
	}

	var ctx sdk.Context
	if p.ctx.BlockHeight() == iavlHeight {
		ctx, _ = p.ctx.CacheContext()
	} else {
		// Get the query context at the given height.
		var err error
		ctx, err = p.getQueryContext(iavlHeight, false)
		if err != nil {
			return nil, err
		}
	}

	// Create a State Plugin with the requested chain height.
	sp := NewPlugin(p.ak, p.bk, p.storeKey, p.cp, p.plf)
	sp.Reset(ctx)
	return sp, nil
}
