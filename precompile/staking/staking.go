//nolint:lll
//go:generate go run github.com/berachain/stargazer/cmd/abigen staking ../contracts/solidity/out/staking.sol/StakingEvents.json  ../staking/contract.abigen.go StakingEvents
//go:generate go run github.com/berachain/stargazer/cmd/abigen staking ../contracts/solidity/out/staking.sol/IStakingModule.json ../staking/contract.abigen.go IStakingModule

package staking

import (
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
