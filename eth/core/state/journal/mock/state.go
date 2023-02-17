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

	"github.com/berachain/stargazer/eth/crypto"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate moq -out ./state.mock.go -pkg mock ../../ Plugin

var Accounts map[common.Address]*Account

type Account struct {
	Balance  *big.Int
	Code     []byte
	CodeHash common.Hash
}

// `NewEmptyStatePlugin` returns an empty `StatePluginMock`.
func NewEmptyStatePlugin() *PluginMock {
	Accounts = make(map[common.Address]*Account)
	return &PluginMock{
		AddBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			Accounts[address] = &Account{
				Balance:  Accounts[address].Balance.Add(Accounts[address].Balance, intMoqParam),
				Code:     Accounts[address].Code,
				CodeHash: Accounts[address].CodeHash,
			}
		},
		CreateAccountFunc: func(address common.Address) {
			Accounts[address] = &Account{
				Balance:  big.NewInt(0),
				CodeHash: crypto.Keccak256Hash(nil),
			}
		},
		DeleteSuicidesFunc: func(addresss []common.Address) {
			for _, addr := range addresss {
				delete(Accounts, addr)
			}
		},
		ExistFunc: func(address common.Address) bool {
			panic("mock out the Exist method")
		},
		FinalizeFunc: func() {
			// no-op
		},
		ForEachStorageFunc: func(address common.Address, fn func(common.Hash, common.Hash) bool) error {
			panic("mock out the ForEachStorage method")
		},
		GetBalanceFunc: func(address common.Address) *big.Int {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			return Accounts[address].Balance
		},
		GetCodeFunc: func(address common.Address) []byte {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			return Accounts[address].Code
		},
		GetCodeHashFunc: func(address common.Address) common.Hash {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			return Accounts[address].CodeHash
		},
		GetCodeSizeFunc: func(address common.Address) int {
			panic("mock out the GetCodeSize method")
		},
		GetCommittedStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			panic("mock out the GetCommittedState method")
		},
		GetNonceFunc: func(address common.Address) uint64 {
			return 0
		},
		GetStateFunc: func(address common.Address, hash common.Hash) common.Hash {
			panic("mock out the GetState method")
		},
		RegistryKeyFunc: func() string {
			return "mockstate"
		},
		RevertToSnapshotFunc: func(n int) {
			// no-op
		},
		SetCodeFunc: func(address common.Address, bytes []byte) {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			Accounts[address] = &Account{
				Balance:  Accounts[address].Balance,
				Code:     bytes,
				CodeHash: common.BytesToHash(bytes),
			}
		},
		SetNonceFunc: func(address common.Address, v uint64) {
			panic("mock out the SetNonce method")
		},
		SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash) {
			panic("mock out the SetState method")
		},
		SnapshotFunc: func() int {
			return 0
		},
		SubBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesnt exist")
			}
			Accounts[address] = &Account{
				Balance: Accounts[address].Balance.Sub(Accounts[address].Balance, intMoqParam),
			}
		},
		TransferBalanceFunc: func(address1 common.Address, address2 common.Address, intMoqParam *big.Int) {
			// no-op
		},
	}
}
