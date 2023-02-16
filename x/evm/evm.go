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

package evm

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	store "cosmossdk.io/store/types"
	modulev1 "github.com/berachain/stargazer/api/stargazer/evm/module/v1"
	"github.com/berachain/stargazer/x/evm/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

//nolint:gochecknoinits // GRRRR fix later.
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

// `DepInjectInput` is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	ModuleKey depinject.OwnModuleKey
	Config    *modulev1.Module
	Key       *store.KVStoreKey

	AccountKeeper AccountKeeper
	BankKeeper    BankKeeper
	StakingKeeper StakingKeeper
}

// `DepInjectOutput` is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out
	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

// `ProvideModule` is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) DepInjectOutput {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		// in.Key,
		// in.StakingKeeper,
		in.AccountKeeper,
		in.BankKeeper,
		authority.String(),
	)

	m := NewAppModule(k,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return DepInjectOutput{Keeper: k, Module: m}
}
