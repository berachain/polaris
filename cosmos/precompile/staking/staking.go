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

	sdkmath "cosmossdk.io/math"

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
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status:     stakingtypes.BondStatusBonded,
		Pagination: cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	// Iterate over all validators and return their addresses.
	addrs := make([]common.Address, 0, len(res.Validators))
	for _, val := range res.Validators {
		var valAddr sdk.ValAddress
		valAddr, err = sdk.ValAddressFromBech32(val.OperatorAddress)
		if err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}
		addrs = append(addrs, cosmlib.ValAddressToEthAddress(valAddr))
	}

	pageResponse := cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination)
	return addrs, pageResponse, nil
}

// GetValidators implements the `getValidators(PageRequest)` method.
func (c *Contract) GetValidators(
	ctx context.Context,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status:     stakingtypes.BondStatusBonded,
		Pagination: cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(res.GetValidators())
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	pageResponse := cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination)
	return vals, pageResponse, nil
}

// GetValidators implements the `getValidator(address)` method.
func (c *Contract) GetValidator(
	ctx context.Context,
	validatorAddr common.Address,
) (generated.IStakingModuleValidator, error) {
	res, err := c.querier.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: cosmlib.AddressToValAddress(validatorAddr).String(),
	})
	if err != nil {
		return generated.IStakingModuleValidator{}, err
	}

	val, err := cosmlib.SdkValidatorsToStakingValidators(
		[]stakingtypes.Validator{res.GetValidator()},
	)
	if err != nil {
		return generated.IStakingModuleValidator{}, err
	}

	// guaranteed not to panic because val is guaranteed to have length 1.
	return val[0], nil
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
	res, err := c.querier.ValidatorDelegations(ctx, &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: cosmlib.AddressToValAddress(validatorAddr).String(),
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleDelegation{}, cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	delegations := make([]generated.IStakingModuleDelegation, 0)
	for _, d := range res.GetDelegationResponses() {
		delegations = append(delegations, generated.IStakingModuleDelegation{
			Delegator: cosmlib.EthAddressFromBech32(d.Delegation.DelegatorAddress),
			Shares:    d.Delegation.Shares.BigInt(),
		})
	}

	return delegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetDelegation implements `getDelegation(address)` method.
func (c *Contract) GetDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) (*big.Int, error) {
	res, err := c.querier.Delegation(ctx, &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: cosmlib.AddressToAccAddress(delegatorAddress).String(),
		ValidatorAddr: cosmlib.AddressToValAddress(validatorAddress).String(),
	})
	if status.Code(err) == codes.NotFound {
		// handle the case where the delegation does not exist
		return big.NewInt(0), nil
	} else if err != nil {
		return nil, err
	}

	delegation := res.GetDelegationResponse()
	if delegation == nil {
		return big.NewInt(0), nil
	}

	return delegation.Balance.Amount.BigInt(), nil
}

// GetUnbondingDelegation implements the `getUnbondingDelegation(address,address)` method.
func (c *Contract) GetUnbondingDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) ([]generated.IStakingModuleUnbondingDelegationEntry, error) {
	res, err := c.querier.UnbondingDelegation(ctx, &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: cosmlib.AddressToAccAddress(delegatorAddress).String(),
		ValidatorAddr: cosmlib.AddressToValAddress(validatorAddress).String(),
	})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleUnbondingDelegationEntry{}, nil
	} else if err != nil {
		return nil, err
	}

	return cosmlib.SdkUDEToStakingUDE(res.GetUnbond().Entries), nil
}

// GetDelegatorUnbondingDelegations implements the `getDelegatorUnbondingDelegations(address)` method.
func (c *Contract) GetDelegatorUnbondingDelegations(
	ctx context.Context,
	delegatorAddress common.Address,
	pagination any,
) ([]generated.IStakingModuleUnbondingDelegation, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.DelegatorUnbondingDelegations(ctx, &stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: cosmlib.Bech32FromEthAddress(delegatorAddress),
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleUnbondingDelegation{},
			cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	unbondingDelegations := make([]generated.IStakingModuleUnbondingDelegation, 0)
	for _, u := range res.GetUnbondingResponses() {
		unbondingDelegations = append(unbondingDelegations,
			generated.IStakingModuleUnbondingDelegation{
				DelegatorAddress: cosmlib.EthAddressFromBech32(u.DelegatorAddress),
				ValidatorAddress: cosmlib.EthAddressFromBech32(u.ValidatorAddress),
				Entries:          cosmlib.SdkUDEToStakingUDE(u.Entries),
			},
		)
	}

	return unbondingDelegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
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
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Delegate(ctx, stakingtypes.NewMsgDelegate(
		cosmlib.Bech32FromEthAddress(vm.UnwrapPolarContext(ctx).MsgSender()), /* todo move to codec */
		cosmlib.AddressToValAddress(validatorAddress).String(),               /* todo move to codec */
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// Undelegate implements the `undelegate(address,uint256)` method.
func (c *Contract) Undelegate(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
) (bool, error) {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Undelegate(ctx, stakingtypes.NewMsgUndelegate(
		cosmlib.Bech32FromEthAddress(vm.UnwrapPolarContext(ctx).MsgSender()), /* todo move to codec */
		cosmlib.AddressToValAddress(validatorAddress).String(),               /* todo move to codec */
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// BeginRedelegate implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegate(
	ctx context.Context,
	srcValidator common.Address,
	dstValidator common.Address,
	amount *big.Int,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.BeginRedelegate(
		ctx,
		stakingtypes.NewMsgBeginRedelegate(
			cosmlib.Bech32FromEthAddress(vm.UnwrapPolarContext(ctx).MsgSender()), /* todo move to codec */
			cosmlib.AddressToValAddress(srcValidator).String(),                   /* todo move to codec */
			cosmlib.AddressToValAddress(dstValidator).String(),                   /* todo move to codec */
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err == nil, err
}

// CancelRedelegate implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegation(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
	creationHeight int64,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.CancelUnbondingDelegation(
		ctx,
		stakingtypes.NewMsgCancelUnbondingDelegation(
			cosmlib.Bech32FromEthAddress(vm.UnwrapPolarContext(ctx).MsgSender()), /* todo move to codec */
			cosmlib.AddressToValAddress(validatorAddress).String(),               /* todo move to codec */
			creationHeight,
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err != nil, err
}

// bondDenom returns the bond denom from the staking module.
func (c *Contract) bondDenom(ctx context.Context) (string, error) {
	res, err := c.querier.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}

	return res.Params.BondDenom, nil
}
