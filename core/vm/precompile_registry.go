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
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/event"
	"github.com/berachain/stargazer/lib/common"
)

// KeyPrefixPrecompileAddress is the prefix for the precompile address to name mapping in the
// precompile kvstore.
var KeyPrefixPrecompileAddress = []byte{0x01}

// `PrecompileRegistry` will store and provide stateless and custom or factory stateful precompiled
// contracts to the EVM.
type PrecompileRegistry struct {
	// `StatelessPrecompiles` is a map of Ethereum addresses
	statelessPrecompiles map[common.Address]PrecompiledContract

	// `statefulPrecompiles` is a map of Ethereum addresses to stateful precompiled contracts.
	statefulPrecompiles map[common.Address]StatefulPrecompiledContract

	// `factoryPrecompiles` is a map of names to factory precompiled contracts.
	factoryPrecompiles map[string]FactoryPrecompiledContract

	// `storeKey` is the store key of the Cosmos KVStore used for storing precompiles.
	storeKey storetypes.StoreKey

	// `eventFactory` is the Ethereum log builder for all Cosmos events emitted during precompile
	// execution.
	eventFactory *precompile.EthereumLogFactory
}

// `NewPrecompileRegistry` creates and returns a new `PrecompileRegistry` for given `storeKey`.
func NewPrecompileRegistry(storeKey storetypes.StoreKey) *PrecompileRegistry {
	return &PrecompileRegistry{
		statelessPrecompiles: make(map[common.Address]PrecompiledContract),
		statefulPrecompiles:  make(map[common.Address]StatefulPrecompiledContract),
		factoryPrecompiles:   make(map[string]FactoryPrecompiledContract),
		storeKey:             storeKey,
		eventFactory:         precompile.NewEthereumLogFactory(),
	}
}

// `GetEventFactory` returns the Ethereum log factory for this precompile registry.
func (pr *PrecompileRegistry) GetEventFactory() *precompile.EthereumLogFactory {
	return pr.eventFactory
}

// `InjectStateless` stores the stateless precompile `pc` at the given Ethereum address `addr`.
func (pr *PrecompileRegistry) InjectStatelessContract(
	addr common.Address,
	pc PrecompiledContract,
) {
	pr.statelessPrecompiles[addr] = pc
}

// `RegisterModule` stores a module's stateful precompile contract (in memory) at a hardcoded
// address (module account address) and registers its events if it has any.
func (pr *PrecompileRegistry) RegisterModule(moduleName string, contract any) error {
	moduleEthAddr := common.BytesToAddress(authtypes.NewModuleAddress(moduleName).Bytes())

	// store the module stateful precompile contract in the hardcoded map
	if spc, ok := contract.(StatefulPrecompiledContract); ok {
		pr.statefulPrecompiles[moduleEthAddr] = spc
	}

	// register the module's events if the precompile contract exposes any events
	if eventsContract, hasEvents := contract.(precompile.HasEvents); hasEvents {
		for _, abiEvent := range eventsContract.ABIEvents() {
			var customModuleAttributes event.ValueDecoders
			if customModule, isCustom := contract.(precompile.HasCustomModuleEvents); isCustom {
				// if contract is of a custom Cosmos module, load its custom attributes' value
				// decoders
				customModuleAttributes =
					customModule.CustomValueDecoders()[precompile.EventType(abiEvent.Name)]
			}
			err := pr.eventFactory.RegisterEvent(moduleEthAddr, abiEvent, customModuleAttributes)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// `Inject` stores any module's factory stateful precompile in the EVM kvstore.
func (pr *PrecompileRegistry) InjectFactoryContract(
	ctx sdk.Context,
	addr common.Address,
	fpc FactoryPrecompiledContract,
) error {
	precompileName := fpc.Name()

	// store stateful precompiled contract object in memory
	if _, ok := pr.factoryPrecompiles[precompileName]; ok {
		return fmt.Errorf("%s precompile is already injected", precompileName)
	}
	pr.factoryPrecompiles[precompileName] = fpc

	// add address-precompileName pair to kv store
	store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
	store.Set(addr.Bytes(), []byte(precompileName))
	return nil
}

// `GetPrecompileFn` returns a `PrecompileGetter` function, to be used by the EVM.
func (pr *PrecompileRegistry) GetPrecompileFn(ctx sdk.Context) PrecompileGetter {
	return func(addr common.Address) (PrecompiledContract, bool) {
		// try hardcoded precompile in memory
		spc, found := pr.statefulPrecompiles[addr]
		if found {
			return spc, found
		}

		// try dynamically loading precompile from kvstore
		store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
		name := store.Get(addr.Bytes())
		if name == nil {
			return nil, false
		}

		spc, found = pr.factoryPrecompiles[string(name)]
		return spc, found
	}
}
