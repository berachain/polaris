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
	"math/big"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/journal"
	"github.com/berachain/stargazer/core/state/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StateReader interface { //nolint:revive // we like the vibe.
	GetBalance(addr common.Address) *big.Int
	GetCode(addr common.Address) []byte
	GetCodeSize(addr common.Address) int
	GetCodeHash(addr common.Address) common.Hash
	GetNonce(addr common.Address) uint64
	GetState(addr common.Address, hash common.Hash) common.Hash
	GetCommittedState(addr common.Address, hash common.Hash) common.Hash
	Exist(addr common.Address) bool
	Empty(addr common.Address) bool
}

var _ StateReader = (*StateDBReader)(nil)

type StateDBReader struct { //nolint:revive // we like the vibe.

	// keepers used for balance and account information
	ak AccountKeeper
	bk BankKeeper

	// the evm denom
	evmDenom string

	// the live context
	liveCtx sdk.Context

	// eth state stores: required for vm.StateDB
	// We store references to these stores, so that we can access them
	// directly, without having to go through the MultiStore interface.
	liveEthState cachekv.StateDBCacheKVStore
}

func NewStateDBReader(
	ctx sdk.Context,
	ak AccountKeeper,
	bk BankKeeper,
	storeKey storetypes.StoreKey,
	evmDenom string,
) *StateDBReader {
	return &StateDBReader{
		ak:           ak,
		bk:           bk,
		evmDenom:     evmDenom,
		liveCtx:      ctx,
		liveEthState: cachekv.NewEvmStore(ctx.KVStore(storeKey), journal.NewManager()),
	}
}

// ==============================================================================
// Balance
// ==============================================================================

func (sdbr *StateDBReader) GetBalance(addr common.Address) *big.Int {
	// Note: bank keeper will return 0 if account/state_object is not found
	return sdbr.bk.GetBalance(sdbr.liveCtx, addr[:], sdbr.evmDenom).Amount.BigInt()
}

// ==============================================================================
// Code
// ==============================================================================

// GetCodeHash implements the GethStateDB interface by returning
// the code hash of account.
func (sdbr *StateDBReader) GetCodeHash(addr common.Address) common.Hash {
	if sdbr.ak.HasAccount(sdbr.liveCtx, addr[:]) {
		if ch := prefix.NewStore(sdbr.liveEthState, types.KeyPrefixCodeHash).
			Get(addr[:]); ch != nil {
			return common.BytesToHash(ch)
		}
		return types.EmptyCodeHash
	}
	// if account at addr does not exist, return ZeroCodeHash
	return types.ZeroCodeHash
}

// GetCode implements the GethStateDB interface by returning
// the code of account (nil if not exists).
func (sdbr *StateDBReader) GetCode(addr common.Address) []byte {
	codeHash := sdbr.GetCodeHash(addr)
	// if account at addr does not exist, GetCodeHash returns ZeroCodeHash so return nil
	// if codeHash is empty, i.e. crypto.Keccak256(nil), also return nil
	if codeHash == types.ZeroCodeHash || codeHash == types.EmptyCodeHash {
		return nil
	}
	return prefix.NewStore(sdbr.liveEthState, types.KeyPrefixCode).Get(codeHash.Bytes())
}

// GetCodeSize implements the GethStateDB interface by returning the size of the
// code associated with the given GethStateDB.
func (sdbr *StateDBReader) GetCodeSize(addr common.Address) int {
	return len(sdbr.GetCode(addr))
}

// ==============================================================================
// State
// ==============================================================================

// GetCommittedState implements the GethStateDB interface by returning the
// committed state of an address.
func (sdbr *StateDBReader) GetCommittedState(
	addr common.Address,
	hash common.Hash,
) common.Hash {
	if value := prefix.NewStore(sdbr.liveEthState.GetParent(),
		types.AddressStoragePrefix(addr)).Get(hash[:]); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// GetState implements the GethStateDB interface by returning the current state of an
// address.
func (sdbr *StateDBReader) GetState(addr common.Address, hash common.Hash) common.Hash {
	if value := prefix.NewStore(sdbr.liveEthState,
		types.AddressStoragePrefix(addr)).Get(hash[:]); value != nil {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

// ==============================================================================
// Account
// ==============================================================================

// Exist implements the GethStateDB interface by reporting whether the given account address
// exists in the state. Notably this also returns true for suicided accounts, which is accounted
// for since, `RemoveAccount()` is not called until Commit.
func (sdbr *StateDBReader) Exist(addr common.Address) bool {
	return sdbr.ak.HasAccount(sdbr.liveCtx, addr[:])
}

// GetNonce implements the GethStateDB interface by returning the nonce
// of an account.
func (sdbr *StateDBReader) GetNonce(addr common.Address) uint64 {
	acc := sdbr.ak.GetAccount(sdbr.liveCtx, addr[:])
	if acc == nil {
		return 0
	}
	return acc.GetSequence()
}

// Empty implements the GethStateDB interface by returning whether the state object
// is either non-existent or empty according to the EIP161 specification
// (balance = nonce = code = 0)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
func (sdbr *StateDBReader) Empty(addr common.Address) bool {
	ch := sdbr.GetCodeHash(addr)
	return sdbr.GetNonce(addr) == 0 &&
		(ch == types.EmptyCodeHash || ch == types.ZeroCodeHash) &&
		sdbr.GetBalance(addr).Sign() == 0
}
