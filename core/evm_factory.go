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

package core

import (
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/params"
)

// `EVMFactory` is used to build new Stargazer `EVM`s.
type EVMFactory struct {
	// `pr` is the precompile registry is responsible for keeping track of the stateful
	// precompiles, events, and errors that are available to the EVM.
	pr *vm.PrecompileRegistry
}

// `NewEVMFactory` creates and returns a new `EVMFactory` with a new `vm.PrecompileRegistry`.
func NewEVMFactory() *EVMFactory {
	return &EVMFactory{
		pr: vm.NewPrecompileRegistry(),
	}
}

// `Build` creates and returns a new `vm.StargazerEVM` with a new `vm.PrecompileHost`.
func (ef *EVMFactory) Build(
	sdb vm.StargazerStateDB,
	blockCtx vm.BlockContext,
	txCtx vm.TxContext,
	chainConfig *params.EthChainConfig,
	noBaseFee bool,
) *vm.StargazerEVM {
	return vm.NewStargazerEVM(
		blockCtx, txCtx, sdb, chainConfig, vm.Config{}, vm.NewPrecompileHost(
			ef.pr,
			sdb,
		),
	)
}
