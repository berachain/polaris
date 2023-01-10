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
	"math/big"

	"github.com/berachain/stargazer/common"
	"github.com/holiman/uint256"
)

type StargazerEVM interface { //nolint:revive
	Config() Config
	Context() BlockContext
	TxContext() TxContext

	Reset(txCtx TxContext, statedb StateDB)
	Cancel()
	Cancelled() bool //nolint: misspell
	Interpreter() *EVMInterpreter
	Call(caller ContractRef, addr common.Address,
		input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	CallCode(caller ContractRef, addr common.Address,
		input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	DelegateCall(caller ContractRef, addr common.Address,
		input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error,
	)
	StaticCall(caller ContractRef, addr common.Address,
		input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error,
	)
	Create(caller ContractRef, code []byte, gas uint64, value *big.Int) (
		ret []byte, contractAddr common.Address, leftOverGas uint64, err error,
	)
	Create2(
		caller ContractRef,
		code []byte,
		gas uint64,
		endowment *big.Int,
		salt *uint256.Int) (
		ret []byte, contractAddr common.Address, leftOverGas uint64, err error,
	)
	// ChainConfig() *params.EthChainConfig

	// ActivePrecompiles(rules params.Rules) []common.Address
	Precompile(addr common.Address) (PrecompiledContract, bool)
	RunPrecompiledContract(
		p StatefulContract,
		addr common.Address,
		input []byte,
		suppliedGas uint64,
		value *big.Int) (
		ret []byte, remainingGas uint64, err error,
	)

	SetTracer(tracer EVMLogger)
	SetDebug(debug bool)
	WithTxContext(txCtx TxContext) *StargazerEVM
	WithBlockContext(blockCtx BlockContext) *StargazerEVM
	// StateDB() state.ExtStateDB
	WithStateDB(stateDB StateDB) *StargazerEVM
}
