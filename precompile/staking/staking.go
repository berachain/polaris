//nolint:lll
//go:generate abigen --pkg generated staking ../contracts/solidity/out/staking.sol/IStakingModule.json --out ./generated/staking.abigen.go --type StakingModule

package staking

import (
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/precompile/contracts/solidity/generated"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	cosmosEventTypes = []string{
		stakingtypes.EventTypeDelegate,
		stakingtypes.EventTypeRedelegate,
		stakingtypes.EventTypeCreateValidator,
		stakingtypes.EventTypeUnbond,
		stakingtypes.EventTypeCancelUnbondingDelegation,
	}
)

var (
	_ precompile.StatefulPrecompileImpl = (*Contract)(nil)
)

type Contract struct {
	vm.PrecompileContainer

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer

	contractAbi abi.ABI
}

// `NewContract` is the constructor of the staking contract.
func NewContract(sk *stakingkeeper.Keeper) *Contract {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		msgServer:   stakingkeeper.NewMsgServerImpl(sk),
		querier:     stakingkeeper.Querier{Keeper: sk},
		contractAbi: contractAbi,
	}
}

// `ABIMethods` implements StatefulPrecompileImpl.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `PrecompileMethods` implements StatefulPrecompileImpl.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return precompile.Methods{}
}

// `ABIEvents` implements StatefulPrecompileImpl.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulPrecompileImpl.
func (c *Contract) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}
