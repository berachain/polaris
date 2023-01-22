// Copyright (C) 2022, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package state

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/cachemulti"
	"github.com/berachain/stargazer/core/state/types"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/crypto"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
)

var (
	// EmptyCodeHash is the code hash of an empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash      = crypto.Keccak256Hash(nil)
	emptyCodeHashBytes = emptyCodeHash.Bytes()
)

// Compile-time assertions to ensure StateDB adheres to StargazerStateDB and PrecompileStateDB.
var (
	_ vm.StargazerStateDB          = (*StateDB)(nil)
	_ precompile.PrecompileStateDB = (*StateDB)(nil)
)

// The StateDB is a very fun and interesting part of the EVM implementation. But if you want to
// join circus you need to know the rules. So here thet are:
//
//  1. You must ensure that the StateDB is only ever used in a single thread. This is because the
//     StateDB is not thread safe. And there are a bunch of optimizations made that are only safe
//     to do in a single thread.
//  2. You must discard the StateDB after use.
//  3. When accessing or mutating the StateDB, you must ensure that the underlying account exists.
//     in the AccountKeeper, for performance reasons, this implementation of the StateDB will not
//     create accounts that do not exist. Notably calling `SetState()` on an account that does not
//     exist is completely possible, and the StateDB will not prevent you doing so. This lazy
//     creation improves performance a ton, as it prevents calling into the ak on
//     every SSTORE. The only accounts that should ever have `SetState()` called on them are
//     accounts that represent smart contracts. Because of this assumption, the only place that we
//     explicitly create accounts is in `CreateAccount()`, since `CreateAccount()` is called when
//     deploying a smart contract.
//  4. Accounts that are sent `evmDenom` coins during an eth transaction, will have an account
//     created for them, automatically by the Bank Module. However, these accounts will have a
//     codeHash of 0x000... This is because the Bank Module does not know that the account is an
//     EVM account, and so it does not set the codeHash. This is totally fine, we just need to
//     check both for both the codeHash being 0x000... as well as the codeHash being 0x567...
type StateDB struct { //nolint: revive // we like the vibe.
	// We maintain a context in the StateDB, so that we can pass it with the correctly
	// configured multi-store to the precompiled contracts.
	ctx sdk.Context

	// Store a reference to the multi-store, in `ctx` so that we can access it directly.
	cms *cachemulti.Store

	// eth state stores: required for vm.StateDB
	// We store references to these stores, so that we can access them
	// directly, without having to go through the MultiStore interface.
	ethStore cachekv.StateDBCacheKVStore

	// keepers used for balance and account information.
	ak types.AccountKeeper
	bk types.BankKeeper

	// Any error that occurs during an sdk module read or write is
	// memoized here and is eventually be returned by `Commit`.
	savedErr error

	// we load the evm denom in the constructor, to prevent going to
	// the params to get it mid interpolation.
	evmDenom string // todo: get from params ( we have a store so like why not )

	// The refund counter, also used by state transitioning.
	refund uint64

	// The storekey used during execution
	storeKey storetypes.StoreKey

	// Per-transaction logs
	logs []*coretypes.Log

	// Transaction and logging bookkeeping
	txHash    common.Hash
	blockHash common.Hash
	txIndex   uint
	logIndex  uint

	// Dirty tracking of suicided accounts, we have to keep track of these manually, in order
	// for the code and state to still be accessible even after the account has been deleted.
	// We chose to keep track of them in a separate slice, rather than a map, because the
	// number of accounts that will be suicided in a single transaction is expected to be
	// very low.
	suicides []common.Address
}

// returns a *StateDB using the MultiStore belonging to ctx.
func NewStateDB(
	ctx sdk.Context,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	storeKey storetypes.StoreKey,
	evmDenom string,
) *StateDB {
	sdb := &StateDB{
		ak:       ak,
		bk:       bk,
		evmDenom: evmDenom,
		storeKey: storeKey,
	}

	// Wire up the `CacheMultiStore` & `sdk.Context`.
	sdb.cms = cachemulti.NewStoreFrom(ctx.MultiStore())
	sdb.ctx = ctx.WithMultiStore(sdb.cms)

	// Store a reference to the EVM state store for performance reasons.
	sdb.ethStore, _ = utils.GetAs[cachekv.StateDBCacheKVStore](sdb.cms.GetKVStore(sdb.storeKey))

	return sdb
}

