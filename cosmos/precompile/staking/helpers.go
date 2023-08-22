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
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	cbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	"pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
)

// delegationHelper is the helper function for `getDelegation`.
func (c *Contract) getDelegationHelper(
	ctx context.Context,
	del sdk.AccAddress,
	val sdk.ValAddress,
) (*big.Int, error) {
	res, err := c.querier.Delegation(ctx, &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: del.String(),
		ValidatorAddr: val.String(),
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

// getDelegationsHelper is the helper function for `getDelegations`.
func (c *Contract) getValidatorDelegationsHelper(
	ctx context.Context,
	val sdk.ValAddress,
	pagination any,
) ([]staking.IStakingModuleDelegation, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.ValidatorDelegations(ctx, &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: val.String(),
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if status.Code(err) == codes.NotFound {
		return []staking.IStakingModuleDelegation{}, cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	delegations := make([]staking.IStakingModuleDelegation, 0)
	for _, d := range res.GetDelegationResponses() {
		delegations = append(delegations, staking.IStakingModuleDelegation{
			Delegator: cosmlib.EthAddressFromBech32(d.Delegation.DelegatorAddress),
			Shares:    d.Delegation.Shares.BigInt(),
		})
	}

	return delegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// getUnbondingDelegationHelper is the helper function for `getUnbondingDelegation`.
func (c *Contract) getUnbondingDelegationHelper(
	ctx context.Context,
	del sdk.AccAddress,
	val sdk.ValAddress,
) ([]staking.IStakingModuleUnbondingDelegationEntry, error) {
	res, err := c.querier.UnbondingDelegation(ctx, &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: del.String(),
		ValidatorAddr: val.String(),
	})
	if status.Code(err) == codes.NotFound {
		return []staking.IStakingModuleUnbondingDelegationEntry{}, nil
	} else if err != nil {
		return nil, err
	}

	return cosmlib.SdkUDEToStakingUDE(res.GetUnbond().Entries), nil
}

// getDelegatorUnbondingDelegationsHelper is the helper function for `getDelegatorUnbondingDelegations`.
func (c *Contract) getDelegatorUnbondingDelegationsHelper(
	ctx context.Context,
	del sdk.AccAddress,
	pagination any,
) ([]staking.IStakingModuleUnbondingDelegation, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.DelegatorUnbondingDelegations(ctx, &stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: del.String(),
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if status.Code(err) == codes.NotFound {
		return []staking.IStakingModuleUnbondingDelegation{},
			cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	unbondingDelegations := make([]staking.IStakingModuleUnbondingDelegation, 0)
	for _, u := range res.GetUnbondingResponses() {
		unbondingDelegations = append(unbondingDelegations,
			staking.IStakingModuleUnbondingDelegation{
				DelegatorAddress: cosmlib.EthAddressFromBech32(u.DelegatorAddress),
				ValidatorAddress: cosmlib.EthAddressFromBech32(u.ValidatorAddress),
				Entries:          cosmlib.SdkUDEToStakingUDE(u.Entries),
			},
		)
	}

	return unbondingDelegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// delegateHelper is the helper function for `delegate`.
func (c *Contract) delegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	validatorAddress sdk.ValAddress,
) (bool, error) {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Delegate(ctx, stakingtypes.NewMsgDelegate(
		cosmlib.AddressToAccAddress(caller).String(), /* todo move to codec */
		validatorAddress.String(),                    /* todo move to codec */
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// undelegateHelper is the helper function for `undelegate`.
func (c *Contract) undelegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	val sdk.ValAddress,
) (bool, error) {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Undelegate(ctx, stakingtypes.NewMsgUndelegate(
		cosmlib.AddressToAccAddress(caller).String(), /* todo move to codec */
		val.String(), /* todo move to codec */
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// beginRedelegateHelper is the helper function for `beginRedelegate`.
func (c *Contract) beginRedelegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	srcVal, dstVal sdk.ValAddress,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.BeginRedelegate(
		ctx,
		stakingtypes.NewMsgBeginRedelegate(
			cosmlib.AddressToAccAddress(caller).String(), /* todo move to codec */
			srcVal.String(), /* todo move to codec */
			dstVal.String(), /* todo move to codec */
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err == nil, err
}

// cancelRedelegateHelper is the helper function for `cancelRedelegate`.
func (c *Contract) cancelUnbondingDelegationHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	val sdk.ValAddress,
	creationHeight int64,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.CancelUnbondingDelegation(
		ctx,
		stakingtypes.NewMsgCancelUnbondingDelegation(
			cosmlib.AddressToAccAddress(caller).String(), /* todo move to codec */
			val.String(), /* todo move to codec */
			creationHeight,
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err != nil, err
}

func (c *Contract) activeValidatorsHelper(
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

func (c *Contract) validatorsHelper(
	ctx context.Context,
	pagination any,
) ([]staking.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
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

// valAddr must be the bech32 address of the validator.
func (c *Contract) validatorHelper(
	ctx context.Context,
	valAddr string,
) (staking.IStakingModuleValidator, error) {
	res, err := c.querier.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: valAddr,
	})
	if err != nil {
		return staking.IStakingModuleValidator{}, err
	}

	val, err := cosmlib.SdkValidatorsToStakingValidators(
		[]stakingtypes.Validator{res.GetValidator()},
	)
	if err != nil {
		return staking.IStakingModuleValidator{}, err
	}

	// guaranteed not to panic because val is guaranteed to have length 1.
	return val[0], nil
}

// bondDenom returns the bond denom from the staking module.
func (c *Contract) bondDenom(ctx context.Context) (string, error) {
	res, err := c.querier.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}

	return res.Params.BondDenom, nil
}
