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
	"github.com/berachain/stargazer/eth/params"
)

// `EVMFactory` is used to build new Stargazer `EVM`s.
type EVMFactory struct {
	// `precompileManager` is responsible for keeping track of the stateful precompile
	// containers that are available to the EVM and executing them.
	precompileManager PrecompileManager
}

// `NewEVMFactory` creates and returns a new `EVMFactory` with the given `precompileManager`.
func NewEVMFactory(precompileManager PrecompileManager) *EVMFactory {
	return &EVMFactory{
		precompileManager: precompileManager,
	}
}

// `Build` creates and returns a new `vm.StargazerEVM`.
func (ef *EVMFactory) Build(
	ssdb StargazerStateDB,
	blockCtx BlockContext,
	chainConfig *params.EthChainConfig,
	noBaseFee bool,
) StargazerEVM {
	return NewStargazerEVM(
		blockCtx, TxContext{}, ssdb, chainConfig, Config{}, ef.precompileManager,
	)
}