// `GetContext` implements `StargazerStateDB`.
func (sdb *StateDB) GetContext() context.Context {
	return sdb.ctx
}

// ===========================================================================
// Account
// ===========================================================================

// CreateAccount implements the GethStateDB interface by creating a new account
// in the account keeper. It will allow accounts to be overridden.
func (sdb *StateDB) CreateAccount(addr common.Address) {
	acc := sdb.ak.NewAccountWithAddress(sdb.ctx, addr[:])

	// save the new account in the account keeper
	sdb.ak.SetAccount(sdb.ctx, acc)

	// initialize the code hash to empty
	sdb.ethStore.Set(types.CodeHashKeyFor(addr), emptyCodeHashBytes)
}

// =============================================================================
// Transaction Handling
// =============================================================================

// Prepare sets the current transaction hash and index and block hash which is
// used for logging events.
func (sdb *StateDB) PrepareForTransition(blockHash, txHash common.Hash, ti, li uint) {
	sdb.blockHash = blockHash
	sdb.txHash = txHash
	sdb.txIndex = ti
	sdb.logIndex = li
}

// Reset clears the journal and other state objects. It also clears the
// refund counter and the access list.
func (sdb *StateDB) Reset(ctx sdk.Context) {
	// TODO: figure out why not fully reallocating the object causes
	// the gas shit to fail
	// sdb.MultiStore = cachemulti.NewStoreFrom(ctx.MultiStore())
	// sdb.ctx = ctx.WithMultiStore(sdb.MultiStore)
	// // // Must support directly accessing the parent store.
	// // sdb.ethStore = sdb.ctx.cms.
	// // 	GetKVStore(sdb.storeKey).(cachekv.StateDBCacheKVStore)
	// sdb.savedErr = nil
	// sdb.refund = 0

	// sdb.logs = make([]*coretypes.Log, 0)
	// sdb.accessList = newAccessList()
	// sdb.suicides = make([]common.Address, 0)
	// TODO: unghetto this
	*sdb = *NewStateDB(ctx, sdb.ak, sdb.bk, sdb.storeKey, sdb.evmDenom)
}

// =============================================================================
// Balance
// =============================================================================

// GetBalance implements `GethStateDB` interface.
func (sdb *StateDB) GetBalance(addr common.Address) *big.Int {
	// Note: bank keeper will return 0 if account/state_object is not found
	return sdb.bk.GetBalance(sdb.ctx, addr[:], sdb.evmDenom).Amount.BigInt()
}

// AddBalance implements the `GethStateDB` interface by adding the given amount
// from the account associated with addr. If the account does not exist, it will be
// created.
func (sdb *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Mint the coins to the evm module account
	if err := sdb.bk.MintCoins(sdb.ctx, types.EvmNamespace, coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}

	// Send the coins from the evm module account to the destination address.
	if err := sdb.bk.SendCoinsFromModuleToAccount(sdb.ctx, types.EvmNamespace, addr[:], coins); err != nil {
		sdb.setErrorUnsafe(err)
	}
}

// SubBalance implements the `GethStateDB` interface by subtracting the given amount
// from the account associated with addr.
func (sdb *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the evm module account.
	if err := sdb.bk.SendCoinsFromAccountToModule(sdb.ctx, addr[:], types.EvmNamespace, coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}

	// Burn the coins from the evm module account.
	if err := sdb.bk.BurnCoins(sdb.ctx, types.EvmNamespace, coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}
}

// `TransferBalance` sends the given amount from one account to another. It will
// error if the sender does not have enough funds to send.
func (sdb *StateDB) TransferBalance(from, to common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the destination address.
	if err := sdb.bk.SendCoins(sdb.ctx, from[:], to[:], coins); err != nil {
		sdb.setErrorUnsafe(err)
	}
}

