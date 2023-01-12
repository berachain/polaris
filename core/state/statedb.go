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
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gevm "github.com/ethereum/go-ethereum/core/vm"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/cachemulti"
	"github.com/berachain/stargazer/core/state/types"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/lib/crypto"
)

type GethStateDB = gevm.StateDB

// ExtStateDB defines an extension to the interface provided by the go-ethereum codebase to
// support additional state transition functionalities. In particular it supports getting the
// cosmos sdk context for natively running stateful precompiled contracts.
type ExtStateDB interface {
	storetypes.MultiStore
	GethStateDB

	// GetContext returns the cosmos sdk context with the statedb multistore attached
	GetContext() sdk.Context

	// GetSavedErr returns the error saved in the statedb
	GetSavedErr() error

	// GetLogs returns the logs generated during the transaction
	Logs() []*coretypes.Log

	// Commit writes the state to the underlying multi-store
	Commit() error
}

type IntraBlockStateDB interface {
	ExtStateDB

	// PrepareForTransition prepares the statedb for a new transition
	// by setting the block hash, tx hash, tx index and tx log index.
	PrepareForTransition(common.Hash, common.Hash, uint, uint)

	// Reset resets the statedb to the initial state.
	Reset(sdk.Context)
}

const (
	keyPrefixCode byte = iota
	keyPrefixHash
)

var (
	// EmptyCodeHash is the code hash of an empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	EmptyCodeHash = crypto.Keccak256Hash(nil)
	ZeroCodeHash  = common.Hash{}

	KeyPrefixCode     = []byte{keyPrefixCode}
	KeyPrefixCodeHash = []byte{keyPrefixHash}
	// todo: add later?
	// StateDBStoreKey                         = sdk.NewKVStoreKey(types.StoreKey).
	_ ExtStateDB = (*StateDB)(nil)
)

// Welcome to the StateDB implementation. This is a wrapper around a cachemulti.Store
// so that precompiles querying Eth-modified state can directly read from the statedb
// objects. It adheres to the Geth vm.StateDB and Cosmos MultiStore interfaces, which allows it
// to be used as a MultiStore in the Cosmos SDK context.
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
	*cachemulti.Store

	// eth state stores: required for vm.StateDB
	// We store references to these stores, so that we can access them
	// directly, without having to go through the MultiStore interface.
	liveEthState cachekv.StateDBCacheKVStore

	// cosmos state ctx: required for sdk.MultiStore
	ctx sdk.Context

	// keepers used for balance and account information
	ak AccountKeeper
	bk BankKeeper

	// DB error.
	// State objects are used by the consensus core and VM which are
	// unable to deal with database-level errors. Any error that occurs
	// during a database read is memoized here and will eventually be returned
	// by StateDB.Commit.
	savedErr error

	// we load the evm denom in the constructor, to prevent going to
	// the params to get it mid interpolation.
	evmDenom string

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
	ak AccountKeeper,
	bk BankKeeper,
	storeKey storetypes.StoreKey,
	evmDenom string,
) *StateDB {
	sdb := &StateDB{
		Store:    cachemulti.NewStoreFrom(ctx.MultiStore()),
		ak:       ak,
		bk:       bk,
		evmDenom: evmDenom,
		storeKey: storeKey,
	}
	sdb.ctx = ctx.WithMultiStore(sdb)

	// Must support directly accessing the parent store.
	sdb.liveEthState, _ = sdb.ctx.MultiStore().
		GetKVStore(storeKey).(cachekv.StateDBCacheKVStore)
	return sdb
}

func (sdb *StateDB) GetEvmDenom() string {
	return sdb.evmDenom
}

