// Copyright (C) 2023, Berachain Foundation. All rights reserved.
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
	"math/big"

	"github.com/berachain/stargazer/eth/core/state/plugin"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
)

var (
	// EmptyCodeHash is the Keccak256 Hash of empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash = crypto.Keccak256Hash(nil)
)

// Compile-time assertion that StateDB implements the vm.StargazerStateDB interface.
var _ vm.StargazerStateDB = (*StateDB)(nil)

type StateDB struct { //nolint:revive // StateDB is a struct that holds the state of the blockchain.
	ctx context.Context

	// The controller is used to manage the plugins
	ctrl Controller

	// Developer provided plugins
	ap AccountPlugin
	bp BalancePlugin
	cp CodePlugin
	sp StoragePlugin

	// Internal plugins
	lp LogPlugin
	rf RefundPlugin

	// Dirty tracking of suicided accounts, we have to keep track of these manually, in order
	// for the code and state to still be accessible even after the account has been deleted.
	// We chose to keep track of them in a separate slice, rather than a map, because the
	// number of accounts that will be suicided in a single transaction is expected to be
	// very low.
	suicides []common.Address
}

func NewStateDB(ctrl Controller) *StateDB {
	// Add the internal plugins to the controller
	ctrl.AddPlugin(plugin.NewRefund())
	ctrl.AddPlugin(plugin.NewLogs())

	// Create the stateDB and populate the developer provided plugins.
	return &StateDB{
		ctrl:     ctrl,
		ap:       ctrl.GetPlugin("account").(AccountPlugin),
		bp:       ctrl.GetPlugin("balance").(BalancePlugin),
		cp:       ctrl.GetPlugin("code").(CodePlugin),
		sp:       ctrl.GetPlugin("storage").(StoragePlugin),
		lp:       ctrl.GetPlugin("logs").(LogPlugin),
		rf:       ctrl.GetPlugin("refund").(RefundPlugin),
		suicides: make([]common.Address, 0),
	}
}

// `GetContext` returns the context of the StateDB.
func (sdb *StateDB) GetContext() context.Context {
	return sdb.ctx
}

// =============================================================================
// Transaction Handling
// =============================================================================

// `Prepare` sets the current transaction hash and index which are
// used when the EVM emits new state logs.
func (sdb *StateDB) Prepare(txHash common.Hash, ti uint) {
	sdb.lp.Prepare(txHash, ti)
}

// `Reset` resets the state object to the initial state.
func (sdb *StateDB) Reset(ctx context.Context) {
	sdb.ctx = ctx
}

// =============================================================================
// Account
// =============================================================================

// `CreateAccount` creates a new account.
func (sdb *StateDB) CreateAccount(addr common.Address) {
	sdb.ap.CreateAccount(sdb.ctx, addr)
}

// GetNonce implements the `GethStateDB` interface by returning the nonce
// of an account.
func (sdb *StateDB) GetNonce(addr common.Address) uint64 {
	return sdb.ap.GetNonce(sdb.ctx, addr)
}

// SetNonce implements the `GethStateDB` interface by setting the nonce
// of an account.
func (sdb *StateDB) SetNonce(addr common.Address, nonce uint64) {
	sdb.ap.SetNonce(sdb.ctx, addr, nonce)
}

// =============================================================================
// Balance
// =============================================================================

// `GetBalance` returns the balance of an account.
func (sdb *StateDB) GetBalance(addr common.Address) *big.Int {
	return sdb.bp.GetBalance(sdb.ctx, addr)
}

// `SubBalance` subtracts an amount from the balance of an account.
func (sdb *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	sdb.bp.SubBalance(sdb.ctx, addr, amount)
}

// `AddBalance` adds an amount to the balance of an account.
func (sdb *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	sdb.bp.AddBalance(sdb.ctx, addr, amount)
}

// `TransferBalance` transfers an amount from one account to another.
func (sdb *StateDB) TransferBalance(sender, receipient common.Address, amount *big.Int) {
	sdb.bp.TransferBalance(sdb.ctx, sender, receipient, amount)
}

// =============================================================================
// Code
// =============================================================================

func (sdb *StateDB) GetCode(addr common.Address) []byte {
	return sdb.cp.GetCodeFromHash(sdb.ctx, sdb.GetCodeHash(addr))
}

