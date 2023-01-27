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
	"context"
	"math/big"

	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
)

// var _ vm.StargazerStateDB = (*StateDB)(nil)

type StateDB struct { //nolint:revive // StateDB is a struct that holds the state of the blockchain.
	ctx context.Context

	ctrl Controller

	// Developer provided plugins
	ap AccountPlugin
	bp BalancePlugin
	cp CodePlugin
	sp StoragePlugin

	// Internal plugins
	rf RefundPlugin

	// Transaction and logging bookkeeping
	txHash  common.Hash
	txIndex uint
	// logs    map[common.Hash]ds.Stack[*coretypes.Log]
	// logSize uint
}

func NewStateDB(ctrl Controller) *StateDB {
	ctrl.AddStore(&RefundPlugin{})
	return &StateDB{ctrl: ctrl}
}

// =============================================================================
// Transaction Handling
// =============================================================================

// `Prepare` sets the current transaction hash and index which are
// used when the EVM emits new state logs.
func (sdb *StateDB) Prepare(txHash common.Hash, ti uint) {
	sdb.txHash = txHash
	sdb.txIndex = ti
	// sdb.logs[txHash] = stack.New[*coretypes.Log](logStackCapacity)
}

// `Reset` resets the state object to the initial state.
func (sdb *StateDB) Reset(ctx context.Context) {
	sdb.ctx = ctx
}

// =============================================================================
// Account
// =============================================================================

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

func (sdb *StateDB) GetBalance(addr common.Address) *big.Int {
	return sdb.bp.GetBalance(sdb.ctx, addr)
}

func (sdb *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	sdb.bp.SubBalance(sdb.ctx, addr, amount)
}

func (sdb *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	sdb.bp.AddBalance(sdb.ctx, addr, amount)
}

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

func (sdb *StateDB) SetState(addr common.Address, key, value common.Hash) {
	// If empty value is given, delete the state entry.
	if len(value) == 0 || (value == common.Hash{}) {
		sdb.sp.DeleteState(sdb.ctx, addr, key)
		return
	}

	// Set the state entry.
	sdb.sp.SetState(sdb.ctx, addr, key, value)
}
