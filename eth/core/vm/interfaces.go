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
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/params"
	libtypes "github.com/berachain/stargazer/lib/types"

	"context"
	"math/big"

	"github.com/berachain/stargazer/lib/common"
)

type (
	// `StargazerEVM` defines an extension to the interface provided by Go-Ethereum to support additional
	// state transition functionalities.
	StargazerEVM interface {
		Reset(txCtx TxContext, sdb GethStateDB)
		Create(caller ContractRef, code []byte,
			gas uint64, value *big.Int,
		) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
		Call(caller ContractRef, addr common.Address, input []byte,
			gas uint64, value *big.Int,
		) (ret []byte, leftOverGas uint64, err error)

		SetTxContext(txCtx TxContext)
		SetTracer(tracer EVMLogger)
		SetDebug(debug bool)
		StateDB() StargazerStateDB
		TxContext() TxContext
		Tracer() EVMLogger
		Context() BlockContext
		ChainConfig() *params.EthChainConfig
	}

	// `StargazerStateDB` defines an extension to the interface provided by Go-Ethereum to support
	// additional state transition functionalities.
	StargazerStateDB interface {
		GethStateDB
		libtypes.Finalizeable

		// `GetContext` returns the Go context associated to the StateDB.
		GetContext() context.Context

		// TransferBalance transfers the balance from one account to another
		TransferBalance(common.Address, common.Address, *big.Int)

		// `BuildLogsAndClear` returns the logs for the tx.
		BuildLogsAndClear(common.Hash, common.Hash, uint, uint) []*coretypes.Log
	}

	// `PrecompileRunner` defines the required function of a vm-specific precompile runner.
	PrecompileRunner interface {
		// `Run` runs a precompile container with the given statedb and returns the remaining gas.
		Run(pc PrecompileContainer, ssdb StargazerStateDB, input []byte,
			caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
		) (ret []byte, remainingGas uint64, err error)
	}

	// `BasePrecompileImpl` is a type for the base precompile implementation, which only needs to
	// provide an Ethereum address of where its contract is found.
	BasePrecompileImpl = libtypes.Registrable[common.Address]
)
