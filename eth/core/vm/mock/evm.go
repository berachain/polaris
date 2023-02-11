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

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/lib/common"
)

//go:generate moq -out ./evm.mock.go -pkg mock ../ StargazerEVM

func NewStargazerEVM() *StargazerEVMMock {
	mockedStargazerEVM := &StargazerEVMMock{
		CallFunc: func(caller vm.ContractRef, addr common.Address,
			input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
			return []byte{}, 0, nil
		},
		ChainConfigFunc: func() *params.ChainConfig {
			return &params.ChainConfig{
				LondonBlock:    big.NewInt(0),
				HomesteadBlock: big.NewInt(0),
			}
		},
		ConfigFunc: func() vm.Config {
			return vm.Config{}
		},
		ContextFunc: func() vm.BlockContext {
			return vm.BlockContext{
				CanTransfer: func(db vm.GethStateDB, addr common.Address, amount *big.Int) bool {
					return true
				},
				BlockNumber: big.NewInt(1), // default to block == 1 to pass all forks,
			}
		},
		CreateFunc: func(caller vm.ContractRef, code []byte,
			gas uint64, value *big.Int) ([]byte, common.Address, uint64, error) {
			return []byte{}, common.Address{}, 0, nil
		},
		SetTxContextFunc: func(txCtx vm.TxContext) {
			// no-op
		},
		StateDBFunc: func() vm.StargazerStateDB {
			return NewEmptyStateDB()
		},
	}
	return mockedStargazerEVM
}
