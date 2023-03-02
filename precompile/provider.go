package precompile

import (
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/precompile/staking"
)

type Provider interface {
	// `GetPrecompile` returns the precompile contract at the given address.
	GetPrecompiles() []vm.RegistrablePrecompile

	IsPrecompileProvider() bool
}

type provider struct {
	precompiles []vm.RegistrablePrecompile
}

func NewProvider(sk *stakingkeeper.Keeper) *provider {
	return &provider{
		precompiles: []vm.RegistrablePrecompile{
			staking.NewPrecompileContract(&sk),
		},
	}
}

func (p *provider) GetPrecompiles() []vm.RegistrablePrecompile {
	return p.precompiles
}

func (p *provider) IsPrecompileProvider() bool {
	return true
}
