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

	"github.com/berachain/stargazer/lib/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type VMInterface interface { //nolint:revive // we like the vibe.
	Reset(txCtx TxContext, sdb StateDB)
	Create(caller ContractRef, code []byte,
		gas uint64, value *uint256.Int,
	) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
	Call(caller ContractRef, addr common.Address, input []byte,
		gas uint64, value *uint256.Int, bailout bool,
	) (ret []byte, leftOverGas uint64, err error)
	Config() Config
	ChainConfig() *params.ChainConfig
	ChainRules() *params.Rules
	Context() BlockContext
	StateDB() StateDB
	TxContext() TxContext
}

// `PrecompileManagerI` is invoked by the EVM to determine if a particular address is one of a
// precompiled contract and allows the EVM to execute a precompiled contract function.
type PrecompileManagerI interface {
	// `Exists` returns a precompiled contract if it was found at `addr`.
	Exists(addr common.Address) (PrecompiledContract, bool)

	// `Run` runs a precompiled contract and returns the leftover gas.
	Run(sdb StateDB, input []byte, caller common.Address,
		value *big.Int, suppliedGas uint64, readonly bool,
	) (ret []byte, leftOverGas uint64, err error)
}
