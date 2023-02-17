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

package mock

import (
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
			panic("mock out the GetCodeSize method")
		},
		GetCommittedStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			panic("mock out the GetCommittedState method")
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
		},
	}
	return mockedStargazerStateDB
}
