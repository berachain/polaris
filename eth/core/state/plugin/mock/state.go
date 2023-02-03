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
	"math/big"

	"github.com/berachain/stargazer/lib/crypto"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate moq -out ./state.mock.go -pkg mock ../../ StatePlugin

var Accounts map[common.Address]*Account

type Account struct {
	Balance  *big.Int
	Code     []byte
	CodeHash common.Hash
}

// `NewEmptyStatePlugin` returns an empty `StatePluginMock`.
func NewEmptyStatePlugin() *StatePluginMock {
	Accounts = make(map[common.Address]*Account)
	return &StatePluginMock{
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
			panic("mock out the TransferBalance method")
		},
	}
}