// =============================================================================
// Nonce
// =============================================================================

// GetNonce implements the `GethStateDB` interface by returning the nonce
// of an account.
func (sdb *StateDB) GetNonce(addr common.Address) uint64 {
	acc := sdb.ak.GetAccount(sdb.ctx, addr[:])
	if acc == nil {
		return 0
	}
	return acc.GetSequence()
}

// SetNonce implements the `GethStateDB` interface by setting the nonce
// of an account.
func (sdb *StateDB) SetNonce(addr common.Address, nonce uint64) {
	// get the account or create a new one if doesn't exist
	acc := sdb.ak.GetAccount(sdb.ctx, addr[:])
	if acc == nil {
		acc = sdb.ak.NewAccountWithAddress(sdb.ctx, addr[:])
	}

	if err := acc.SetSequence(nonce); err != nil {
		sdb.setErrorUnsafe(err)
	}

	sdb.ak.SetAccount(sdb.ctx, acc)
}

// =============================================================================
// Code
// =============================================================================

// GetCodeHash implements the `GethStateDB` interface by returning
// the code hash of account.
func (sdb *StateDB) GetCodeHash(addr common.Address) common.Hash {
	if sdb.ak.HasAccount(sdb.ctx, addr[:]) {
		if ch := sdb.ethStore.Get(types.CodeHashKeyFor(addr)); ch != nil {
			return common.BytesToHash(ch)
		}
		return emptyCodeHash
	}
	// if account at addr does not exist, return ZeroCodeHash
	return common.Hash{}
}

// GetCode implements the `GethStateDB` interface by returning
// the code of account (nil if not exists).
func (sdb *StateDB) GetCode(addr common.Address) []byte {
	codeHash := sdb.GetCodeHash(addr)
	// if account at addr does not exist, GetCodeHash returns ZeroCodeHash so return nil
	// if codeHash is empty, i.e. crypto.Keccak256(nil), also return nil
	if (codeHash == common.Hash{}) || codeHash == emptyCodeHash {
		return nil
	}
	return sdb.ethStore.Get(types.CodeKeyFor(codeHash))
}

// SetCode implements the `GethStateDB` interface by setting the code hash and
// code for the given account.
func (sdb *StateDB) SetCode(addr common.Address, code []byte) {
	codeHash := crypto.Keccak256Hash(code)

	sdb.ethStore.Set(types.CodeHashKeyFor(addr), codeHash[:])
	// store or delete code
	if len(code) == 0 {
		sdb.ethStore.Delete(types.CodeKeyFor(codeHash))
	} else {
		sdb.ethStore.Set(types.CodeKeyFor(codeHash), code)
	}
}

// GetCodeSize implements the `GethStateDB` interface by returning the size of the
// code associated with the given `GethStateDB`.
func (sdb *StateDB) GetCodeSize(addr common.Address) int {
	return len(sdb.GetCode(addr))
}

// =============================================================================
// Refund
// =============================================================================

// `AddRefund` implements the `GethStateDB` interface by adding gas to the
// refund counter.
func (sdb *StateDB) AddRefund(gas uint64) {
	sdb.cms.JournalMgr.Push(&RefundChange{sdb, sdb.refund})
	sdb.refund += gas
}

// `SubRefund` implements the `GethStateDB` interface by subtracting gas from the
// refund counter. If the gas is greater than the refund counter, it will panic.
func (sdb *StateDB) SubRefund(gas uint64) {
	sdb.cms.JournalMgr.Push(&RefundChange{sdb, sdb.refund})
	if gas > sdb.refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, sdb.refund))
	}
	sdb.refund -= gas
}

// `GetRefund` implements the `GethStateDB` interface by returning the current
// value of the refund counter.
func (sdb *StateDB) GetRefund() uint64 {
	return sdb.refund
}

// =============================================================================
// State
// =============================================================================

