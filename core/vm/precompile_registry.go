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

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/vm/precompile/events"
)

// KeyPrefixPrecompileAddress is the prefix for the precompile address to name mapping in the
// precompile kvstore.
var KeyPrefixPrecompileAddress = []byte{0x01}

// Registry will store and provide custom, stateful precompiled smart contracts.
type Registry struct {
	hardcodedPrecompiles      map[common.Address]StatefulContract
	namesToFactoryPrecompiles map[string]StatefulContract
	storeKey                  storetypes.StoreKey

	// custom Cosmos <> EVM events registry
	events *events.Registry
}

func NewRegistry(storeKey storetypes.StoreKey) *Registry {
	return &Registry{
		hardcodedPrecompiles:      make(map[common.Address]StatefulContract),
		namesToFactoryPrecompiles: make(map[string]StatefulContract),
		storeKey:                  storeKey,
		events:                    events.NewRegistry(),
	}
}

func (pr *Registry) GetEventsRegistry() *events.Registry {
	return pr.events
}

// RegisterModule stores a module's evm stateful precompile contract (in memory) at hardcoded
// addresses and registers its events if it has any.
func (pr *Registry) RegisterModule(moduleName string, contract any) {
	moduleAddr := common.BytesToAddress(authtypes.NewModuleAddress(moduleName).Bytes())

	if spc, ok := contract.(StatefulContract); ok {
		pr.hardcodedPrecompiles[moduleAddr] = spc
	}
	if eventsContract, ok := contract.(events.HasEvents); ok {
		pr.events.RegisterModule(&moduleAddr, eventsContract)
	}
}

// Inject stores any module's factory stateful precompile in the precompile kvstore.
func (pr *Registry) Inject(
	ctx sdk.Context,
	addr common.Address,
	fpc FactoryContract,
) error {
	precompileName := fpc.Name()

	// store stateful precompiled contract object in memory
	if _, ok := pr.namesToFactoryPrecompiles[precompileName]; ok {
		return fmt.Errorf("%s precompile is already injected", precompileName)
	}
	pr.namesToFactoryPrecompiles[precompileName] = fpc

	// add address-precompileName pair to kv store
	store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
	store.Set(addr.Bytes(), []byte(precompileName))
	return nil
}

func (pr *Registry) GetPrecompileFn(ctx sdk.Context) Getter {
	return func(addr common.Address) (PrecompiledContract, bool) {
		// try hardcoded precompile in memory
		spc, found := pr.hardcodedPrecompiles[addr]
		if found {
			return spc, found
		}

		// try dynamically loading precompile from kvstore
		store := prefix.NewStore(ctx.KVStore(pr.storeKey), KeyPrefixPrecompileAddress)
		name := store.Get(addr.Bytes())
		if name == nil {
			return nil, false
		}

		spc, found = pr.namesToFactoryPrecompiles[string(name)]
		return spc, found
	}
}
