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

	"github.com/berachain/polaris/eth/core/state"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate moq -out ./state.mock.go -skip-ensure -pkg mock ../ Plugin

var Accounts map[common.Address]*Account

type Account struct {
	Balance  *big.Int
	Code     []byte
	CodeHash common.Hash
	Nonce    uint64
}

// NewEmptyStatePlugin returns an empty `StatePluginMock`.
func NewEmptyStatePlugin() *PluginMock {
	Accounts = make(map[common.Address]*Account)
	return &PluginMock{
		AddBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			if _, ok := Accounts[address]; !ok {
				Accounts[address] = &Account{
					Balance: intMoqParam,
				}
			}
			Accounts[address] = &Account{
				Balance:  Accounts[address].Balance.Add(Accounts[address].Balance, intMoqParam),
				Code:     Accounts[address].Code,
				CodeHash: Accounts[address].CodeHash,
			}
		},
		CloneFunc: func() state.Plugin {
			return nil
		},
		CreateAccountFunc: func(address common.Address) {
			Accounts[address] = &Account{
				Balance:  new(big.Int),
				CodeHash: crypto.Keccak256Hash(nil),
			}
		},
		DeleteAccountsFunc: func(addresss []common.Address) {
			for _, addr := range addresss {
				delete(Accounts, addr)
			}
		},
		ExistFunc: func(address common.Address) bool {
			_, ok := Accounts[address]
			return ok
		},
		EmptyFunc: func(address common.Address) bool {
			panic("mock out the Empty method")
		},
		ErrorFunc: func() error {
			return nil
		},
		FinalizeFunc: func() {
			// no-op
		},
		ForEachStorageFunc: func(address common.Address, fn func(common.Hash, common.Hash) bool) error {
			panic("mock out the ForEachStorage method")
		},
		GetBalanceFunc: func(address common.Address) *big.Int {
			if _, ok := Accounts[address]; !ok {
				return new(big.Int).SetUint64(0)
			}
			return Accounts[address].Balance
		},
		GetCodeFunc: func(address common.Address) []byte {
			if _, ok := Accounts[address]; !ok {
				return nil
			}
			return Accounts[address].Code
		},
		GetCodeHashFunc: func(address common.Address) common.Hash {
			if _, ok := Accounts[address]; !ok {
				return common.Hash{}
			}
			return Accounts[address].CodeHash
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
				panic("acct doesn't exist")
			}
			Accounts[address] = &Account{
				Balance:  Accounts[address].Balance,
				Code:     bytes,
				CodeHash: common.BytesToHash(bytes),
			}
		},
		SetNonceFunc: func(address common.Address, v uint64) {
			Accounts[address].Nonce = v
		},
		SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash) {
			panic("mock out the SetState method")
		},
		SnapshotFunc: func() int {
			return 0
		},
		SubBalanceFunc: func(address common.Address, intMoqParam *big.Int) {
			if _, ok := Accounts[address]; !ok {
				panic("acct doesn't exist")
			}
			Accounts[address] = &Account{
				Balance:  Accounts[address].Balance.Sub(Accounts[address].Balance, intMoqParam),
				Code:     Accounts[address].Code,
				CodeHash: Accounts[address].CodeHash,
			}
		},
	}
}