// `GetCommittedState` implements the `GethStateDB` interface by returning the
// committed state of slot in the given address.
func (sdb *StateDB) GetCommittedState(
	addr common.Address,
	slot common.Hash,
) common.Hash {
	return sdb.getStateFromStore(sdb.ethStore.GetParent(), addr, slot)
}

// `GetState` implements the `GethStateDB` interface by returning the current state
// of slot in the given address.
func (sdb *StateDB) GetState(addr common.Address, slot common.Hash) common.Hash {
	return sdb.getStateFromStore(sdb.ethStore, addr, slot)
}

// `getStateFromStore` returns the current state of the slot in the given address.
func (sdb *StateDB) getStateFromStore(
	store storetypes.KVStore,
	addr common.Address, slot common.Hash,
) common.Hash {
	if value := store.Get(types.StateKeyFor(addr, slot)); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// `SetState` implements the `GethStateDB` interface by setting the state of an
// address.
func (sdb *StateDB) SetState(addr common.Address, key, value common.Hash) {
	// For performance reasons, we don't check to ensure the account exists before we execute.
	// This is reasonably safe since under normal operation, SetState is only ever called by the
	// SSTORE opcode in the EVM, which will only ever be called on an account that exists, since
	// it would with 100% certainty have been created by a prior Create, thus setting its code
	// hash.
	//
	// CONTRACT: never manually call SetState outside of `opSstore`, or InitGenesis.

	// If empty value is given, delete the state entry.
	if len(value) == 0 || (value == common.Hash{}) {
		sdb.ethStore.Delete(types.StateKeyFor(addr, key))
		return
	}

	// Set the state entry.
	sdb.ethStore.Set(types.StateKeyFor(addr, key), value[:])
}

// =============================================================================
// Suicide
// =============================================================================

// Suicide implements the GethStateDB interface by marking the given address as suicided.
// This clears the account balance, but the code and state of the address remains available
// until after Commit is called.
func (sdb *StateDB) Suicide(addr common.Address) bool {
	// only smart contracts can commit suicide
	ch := sdb.GetCodeHash(addr)
	if (ch == common.Hash{}) || ch == emptyCodeHash {
		return false
	}

	// Reduce it's balance to 0.
	bal := sdb.bk.GetBalance(sdb.ctx, addr[:], sdb.evmDenom)
	sdb.SubBalance(addr, bal.Amount.BigInt())

	// Mark the underlying account for deletion in `Commit()`.
	sdb.suicides = append(sdb.suicides, addr)
	return true
}

// `HasSuicided` implements the `GethStateDB` interface by returning if the contract was suicided
// in current transaction.
func (sdb *StateDB) HasSuicided(addr common.Address) bool {
	for _, suicide := range sdb.suicides {
		if bytes.Equal(suicide[:], addr[:]) {
			return true
		}
	}
	return false
}

// =============================================================================
// Exist & Empty
// =============================================================================

// `Exist` implements the `GethStateDB` interface by reporting whether the given account address
// exists in the state. Notably this also returns true for suicided accounts, which is accounted
// for since, `RemoveAccount()` is not called until Commit.
func (sdb *StateDB) Exist(addr common.Address) bool {
	return sdb.ak.HasAccount(sdb.ctx, addr[:])
}

// `Empty` implements the `GethStateDB` interface by returning whether the state object
// is either non-existent or empty according to the EIP161 specification
// (balance = nonce = code = 0)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
func (sdb *StateDB) Empty(addr common.Address) bool {
	ch := sdb.GetCodeHash(addr)
	return sdb.GetNonce(addr) == 0 &&
		(ch == emptyCodeHash || ch == common.Hash{}) &&
		sdb.GetBalance(addr).Sign() == 0
}

// =============================================================================
// Snapshot
// =============================================================================

// `RevertToSnapshot` implements `StateDB`.
func (sdb *StateDB) RevertToSnapshot(id int) {
	// revert and discard all journal entries after snapshot id
	sdb.cms.JournalMgr.PopToSize(id)
}

// `Snapshot` implements `StateDB`.
func (sdb *StateDB) Snapshot() int {
	return sdb.cms.JournalMgr.Size()
}

// =============================================================================
// Logs
// =============================================================================

// AddLog implements the GethStateDB interface by adding the given log to the current transaction.
func (sdb *StateDB) AddLog(log *coretypes.Log) {
	sdb.cms.JournalMgr.Push(&AddLogChange{sdb})
	log.TxHash = sdb.txHash
	log.BlockHash = sdb.blockHash
	log.TxIndex = sdb.txIndex
	log.Index = sdb.logIndex
	sdb.logs = append(sdb.logs, log)
	sdb.logIndex++ // erigon intra
}

// Logs returns the logs of current transaction.
func (sdb *StateDB) Logs() []*coretypes.Log {
	return sdb.logs
}

// =============================================================================
// ForEachStorage
// =============================================================================

// `ForEachStorage` implements the `GethStateDB` interface by iterating through the contract state
// contract storage, the iteration order is not defined.
//
// Note: We do not support iterating through any storage that is modified before calling
// `ForEachStorage`; only committed state is iterated through.
func (sdb *StateDB) ForEachStorage(
	addr common.Address,
	cb func(key, value common.Hash) bool,
) error {
	it := sdk.KVStorePrefixIterator(sdb.ethStore, types.AddressStoragePrefix(addr))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		committedValue := it.Value()
		if len(committedValue) > 0 {
			if !cb(common.BytesToHash(it.Key()), common.BytesToHash(committedValue)) {
				// stop iteration
				return nil
			}
		}
	}

	return nil
}

