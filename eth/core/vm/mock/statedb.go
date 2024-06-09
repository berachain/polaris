// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

//go:generate moq -out ./statedb.mock.go -pkg mock ../ PolarStateDB

// NewEmptyStateDB creates a new `StateDBMock` instance.
func NewEmptyStateDB() *PolarStateDBMock {
	mockedPolarStateDB := &PolarStateDBMock{
		AddAddressToAccessListFunc: func(addr common.Address) {

		},
		AddBalanceFunc: func(address common.Address, intMoqParam *big.Int) {

		},
		AddLogFunc: func(log *ethtypes.Log) {

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
		GetNonceFunc: func(address common.Address) uint64 {
			return 0
		},
		GetRefundFunc: func() uint64 {
			return 0
		},
		GetStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			return common.Hash{}
		},
		HasSelfDestructedFunc: func(address common.Address) bool {
			return false
		},
		PrepareFunc: func(rules params.Rules, sender common.Address,
			coinbase common.Address, dest *common.Address,
			precompiles []common.Address, txAccesses ethtypes.AccessList,
		) {
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
		SelfDestructFunc: func(address common.Address) {
		},
	}
	return mockedPolarStateDB
}
