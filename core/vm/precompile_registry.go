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
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
)

// KeyPrefixPrecompileAddress is the prefix for the precompile address to name mapping in the
// precompile kvstore.
var KeyPrefixPrecompileAddress = []byte{0x01}

// `PrecompileRegistry` will store and provide stateless and custom or dynamic stateful
// precompiled contracts to the EVM.
type PrecompileRegistry struct {
	// `StatelessPrecompiles` is a map of Ethereum addresses to stateless precompiled contracts.
	statelessPrecompiles map[common.Address]PrecompiledContract

	// `statefulPrecompiles` is a map of Ethereum addresses to stateful precompiled contracts.
	statefulPrecompiles map[common.Address]StatefulPrecompiledContract

	// `dynamicPrecompiles` is a map of names to dynamic precompiled contracts.
	dynamicPrecompiles map[string]DynamicPrecompiledContract

	// `storeKey` is the store key of the Cosmos KVStore used for storing the addresses of dynamic
	// precompiled contracts.
	storeKey storetypes.StoreKey

	// `logFactory` is the Ethereum log builder for all Cosmos events emitted during precompile
	// execution.
	logFactory *precompile.LogFactory
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewPrecompileRegistry` creates and returns a new `PrecompileRegistry` for given `storeKey`.
func NewPrecompileRegistry(storeKey storetypes.StoreKey) *PrecompileRegistry {
	return &PrecompileRegistry{
		statelessPrecompiles: make(map[common.Address]PrecompiledContract),
		statefulPrecompiles:  make(map[common.Address]StatefulPrecompiledContract),
		dynamicPrecompiles:   make(map[string]DynamicPrecompiledContract),
		storeKey:             storeKey,
		logFactory:           precompile.NewLogFactory(),
	}
}

// ==============================================================================
// Setters
// ==============================================================================

// `RegisterStateless` stores the stateless precompile `pc` at the given Ethereum address `addr`.
func (pr *PrecompileRegistry) RegisterStatelessContract(
	addr common.Address,
	pc PrecompiledContract,
) error {
	if _, found := pr.statelessPrecompiles[addr]; found {
		return errors.Wrap(ErrPrecompileAlreadyRegistered, addr.String())
	}
	pr.statelessPrecompiles[addr] = pc
	return nil
}

// `RegisterModule` stores a module's stateful precompile contract (in memory) at a hardcoded
// address (module account address) and registers its events if it has any.
func (pr *PrecompileRegistry) RegisterModule(moduleName string, contract any) error {
	moduleEthAddr := common.BytesToAddress(authtypes.NewModuleAddress(moduleName).Bytes())

	// store the module stateful precompile contract in the hardcoded map
	if spc, isStateful := contract.(StatefulPrecompiledContract); isStateful {
		if _, found := pr.statefulPrecompiles[moduleEthAddr]; found {
			return errors.Wrap(ErrPrecompileAlreadyRegistered, moduleName)
		}
		pr.statefulPrecompiles[moduleEthAddr] = spc
	}

	// register the module's events if the precompile contract exposes any events
	if eventsContract, hasEvents := contract.(precompile.HasEvents); hasEvents {
		for _, abiEvent := range eventsContract.ABIEvents() {
			var customModuleAttributes log.ValueDecoders
			if customEvents, isCustom := contract.(precompile.HasCustomEvents); isCustom {
				// if contract is of a custom Cosmos module, load its custom attributes' value
				// decoders
				customModuleAttributes =
					customEvents.CustomValueDecoders()[precompile.EventType(abiEvent.Name)]
			}
			err := pr.logFactory.RegisterEvent(moduleEthAddr, abiEvent, customModuleAttributes)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// `RegisterDynamicContract` stores any module's dynamic stateful precompile in the EVM kvstore.
func (pr *PrecompileRegistry) RegisterDynamicContract(
	ctx sdk.Context,
	addr common.Address,
	dpc DynamicPrecompiledContract,
) error {
	precompileName := dpc.Name()

	// store dyanmic precompiled contract object in memory
	if _, ok := pr.dynamicPrecompiles[precompileName]; ok {
		return errors.Wrap(ErrPrecompileAlreadyRegistered, precompileName)
	}
	pr.dynamicPrecompiles[precompileName] = dpc

	// add address-precompileName pair to kv store
	store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
	store.Set(addr.Bytes(), []byte(precompileName))
	return nil
}

// ==============================================================================
// Getters
// ==============================================================================

// `GetLogFactory` returns the Ethereum log dynamic for this precompile registry.
func (pr *PrecompileRegistry) GetLogFactory() *precompile.LogFactory {
	return pr.logFactory
}

// `GetStatelessContract` returns a stateless precompile in memory at the given `addr`.
func (pr *PrecompileRegistry) GetStatelessContract(addr common.Address) (PrecompiledContract, bool) {
	pc, found := pr.statelessPrecompiles[addr]
	return pc, found
}

// `GetStatefulContract` returns a stateful precompile in memory at the given `addr`.
func (pr *PrecompileRegistry) GetStatefulContract(addr common.Address) (PrecompiledContract, bool) {
	pc, found := pr.statefulPrecompiles[addr]
	return pc, found
}

// `GetDynamicContract` returns a dynamic precompile from `ctx` stores at the given `addr`.
func (pr *PrecompileRegistry) GetDynamicContract(
	ctx sdk.Context,
	addr common.Address,
) (PrecompiledContract, bool) {
	store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
	name := store.Get(addr.Bytes())
	if name == nil {
		return nil, false
	}
	dpc, found := pr.dynamicPrecompiles[string(name)]
	return dpc, found
}