// `Commit` is called when we are complete with the state transition and want to commit the changes
// to the underlying store.
func (sdb *StateDB) Commit() error {
	// If we saw an error during the execution, we return it here.
	if sdb.savedErr != nil {
		return sdb.savedErr
	}

	// Manually delete all suicidal accounts.
	for _, suicidalAddr := range sdb.suicides {
		acct := sdb.ak.GetAccount(sdb.ctx, suicidalAddr[:])
		if acct == nil {
			// handles the double suicide case
			continue
		}

		// clear storage
		if err := sdb.ForEachStorage(suicidalAddr,
			func(key, _ common.Hash) bool {
				sdb.SetState(suicidalAddr, key, common.Hash{})
				return true
			}); err != nil {
			return err
		}

		// clear the codehash from this account
		sdb.ethStore.Delete(types.CodeHashKeyFor(suicidalAddr))

		// remove auth account
		sdb.ak.RemoveAccount(sdb.ctx, acct)
	}

	// write all cache stores to parent stores, effectively writing temporary state in ctx to
	// the underlying parent store.
	sdb.cms.CacheMultiStore().Write()
	return nil
}

// =============================================================================
// Saved Errors
// =============================================================================

// Any errors that pop up during store operations should be checked here.
// Called upon the conclusion.
func (sdb *StateDB) GetSavedErr() error {
	return sdb.savedErr
}

// `setErrorUnsafe` sets error but should be called in medhods that already have locks.
func (sdb *StateDB) setErrorUnsafe(err error) {
	if sdb.savedErr == nil {
		sdb.savedErr = err
	}
}

// =============================================================================
// AccessList
// =============================================================================

func (sdb *StateDB) PrepareAccessList(
	sender common.Address,
	dst *common.Address,
	precompiles []common.Address,
	list coretypes.AccessList,
) {
	panic("not implemented, as accesslists are not valuable in the Cosmos-SDK context")
}

func (sdb *StateDB) AddAddressToAccessList(addr common.Address) {
	panic("not implemented, as accesslists are not valuable in the Cosmos-SDK context")
}

func (sdb *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	panic("not implemented, as accesslists are not valuable in the Cosmos-SDK context")
}

func (sdb *StateDB) AddressInAccessList(addr common.Address) bool {
	return false
}

func (sdb *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	return false, false
}

// =============================================================================
// PreImage
// =============================================================================

// AddPreimage implements the the `StateDB“ interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *StateDB) AddPreimage(hash common.Hash, preimage []byte) {}