func (sdb *StateDB) SetCode(addr common.Address, code []byte) {
	codeHash := crypto.Keccak256Hash(code)
	// store or delete code
	// todo: nil vs 0? codehash?
	sdb.cp.SetCodeHash(sdb.ctx, addr, codeHash)
	if len(code) == 0 {
		sdb.cp.DeleteCode(sdb.ctx, addr)
	} else {
		sdb.cp.SetCode(sdb.ctx, codeHash, code)
	}
}

func (sdb *StateDB) GetCodeSize(addr common.Address) int {
	return len(sdb.GetCode(addr))
}

func (sdb *StateDB) GetCodeHash(addr common.Address) common.Hash {
	return sdb.cp.GetCodeHash(sdb.ctx, addr)
}

// =============================================================================
// Refund
// =============================================================================

// `AddRefund` implements the `GethStateDB` interface by adding gas to the
// refund counter.
func (sdb *StateDB) AddRefund(gas uint64) {
	sdb.rf.Add(gas)
}

// `SubRefund` implements the `GethStateDB` interface by subtracting gas from the
// refund counter. If the gas is greater than the refund counter, it will panic.
func (sdb *StateDB) SubRefund(gas uint64) {
	sdb.rf.Sub(gas)
}

// `GetRefund` implements the `GethStateDB` interface by returning the current
// value of the refund counter.
func (sdb *StateDB) GetRefund() uint64 {
	return sdb.rf.Get()
}

// =============================================================================
// Storage
// =============================================================================

func (sdb *StateDB) GetState(addr common.Address, key common.Hash) common.Hash {
	return sdb.sp.GetState(sdb.ctx, addr, key)
}

func (sdb *StateDB) GetCommittedState(addr common.Address, key common.Hash) common.Hash {
	return sdb.sp.GetCommittedState(sdb.ctx, addr, key)
}

func (sdb *StateDB) SetState(addr common.Address, key, value common.Hash) {
	// If empty value is given, delete the state entry.
	if len(value) == 0 || (value == common.Hash{}) {
		sdb.sp.SetState(sdb.ctx, addr, key, value)
		return
	}

	// Set the state entry.
	sdb.sp.DeleteState(sdb.ctx, addr, key)
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
	sdb.SubBalance(addr, sdb.GetBalance(addr))

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
	return sdb.ap.HasAccount(sdb.ctx, addr)
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
// Logs
// =============================================================================

// `AddLog` implements the `GethStateDB` interface by adding a log to the current
// transaction.
func (sdb *StateDB) AddLog(log *coretypes.Log) {
	sdb.lp.AddLog(log)
}

// `GetLogs` implements the `GethStateDB` interface by returning the logs for the.
func (sdb *StateDB) GetLogs(txHash common.Hash, blockHash common.Hash) []*coretypes.Log {
	return sdb.lp.GetLogs(txHash, blockHash)
}

// =============================================================================
// Snapshot
// =============================================================================

// `RevertToSnapshot` implements `StateDB`.
func (sdb *StateDB) RevertToSnapshot(id int) {
	// revert and discard all journal entries after snapshot id
	sdb.ctrl.RevertToSnapshot(id)
}

// `Snapshot` implements `StateDB`.
func (sdb *StateDB) Snapshot() int {
	return sdb.ctrl.Snapshot()
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
	return sdb.sp.ForEachStorage(sdb.ctx, addr, cb)
}

// `FinalizeTx` is called when we are complete with the state transition and want to commit the changes
// to the underlying store.
func (sdb *StateDB) FinalizeTx() error {
	// Manually delete all suicidal accounts.
	for _, suicidalAddr := range sdb.suicides {
		if !sdb.ap.HasAccount(sdb.ctx, suicidalAddr) {
			// handles the double suicide case
			continue
		}

		// clear storage
		_ = sdb.ForEachStorage(suicidalAddr,
			func(key, _ common.Hash) bool {
				sdb.SetState(suicidalAddr, key, common.Hash{})
				return true
			})

		// clear the codehash from this account
		sdb.cp.SetCodeHash(sdb.ctx, suicidalAddr, common.Hash{})

		// remove auth account
		sdb.ap.DeleteAccount(sdb.ctx, suicidalAddr)
	}
	sdb.ctrl.Finalize()
	return nil
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
