// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package staking

import (
	"context"
	"math/big"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	cbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the precompile contract for the staking module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer
}

// NewContract is the constructor of the staking contract.
func NewPrecompileContract(sk *stakingkeeper.Keeper) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.StakingModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(stakingtypes.ModuleName)),
		),
		msgServer: stakingkeeper.NewMsgServerImpl(sk),
		querier:   stakingkeeper.Querier{Keeper: sk},
	}
}

// GetActiveValidators implements the `getActiveValidators(PageRequest)` method.
func (c *Contract) GetActiveValidators(
	ctx context.Context,
	pagination any,
) ([]common.Address, cbindings.CosmosPageResponse, error) {
	return c.activeValidatorsHelper(ctx, pagination)
}

// GetValidators implements the `getValidators(PageRequest)` method.
func (c *Contract) GetValidators(
	ctx context.Context,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	return c.validatorsHelper(ctx, pagination)
}

// GetValidators implements the `getValidator(address)` method.
func (c *Contract) GetValidator(
	ctx context.Context,
	validatorAddr common.Address,
) (generated.IStakingModuleValidator, error) {
	return c.validatorHelper(ctx, sdk.ValAddress(validatorAddr[:]).String())
}

// GetDelegatorValidators implements the `getDelegatorValidators(address)` method.
func (c *Contract) GetDelegatorValidators(
	ctx context.Context,
	delegatorAddr common.Address,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.DelegatorValidators(ctx, &stakingtypes.QueryDelegatorValidatorsRequest{
		DelegatorAddr: cosmlib.Bech32FromEthAddress(delegatorAddr),
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(res.GetValidators())
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	return vals, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetValidatorDelegations implements the `getValidatorDelegations(address,PageRequest)` method.
func (c *Contract) GetValidatorDelegations(
	ctx context.Context,
	validatorAddr common.Address,
	pagination any,
) ([]generated.IStakingModuleDelegation, cbindings.CosmosPageResponse, error) {
	return c.getValidatorDelegationsHelper(ctx, cosmlib.AddressToValAddress(validatorAddr), pagination)
}

// GetDelegation implements `getDelegation(address)` method.
func (c *Contract) GetDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) (*big.Int, error) {
	return c.getDelegationHelper(
		ctx,
		cosmlib.AddressToAccAddress(delegatorAddress),
		cosmlib.AddressToValAddress(validatorAddress),
	)
}

// GetUnbondingDelegation implements the `getUnbondingDelegation(address,address)` method.
func (c *Contract) GetUnbondingDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) ([]generated.IStakingModuleUnbondingDelegationEntry, error) {
	return c.getUnbondingDelegationHelper(
		ctx, cosmlib.AddressToAccAddress(delegatorAddress), cosmlib.AddressToValAddress(validatorAddress),
	)
}

// GetDelegatorUnbondingDelegations implements the `getDelegatorUnbondingDelegations(address)` method.
func (c *Contract) GetDelegatorUnbondingDelegations(
	ctx context.Context,
	delegatorAddress common.Address,
	pagination any,
) ([]generated.IStakingModuleUnbondingDelegation, cbindings.CosmosPageResponse, error) {
	return c.getDelegatorUnbondingDelegationsHelper(
		ctx, cosmlib.AddressToAccAddress(delegatorAddress), pagination,
	)
}

// GetRedelegations implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegations(
	ctx context.Context,
	delegatorAddress common.Address,
	srcValidator common.Address,
	dstValidator common.Address,
	pagination any,
) ([]generated.IStakingModuleRedelegationEntry, cbindings.CosmosPageResponse, error) {
	del := cosmlib.Bech32FromEthAddress(delegatorAddress)
	rsp, err := c.querier.Redelegations(
		ctx,
		&stakingtypes.QueryRedelegationsRequest{
			DelegatorAddr:    cosmlib.Bech32FromEthAddress(delegatorAddress),
			SrcValidatorAddr: cosmlib.AddressToValAddress(srcValidator).String(),
			DstValidatorAddr: cosmlib.AddressToValAddress(dstValidator).String(),
			Pagination:       cosmlib.ExtractPageRequestFromInput(pagination),
		},
	)
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleRedelegationEntry{}, cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	var redelegationEntryResponses []stakingtypes.RedelegationEntryResponse
	for _, r := range rsp.GetRedelegationResponses() {
		redel := r.GetRedelegation()
		if redel.DelegatorAddress == del &&
			redel.ValidatorSrcAddress == srcValidator.String() &&
			redel.ValidatorDstAddress == dstValidator.String() {
			redelegationEntryResponses = r.GetEntries()
			break
		}
	}
	redelegationEntries := make(
		[]stakingtypes.RedelegationEntry, 0, len(redelegationEntryResponses),
	)
	for _, entryRsp := range redelegationEntryResponses {
		redelegationEntries = append(redelegationEntries, entryRsp.GetRedelegationEntry())
	}

	return cosmlib.SdkREToStakingRE(redelegationEntries),
		cosmlib.SdkPageResponseToEvmPageResponse(rsp.Pagination),
		err
}

// Delegate implements the `delegate(address,uint256)` method.
func (c *Contract) Delegate(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
) (bool, error) {
	return c.delegateHelper(
		ctx,
		vm.UnwrapPolarContext(ctx).MsgSender(),
		amount,
		cosmlib.AddressToValAddress(validatorAddress),
	)
}

// Undelegate implements the `undelegate(address,uint256)` method.
func (c *Contract) Undelegate(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
) (bool, error) {
	return c.undelegateHelper(
		ctx,
		vm.UnwrapPolarContext(ctx).MsgSender(),
		amount,
		cosmlib.AddressToValAddress(validatorAddress),
	)
}

// BeginRedelegate implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegate(
	ctx context.Context,
	srcValidator common.Address,
	dstValidator common.Address,
	amount *big.Int,
) (bool, error) {
	return c.beginRedelegateHelper(
		ctx,
		vm.UnwrapPolarContext(ctx).MsgSender(),
		amount,
		cosmlib.AddressToValAddress(srcValidator),
		cosmlib.AddressToValAddress(dstValidator),
	)
}

// CancelRedelegate implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegation(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
	creationHeight int64,
) (bool, error) {
	return c.cancelUnbondingDelegationHelper(
		ctx,
		vm.UnwrapPolarContext(ctx).MsgSender(),
		amount,
		cosmlib.AddressToValAddress(validatorAddress),
		creationHeight,
	)
}
