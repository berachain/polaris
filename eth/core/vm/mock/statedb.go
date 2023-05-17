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

	"github.com/ethereum/go-ethereum/params"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"
)

//go:generate moq -out ./statedb.mock.go -pkg mock ../ PolarisStateDB

// NewEmptyStateDB creates a new `StateDBMock` instance.
func NewEmptyStateDB() *PolarisStateDBMock {
	mockedPolarisStateDB := &PolarisStateDBMock{
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
		GetBlockLogsAndClearFunc: func(common.Hash) []*types.Log {
			return nil
		},
		CreateAccountFunc: func(address common.Address) {

		},
		EmptyFunc: func(address common.Address) bool {
			return true
		},
		ErrorFunc: func() error {
			return nil
		},
		ExistFunc: func(address common.Address) bool {
			return false
		},
		FinaliseFunc: func(bool) {
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
		GetLogsFunc: func(hash common.Hash, blockNumber uint64, blockHash common.Hash) []*types.Log {
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
		PrepareFunc: func(rules params.Rules, sender common.Address,
			coinbase common.Address, dest *common.Address,
			precompiles []common.Address, txAccesses types.AccessList,
		) {
			// no-op
		},
		ResetFunc: func(common.Hash, int) {
			// no-op
		},
		RevertToSnapshotFunc: func(n int) {

		},
		SetCodeFunc: func(address common.Address, bytes []byte) {

		},
		SetNonceFunc: func(address common.Address, v uint64) {},
		SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash) {

		},
		SetTxContextFunc: func(thash common.Hash, ti int) {},
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
	}
	return mockedPolarisStateDB
}
