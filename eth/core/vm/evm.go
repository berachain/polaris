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

package vm

import (
	"context"

	"github.com/berachain/stargazer/eth/params"
)

// Compile-time assertion to ensure `StargazerEVM` implements `VMInterface`.
var _ StargazerEVM = (*stargazerEVM)(nil)

// `StargazerEVM` is the wrapper for the Go-Ethereum EVM.
type stargazerEVM struct {
	*GethEVM
}

// `NewStargazerEVM` creates and returns a new `StargazerEVM`.
func NewStargazerEVM(
	blockCtx BlockContext,
	txCtx TxContext,
	stateDB StargazerStateDB,
	chainConfig *params.EthChainConfig,
	config Config,
	pcmgr PrecompileManager,
) StargazerEVM {
	return &stargazerEVM{
		GethEVM: NewGethEVMWithPrecompiles(
			blockCtx, txCtx, stateDB, chainConfig, config, pcmgr,
		),
	}
}

// MUST BE CALLED EVERY TIME THE EVM IS REUSED, I.E. BEFORE EVERY TRANSACTION.
func (evm *stargazerEVM) ResetPrecompileManager(ctx context.Context) {
	_ = evm.PrecompileManager.Reset(ctx)
}

func (evm *stargazerEVM) SetTxContext(txCtx TxContext) {
	evm.GethEVM.TxContext = txCtx
}

func (evm *stargazerEVM) Context() BlockContext {
	return evm.GethEVM.Context
}

func (evm *stargazerEVM) StateDB() StargazerStateDB {
	return evm.GethEVM.StateDB.(StargazerStateDB)
}

func (evm *stargazerEVM) SetTracer(tracer EVMLogger) {
	evm.Config.Tracer = tracer
}

func (evm *stargazerEVM) SetDebug(debug bool) {
	evm.Config.Debug = debug
}

func (evm *stargazerEVM) Tracer() EVMLogger {
	return evm.Config.Tracer
}

func (evm *stargazerEVM) TxContext() TxContext {
	return evm.GethEVM.TxContext
}
