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

	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to enforce `PrecompileManager` implements `PrecompileManagerI`.
var _ PrecompileManagerI = (*PrecompileManager)(nil)

// `PrecompileManager` is a struct that holds the registry of all precompiled contracts and a
// container that can run any precompiled contract.
type PrecompileManager struct {
	// `registry` is used to get an existing precompiled contract.
	registry *PrecompileRegistry
}

// `NewPrecompileManager` creates and returns a new `PrecompileManager` with the given context
// `ctx` and `registry`.
func NewPrecompileManager(registry *PrecompileRegistry) *PrecompileManager {
	return &PrecompileManager{
		registry: registry,
	}
}

// `Exists` returns a precompiled contract if it is found at the given `addr`.
func (pm *PrecompileManager) Exists(addr common.Address) (PrecompiledContract, bool) {
	// try stateless precompiles
	if pc, found := pm.registry.GetStatelessContract(addr); found {
		return pc, found
	}

	// try stateful precompiles
	if spc, found := pm.registry.GetStatefulContract(addr); found {
		return spc, found
	}

	// no precompile at given address found
	return nil, false
}

// `Run` uses the precompile container to execute the logic of a precompile contract function and
// consume gas appropriately.
func (pm *PrecompileManager) Run(
	sdb StateDB, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// TODO: implement once precompile container implemented.
	panic("implement me")
}
