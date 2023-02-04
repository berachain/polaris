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

	stargazercorevm "github.com/berachain/stargazer/eth/core/vm"
	"github.com/ethereum/go-ethereum/common"
	ethereumcorevm "github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

//go:generate moq -out ./evm.mock.go -pkg mock ../ StargazerEVM

func NewStargazerEVM() *StargazerEVMMock {
	mockedStargazerEVM := &StargazerEVMMock{
		CallFunc: func(caller ethereumcorevm.ContractRef, addr common.Address,
			input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
			return []byte{}, 0, nil
		},
		ChainConfigFunc: func() *params.ChainConfig {
			return &params.ChainConfig{
				LondonBlock:    big.NewInt(0),
				HomesteadBlock: big.NewInt(0),
			}
		},
		ContextFunc: func() ethereumcorevm.BlockContext {
			return stargazercorevm.BlockContext{
				CanTransfer: func(db stargazercorevm.GethStateDB, addr common.Address, amount *big.Int) bool {
					return true
				},
				BlockNumber: big.NewInt(1), // default to block == 1 to pass all forks,
			}
		},
		CreateFunc: func(caller ethereumcorevm.ContractRef, code []byte,
			gas uint64, value *big.Int) ([]byte, common.Address, uint64, error) {
			return []byte{}, common.Address{}, 0, nil
		},
		ResetFunc: func(txCtx ethereumcorevm.TxContext, sdb ethereumcorevm.StateDB) {
			panic("mock out the Reset method")
		},
		SetDebugFunc: func(debug bool) {
			// no-op
		},
		SetTracerFunc: func(tracer ethereumcorevm.EVMLogger) {
			// no-op
		},
		SetTxContextFunc: func(txCtx ethereumcorevm.TxContext) {
			// no-op
		},
		StateDBFunc: func() stargazercorevm.StargazerStateDB {
			return NewEmptyStateDB()
		},
		TracerFunc: func() ethereumcorevm.EVMLogger {
			return &EVMLoggerMock{}
		},
		TxContextFunc: func() ethereumcorevm.TxContext {
			panic("mock out the TxContext method")
		},
	}
	return mockedStargazerEVM
}