// CreateAccount implements the GethStateDB interface by creating a new account
// in the account keeper. It will allow accounts to be overridden.
func (sdb *StateDB) CreateAccount(addr common.Address) {
	acc := sdb.ak.NewAccountWithAddress(sdb.ctx, addr[:])

	// save the new account in the account keeper
	sdb.ak.SetAccount(sdb.ctx, acc)

	// initialize the code hash to empty
	prefix.NewStore(sdb.liveEthState, KeyPrefixCodeHash).Set(addr[:], EmptyCodeHash[:])
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
	// // sdb.liveEthState = sdb.ctx.MultiStore().
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

// GetBalance implements GethStateDB interface.
func (sdb *StateDB) GetBalance(addr common.Address) *big.Int {
	// Note: bank keeper will return 0 if account/state_object is not found
	return sdb.bk.GetBalance(sdb.ctx, addr[:], sdb.evmDenom).Amount.BigInt()
}

// AddBalance implements the GethStateDB interface by adding the given amount
// from the account associated with addr. If the account does not exist, it will be
// created.
func (sdb *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Mint the coins to the evm module account
	if err := sdb.bk.MintCoins(sdb.ctx, "evm", coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}

	// Send the coins from the evm module account to the destination address.
	if err := sdb.bk.SendCoinsFromModuleToAccount(sdb.ctx, "evm", addr[:], coins); err != nil {
		sdb.setErrorUnsafe(err)
	}
}

// SubBalance implements the GethStateDB interface by subtracting the given amount
// from the account associated with addr.
func (sdb *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the evm module account.
	if err := sdb.bk.SendCoinsFromAccountToModule(sdb.ctx, addr[:], "evm", coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}

	// Burn the coins from the evm module account.
	if err := sdb.bk.BurnCoins(sdb.ctx, "evm", coins); err != nil {
		sdb.setErrorUnsafe(err)
		return
	}
}

// `SendBalance` sends the given amount from one account to another. It will
// error if the sender does not have enough funds to send.
func (sdb *StateDB) SendBalance(from, to common.Address, amount *big.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(sdb.evmDenom, sdk.NewIntFromBigInt(amount)))

	// Send the coins from the source address to the destination address.
	if err := sdb.bk.SendCoins(sdb.ctx, from[:], to[:], coins); err != nil {
		sdb.setErrorUnsafe(err)
	}
}

// =============================================================================
// Nonce
// =============================================================================

// GetNonce implements the GethStateDB interface by returning the nonce
// of an account.
func (sdb *StateDB) GetNonce(addr common.Address) uint64 {
	acc := sdb.ak.GetAccount(sdb.ctx, addr[:])
	if acc == nil {
		return 0
	}
	return acc.GetSequence()
}

// SetNonce implements the GethStateDB interface by setting the nonce
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

// GetCodeHash implements the GethStateDB interface by returning
// the code hash of account.
func (sdb *StateDB) GetCodeHash(addr common.Address) common.Hash {
	if sdb.ak.HasAccount(sdb.ctx, addr[:]) {
		if ch := prefix.NewStore(sdb.liveEthState,
			KeyPrefixCodeHash).Get(addr[:]); ch != nil {
			return common.BytesToHash(ch)
		}
		return EmptyCodeHash
	}
	// if account at addr does not exist, return ZeroCodeHash
	return ZeroCodeHash
}

// GetCode implements the GethStateDB interface by returning
// the code of account (nil if not exists).
func (sdb *StateDB) GetCode(addr common.Address) []byte {
	codeHash := sdb.GetCodeHash(addr)
	// if account at addr does not exist, GetCodeHash returns ZeroCodeHash so return nil
	// if codeHash is empty, i.e. crypto.Keccak256(nil), also return nil
	if codeHash == ZeroCodeHash || codeHash == EmptyCodeHash {
		return nil
	}
	return prefix.NewStore(sdb.liveEthState, KeyPrefixCode).Get(codeHash.Bytes())
}

// SetCode implements the GethStateDB interface by setting the code hash and
// code for the given account.
func (sdb *StateDB) SetCode(addr common.Address, code []byte) {
	codeHash := crypto.Keccak256Hash(code)

	prefix.NewStore(sdb.liveEthState, KeyPrefixCodeHash).Set(addr[:], codeHash[:])

	// store or delete code
	if len(code) == 0 {
		prefix.NewStore(sdb.liveEthState, KeyPrefixCode).Delete(codeHash[:])
	} else {
		prefix.NewStore(sdb.liveEthState, KeyPrefixCode).Set(codeHash[:], code)
	}
}

// GetCodeSize implements the GethStateDB interface by returning the size of the
// code associated with the given GethStateDB.
func (sdb *StateDB) GetCodeSize(addr common.Address) int {
	return len(sdb.GetCode(addr))
}

// =============================================================================
// Refund
// =============================================================================

// AddRefund implements the GethStateDB interface by adding gas to the
// refund counter.
func (sdb *StateDB) AddRefund(gas uint64) {
	sdb.JournalMgr.Push(&RefundChange{sdb, sdb.refund})
	sdb.refund += gas
}

