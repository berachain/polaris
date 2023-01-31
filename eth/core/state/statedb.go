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

	"github.com/berachain/stargazer/eth/core/state/plugin"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
	"github.com/berachain/stargazer/lib/snapshot"
)

// `StatePlugin` is the plugin that holds the state of the evm.

var (
	// EmptyCodeHash is the Keccak256 Hash of empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash = crypto.Keccak256Hash(nil)
)

// Compile-time assertion that StateDB implements the vm.StargazerStateDB interface.
var _ vm.StargazerStateDB = (*StateDB)(nil)

type StateDB struct { //nolint:revive // we live the vibe.
	// The controller is used to manage the plugins
	ctrl snapshot.Controller

	// References to the plugins in the controller.
	StatePlugin
	RefundPlugin
	LogsPlugin

	// Dirty tracking of suicided accounts, we have to keep track of these manually, in order
	// for the code and state to still be accessible even after the account has been deleted.
	// We chose to keep track of them in a separate slice, rather than a map, because the
	// number of accounts that will be suicided in a single transaction is expected to be
	// very low.
	suicides []common.Address
}

func NewStateDB(sp StatePlugin, ctrl snapshot.Controller) *StateDB {
	// Build snapshottables for logs and refunds.
	lp := plugin.NewLogs()
	rf := plugin.NewRefund()

	// Register the snapshottables with the controller.
	_ = ctrl.Control(plugin.RefundName, lp)
	_ = ctrl.Control(plugin.LogsName, rf)
	_ = ctrl.Control(plugin.BackendName, sp)

	// Create the `StateDB` and populate the developer provided plugins.
	return &StateDB{
		StatePlugin:  sp,
		LogsPlugin:   lp,
		RefundPlugin: rf,
		ctrl:         ctrl,
		suicides:     make([]common.Address, 1), // very rare to suicide, so we alloc 1 slot.
	}
}

// =============================================================================
// Transaction Handling
// =============================================================================

// // `Prepare` sets the current transaction hash and index which are
// // used when the EVM emits new state logs.
// func (sdb *StateDB) Prepare(txHash common.Hash, ti uint) {
// 	sdb.Prepare(txHash, ti)
// }

// `Reset` resets the state object to the initial state.
func (sdb *StateDB) Reset(ctx context.Context) {
	// sdb.ctx = ctx
}

// =============================================================================
// Suicide
// =============================================================================

// Suicide implements the StargazerStateDB interface by marking the given address as suicided.
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

// `HasSuicided` implements the `StargazerStateDB` interface by returning if the contract was suicided
// in current transaction.
func (sdb *StateDB) HasSuicided(addr common.Address) bool {
	for _, suicide := range sdb.suicides {
		if bytes.Equal(suicide[:], addr[:]) {
			return true
		}
	}
	return false
}

// `Empty` implements the `StargazerStateDB` interface by returning whether the state object
// is either non-existent or empty according to the EIP161 epecification
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
	sdb.ctrl.RevertToSnapshot(id)
}

// `Snapshot` implements `StateDB`.
func (sdb *StateDB) Snapshot() int {
	return sdb.ctrl.Snapshot()
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
