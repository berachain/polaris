// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

package vm

import (
	"math/big"

	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/params"
)

type VMInterface interface { //nolint:revive // we like the vibe.
	Reset(txCtx TxContext, sdb state.BaseStateDB)
	Create(caller ContractRef, code []byte,
		gas uint64, value *big.Int,
	) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
	Call(caller ContractRef, addr common.Address, input []byte,
		gas uint64, value *big.Int,
	) (ret []byte, leftOverGas uint64, err error)
	Config() Config
	ChainConfig() *params.EthChainConfig
	Context() BlockContext
	StateDB() state.StargazerStateDB
	SetTxContext(TxContext)
	SetTracer(EVMLogger)
	SetDebug(bool)
}

var _ VMInterface = (*StargazerEVM)(nil)

// `StargazerEVM` is the wrapper for the Go-Ethereum EVM.
type StargazerEVM struct {
	*GethEVM
}

// `NewStargazerEVM` creates and returns a new `StargazerEVM`.
func NewStargazerEVM(
	blockCtx BlockContext,
	txCtx TxContext,
	stateDB state.BaseStateDB,
	chainConfig *params.EthChainConfig,
	config Config,
	precompileHost precompile.Host,
) *StargazerEVM {
	return &StargazerEVM{
		GethEVM: NewGethEVMWithPrecompileHost(blockCtx, txCtx, stateDB, chainConfig, config, precompileHost),
	}
}

func (sge *StargazerEVM) SetDebug(debug bool) {
	sge.GethEVM.Config.Debug = debug
}

func (sge *StargazerEVM) SetTracer(tracer EVMLogger) {
	sge.GethEVM.Config.Tracer = tracer
}

func (sge *StargazerEVM) SetTxContext(txCtx TxContext) {
	sge.GethEVM.TxContext = txCtx
}

func (sge *StargazerEVM) StateDB() state.StargazerStateDB {
	return sge.GethEVM.StateDB.(state.StargazerStateDB)
}

func (sge *StargazerEVM) Config() Config {
	return sge.GethEVM.Config
}

func (sge *StargazerEVM) Context() BlockContext {
	return sge.GethEVM.Context
}