// SubRefund implements the GethStateDB interface by subtracting gas from the
// refund counter. If the gas is greater than the refund counter, it will panic.
func (sdb *StateDB) SubRefund(gas uint64) {
	sdb.JournalMgr.Push(&RefundChange{sdb, sdb.refund})
	if gas > sdb.refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, sdb.refund))
	}
	sdb.refund -= gas
}

// GetRefund implements the GethStateDB interface by returning the current
// value of the refund counter.
func (sdb *StateDB) GetRefund() uint64 {
	return sdb.refund
}

// =============================================================================
// State
// =============================================================================

// GetCommittedState implements the GethStateDB interface by returning the
// committed state of an address.
func (sdb *StateDB) GetCommittedState(
	addr common.Address,
	hash common.Hash,
) common.Hash {
	if value := prefix.NewStore(sdb.liveEthState.GetParent(),
		types.AddressStoragePrefix(addr)).Get(hash[:]); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// GetState implements the GethStateDB interface by returning the current state of an
// address.
func (sdb *StateDB) GetState(addr common.Address, hash common.Hash) common.Hash {
	if value := prefix.NewStore(sdb.liveEthState,
		types.AddressStoragePrefix(addr)).Get(hash[:]); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// SetState implements the GethStateDB interface by setting the state of an
// address.
func (sdb *StateDB) SetState(addr common.Address, key, value common.Hash) {
	// For performance reasons, we don't check to ensure the account exists before we execute.
	// This is reasonably safe since under normal operation, SetState is only ever called by the
	// SSTORE opcode in the EVM, which will only ever be called on an account that exists, since
	// it would with 100% certainty have been created by a prior Create, thus setting its code
	// hash.
	// CONTRACT: never manually call SetState outside of the the `opSstore` in the interpreter.

	store := prefix.NewStore(sdb.liveEthState, types.AddressStoragePrefix(addr))
	if len(value) == 0 || value == ZeroCodeHash {
		store.Delete(key[:])
	} else {
		store.Set(key[:], value[:])
	}
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
	if ch == ZeroCodeHash || ch == EmptyCodeHash {
		return false
	}

	// Reduce it's balance to 0.
	bal := sdb.bk.GetBalance(sdb.ctx, addr[:], sdb.evmDenom)
	sdb.SubBalance(addr, bal.Amount.BigInt())

	// Mark the underlying account for deletion in `Commit()`.
	sdb.suicides = append(sdb.suicides, addr)
	return true
}

// HasSuicided implements the GethStateDB interface by returning if the contract was suicided
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

// Exist implements the GethStateDB interface by reporting whether the given account address
// exists in the state. Notably this also returns true for suicided accounts, which is accounted
// for since, `RemoveAccount()` is not called until Commit.
func (sdb *StateDB) Exist(addr common.Address) bool {
	return sdb.ak.HasAccount(sdb.ctx, addr[:])
}

// Empty implements the GethStateDB interface by returning whether the state object
// is either non-existent or empty according to the EIP161 specification
// (balance = nonce = code = 0)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
func (sdb *StateDB) Empty(addr common.Address) bool {
	ch := sdb.GetCodeHash(addr)
	return sdb.GetNonce(addr) == 0 &&
		(ch == EmptyCodeHash || ch == ZeroCodeHash) &&
		sdb.GetBalance(addr).Sign() == 0
}

// =============================================================================
// Snapshot
// =============================================================================

// `RevertToSnapshot` implements `StateDB`.
func (sdb *StateDB) RevertToSnapshot(id int) {
	// revert and discard all journal entries after snapshot id
	sdb.JournalMgr.PopToSize(id)
}

func (sdb *StateDB) Snapshot() int {
	return sdb.JournalMgr.Size()
}

// =============================================================================
// Logs
// =============================================================================

// AddLog implements the GethStateDB interface by adding the given log to the current transaction.
func (sdb *StateDB) AddLog(log *coretypes.Log) {
	sdb.JournalMgr.Push(&AddLogChange{sdb})
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

// ForEachStorage implements the GethStateDB interface by iterating through the contract state
// contract storage, the iteration order is not defined.
//
// Note: We do not support iterating through any storage that is modified before calling
// ForEachStorage; only committed state is iterated through.
func (sdb *StateDB) ForEachStorage(
	addr common.Address,
	cb func(key, value common.Hash) bool,
) error {
	it := sdk.KVStorePrefixIterator(sdb.liveEthState, types.AddressStoragePrefix(addr))
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

// AddPreimage implements the the GethStateDB interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *StateDB) AddPreimage(hash common.Hash, preimage []byte) {}

// =============================================================================
// MultiStore
// =============================================================================

// Commit implements storetypes.MultiStore by writing the dirty states of the
// liveStateCtx to the committedStateCtx. It also handles sucidal accounts.
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
		prefix.NewStore(sdb.liveEthState, KeyPrefixCodeHash).Delete(suicidalAddr[:])

		// remove auth account
		sdb.ak.RemoveAccount(sdb.ctx, acct)
	}

	// write all cache stores to parent stores, effectively writing temporary state in ctx to
	// the underlying parent store.
	sdb.Store.Write()
	return nil
}

// =============================================================================
// ExtStateDB
// =============================================================================

// GetContext implements ExtStateDB
// returns the StateDB's live context.
func (sdb *StateDB) GetContext() sdk.Context {
	return sdb.ctx
}

// GetSavedErr implements ExtStateDB
// any errors that pop up during store operations should be checked here
// called upon the conclusion.
func (sdb *StateDB) GetSavedErr() error {
	return sdb.savedErr
}

// setErrorUnsafe sets error but should be called in medhods that already have locks.
func (sdb *StateDB) setErrorUnsafe(err error) {
	if sdb.savedErr == nil {
		sdb.savedErr = err
	}
}

// =============================================================================
// Genesis
// =============================================================================

// // ImportStateData imports the given genesis accounts into the state database. We pass in a
// // temporary context to bypass the StateDB caching capabilities to speed up the import.
// func (sdb *StateDB) ImportStateData(accounts []types.GenesisAccount) error {
// 	for _, account := range accounts {
// 		addr := common.HexToAddress(account.Address)

// 		acc := sdb.ak.GetAccount(sdb.ctx, addr[:])
// 		if acc == nil {
// 			panic(fmt.Errorf("account not found for address %s", account.Address))
// 		}

// 		code := common.Hex2Bytes(account.Code)
// 		codeHash := crypto.Keccak256Hash(code)
// 		storedCodeHash := sdb.GetCodeHash(addr)

// 		// we ignore the empty Code hash checking, see ethermint PR#1234
// 		if len(account.Code) != 0 && storedCodeHash != codeHash {
// 			s := "the evm state code doesn't match with the codehash\n"
// 			panic(fmt.Sprintf("%s account: %s , evm state codehash: %v,"+
// 				" ethAccount codehash: %v, evm state code: %s\n",
// 				s, account.Address, codeHash, storedCodeHash, account.Code))
// 		}

// 		// Set the code for this contract
// 		sdb.SetCode(addr, code)

// 		// Set the state for this contract
// 		for _, storage := range account.Storage {
// 			sdb.SetState(addr, common.HexToHash(storage.Key), common.HexToHash(storage.Value))
// 		}

// 		// Commit the changes to save them to the underlying KVStore.
// 		if err := sdb.Commit(); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // ExportStateData exports all the data stored in the StateDB to a GenesisAccount slice.
// // We pass in a temporary context to bypass the StateDB caching capabilities to
// // speed up the import.
// func (sdb *StateDB) ExportStateData() (ethAccs []types.GenesisAccount) {
// 	sdb.ak.IterateAccounts(sdb.ctx, func(account authtypes.AccountI) bool {
// 		addr := common.BytesToAddress(account.GetAddress().Bytes())

// 		// Load Storage
// 		var storage types.Storage
// 		if err := sdb.ForEachStorage(addr,
// 			func(key, value common.Hash) bool {
// 				storage = append(storage, types.NewState(key, value))
// 				return true
// 			},
// 		); err != nil {
// 			panic(err)
// 		}

// 		genAccount := types.GenesisAccount{
// 			Address: addr.String(),
// 			Code:    common.Bytes2Hex(sdb.GetCode(addr)),
// 			Storage: storage,
// 		}

// 		ethAccs = append(ethAccs, genAccount)
// 		return false
// 	})
// 	return ethAccs
// }

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
	panic("not implemented, as accesslists are not valuable in the Cosmos-SDK context")
}

func (sdb *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	panic("not implemented, as accesslists are not valuable in the Cosmos-SDK context")
}
