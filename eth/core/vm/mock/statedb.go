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

		},
		AddBalanceFunc: func(address common.Address, intMoqParam *big.Int) {

		},
		AddLogFunc: func(log *types.Log) {

		},
		AddPreimageFunc: func(hash common.Hash, bytes []byte) {

		},
		AddRefundFunc: func(v uint64) {

		},
		AddSlotToAccessListFunc: func(addr common.Address, slot common.Hash) {

		},
		AddressInAccessListFunc: func(addr common.Address) bool {
			return false
		},
		CreateAccountFunc: func(address common.Address) {

		},
		EmptyFunc: func(address common.Address) bool {
			return true
		},
		ExistFunc: func(address common.Address) bool {
			return false
		},
		FinalizeFunc: func() {
			// no-op
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
			return 0
		},
		GetCommittedStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			return common.Hash{}
		},
		BuildLogsAndClearFunc: func(common.Hash, common.Hash, uint, uint) []*types.Log {
			return []*types.Log{}
		},
		GetNonceFunc: func(address common.Address) uint64 {
			return 0
		},
		GetRefundFunc: func() uint64 {
			return 0
		},
		GetStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			return common.Hash{}
		},
		HasSuicidedFunc: func(address common.Address) bool {
			return false
		},
		PrepareAccessListFunc: func(sender common.Address, dest *common.Address,
			precompiles []common.Address, txAccesses types.AccessList) {

		},
		ResetFunc: func(contextMoqParam context.Context) {
			// no-op
		},
		RevertToSnapshotFunc: func(n int) {

		},
		SetCodeFunc: func(address common.Address, bytes []byte) {

		},
		SetNonceFunc: func(address common.Address, v uint64) {},
		SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash) {

		},
		SlotInAccessListFunc: func(addr common.Address, slot common.Hash) (bool, bool) {
			return false, false
		},
		SnapshotFunc: func() int {
			return 0
		},
		SubBalanceFunc: func(address common.Address, intMoqParam *big.Int) {

		},
		SubRefundFunc: func(v uint64) {

		},
		SuicideFunc: func(address common.Address) bool {
			return false
		},
		TransferBalanceFunc: func(address1 common.Address, address2 common.Address, intMoqParam *big.Int) {
		},
	}
	return mockedStargazerStateDB
}
