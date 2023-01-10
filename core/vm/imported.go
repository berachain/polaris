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

import gevm "github.com/ethereum/go-ethereum/core/vm"

type (
	AccountRef          = gevm.AccountRef
	BlockContext        = gevm.BlockContext
	Config              = gevm.Config
	ContractRef         = gevm.ContractRef
	EVMInterpreter      = gevm.EVMInterpreter
	EVMLogger           = gevm.EVMLogger
	GetHashFunc         = gevm.GetHashFunc
	GethEVM             = gevm.EVM
	PrecompiledContract = gevm.PrecompiledContract
	StateDB             = gevm.StateDB
	TxContext           = gevm.TxContext
)

var (
	ErrExecutionReverted              = gevm.ErrExecutionReverted
	ErrOutOfGas                       = gevm.ErrOutOfGas
	EthActivePrecompiles              = gevm.ActivePrecompiles
	EthRunStatefulPrecompiledContract = gevm.RunStatefulPrecompiledContract
	NewGethEVM                        = gevm.NewEVM
	PrecompiledAddressesBerlin        = gevm.PrecompiledAddressesBerlin
	PrecompiledContractsBerlin        = gevm.PrecompiledContractsBerlin
	PrecompiledContractsByzantium     = gevm.PrecompiledContractsByzantium
	PrecompiledContractsHomestead     = gevm.PrecompiledContractsHomestead
	PrecompiledContractsIstanbul      = gevm.PrecompiledContractsIstanbul
)
