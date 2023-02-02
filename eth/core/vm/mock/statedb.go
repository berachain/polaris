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

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate moq -out ./statedb.mock.go -pkg mock ../ StargazerStateDB

// `NewEmptyStateDB` creates a new `StateDBMock` instance.
func NewEmptyStateDB() *StargazerStateDBMock {
	mockedStargazerStateDB := &StargazerStateDBMock{
		AddAddressToAccessListFunc: func(addr common.Address) {
			panic("mock out the AddAddressToAccessList method")
		},
		AddBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			panic("mock out the AddBalance method")
		},
		AddLogFunc: func(log *types.Log) {
			panic("mock out the AddLog method")
		},
		AddPreimageFunc: func(hash common.Hash, bytes []byte) {
			panic("mock out the AddPreimage method")
		},
		AddRefundFunc: func(v uint64) {
			panic("mock out the AddRefund method")
		},
		AddSlotToAccessListFunc: func(addr common.Address, slot common.Hash) {
			panic("mock out the AddSlotToAccessList method")
		},
		AddressInAccessListFunc: func(addr common.Address) bool {
			panic("mock out the AddressInAccessList method")
		},
		CreateAccountFunc: func(address common.Address) {
			panic("mock out the CreateAccount method")
		},
		EmptyFunc: func(address common.Address) bool {
			return true
		},
		ExistFunc: func(address common.Address) bool {
			return false
		},
		FinalizeFunc: func() {
			panic("mock out the Finalize method")
		},
		ForEachStorageFunc: func(address common.Address, fn func(common.Hash, common.Hash) bool) error {
			panic("mock out the ForEachStorage method")
		},
		GetBalanceFunc: func(address common.Address) *big.Int {
			return big.NewInt(0)
		},
		GetCodeFunc: func(address common.Address) []byte {
			return []byte{}
		},
		GetCodeHashFunc: func(address common.Address) common.Hash {
			return common.Hash{}
		},
		GetCodeSizeFunc: func(address common.Address) int {
			panic("mock out the GetCodeSize method")
		},
		GetCommittedStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			panic("mock out the GetCommittedState method")
		},
		GetContextFunc: func() context.Context {
			panic("mock out the GetContext method")
		},
		BuildLogsAndClearFunc: func(common.Hash, common.Hash, uint, uint) []*types.Log {
			panic("mock out the GetLogs method")
		},
		GetNonceFunc: func(address common.Address) uint64 {
			return 0
		},
		GetRefundFunc: func() uint64 {
			return 0
		},
		GetStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			panic("mock out the GetState method")
		},
		HasSuicidedFunc: func(address common.Address) bool {
			panic("mock out the HasSuicided method")
		},
		PrepareAccessListFunc: func(sender common.Address, dest *common.Address,
			precompiles []common.Address, txAccesses types.AccessList) {
			panic("mock out the PrepareAccessList method")
		},
		RevertToSnapshotFunc: func(n int) {
			panic("mock out the RevertToSnapshot method")
		},
		SetCodeFunc: func(address common.Address, bytes []byte) {
			panic("mock out the SetCode method")
		},
		SetNonceFunc: func(address common.Address, v uint64) {},
		SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash) {
			panic("mock out the SetState method")
		},
		SlotInAccessListFunc: func(addr common.Address, slot common.Hash) (bool, bool) {
			panic("mock out the SlotInAccessList method")
		},
		SnapshotFunc: func() int {
			panic("mock out the Snapshot method")
		},
		SubBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			panic("mock out the SubBalance method")
		},
		SubRefundFunc: func(v uint64) {
			panic("mock out the SubRefund method")
		},
		SuicideFunc: func(address common.Address) bool {
			panic("mock out the Suicide method")
		},
		TransferBalanceFunc: func(address1 common.Address, address2 common.Address, intMoqParam *big.Int) {
			panic("mock out the TransferBalance method")
		},
	}
	return mockedStargazerStateDB
}
