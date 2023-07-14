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

// getUnbondingDelegationHelper is the helper function for `getUnbondingDelegation`.
func (c *Contract) getUnbondingDelegationHelper(
	ctx context.Context,
	del sdk.AccAddress,
	val sdk.ValAddress,
) ([]any, error) {
	res, err := c.querier.UnbondingDelegation(ctx, &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: del.String(),
		ValidatorAddr: val.String(),
	})
	if status.Code(err) == codes.NotFound {
		return []any{[]stakingtypes.UnbondingDelegationEntry{}}, nil
	} else if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkUDEToStakingUDE(res.GetUnbond().Entries)}, nil
}

// getRedelegationsHelper is the helper function for `getRedelegations.
func (c *Contract) getRedelegationsHelper(
	ctx context.Context,
	del sdk.AccAddress,
	srcValidator sdk.ValAddress,
	dstValidator sdk.ValAddress,
) ([]any, error) {
	rsp, err := c.querier.Redelegations(
		ctx,
		&stakingtypes.QueryRedelegationsRequest{
			DelegatorAddr:    del.String(),
			SrcValidatorAddr: srcValidator.String(),
			DstValidatorAddr: dstValidator.String(),
		},
	)
	if status.Code(err) == codes.NotFound {
		return []any{[]stakingtypes.RedelegationEntry{}}, nil
	} else if err != nil {
		return nil, err
	}

	var redelegationEntryResponses []stakingtypes.RedelegationEntryResponse
	for _, r := range rsp.GetRedelegationResponses() {
		redel := r.GetRedelegation()
		if redel.DelegatorAddress == del.String() &&
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

	return []any{cosmlib.SdkREToStakingRE(redelegationEntries)}, err
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
		cosmlib.AddressToAccAddress(caller),
		validatorAddress,
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
		cosmlib.AddressToAccAddress(caller),
		val,
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
			cosmlib.AddressToAccAddress(caller),
			srcVal,
			dstVal,
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
			cosmlib.AddressToAccAddress(caller),
			val,
			creationHeight,
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err != nil, err
}

func (c *Contract) activeValidatorsHelper(ctx context.Context) ([]common.Address, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status: stakingtypes.BondStatusBonded,
	})
	if err != nil {
		return nil, err
	}

	// Iterate over all validators and return their addresses.
	addrs := make([]common.Address, 0, len(res.Validators))
	for _, val := range res.Validators {
		var valAddr sdk.ValAddress
		valAddr, err = sdk.ValAddressFromBech32(val.OperatorAddress)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, cosmlib.ValAddressToEthAddress(valAddr))
	}
	return addrs, nil
}

func (c *Contract) validatorsHelper(ctx context.Context) ([]any, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status: stakingtypes.BondStatusBonded,
	})
	if err != nil {
		return nil, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(res.GetValidators())
	if err != nil {
		return nil, err
	}

	return []any{vals}, nil
}

// valAddr must be the bech32 address of the validator.
func (c *Contract) validatorHelper(ctx context.Context, valAddr string) ([]any, error) {
	res, err := c.querier.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: valAddr,
	})
	if err != nil {
		return nil, err
	}

	val, err := cosmlib.SdkValidatorsToStakingValidators([]stakingtypes.Validator{res.GetValidator()})
	if err != nil {
		return nil, err
	}

	// guaranteed not to panic because val is guaranteed to have length 1.
	return []any{val[0]}, nil
}

// accAddr must be the bech32 address of the delegator.
func (c *Contract) delegatorValidatorsHelper(ctx context.Context, accAddr string) ([]any, error) {
	res, err := c.querier.DelegatorValidators(ctx, &stakingtypes.QueryDelegatorValidatorsRequest{
		DelegatorAddr: accAddr,
	})
	if err != nil {
		return nil, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(res.GetValidators())
	if err != nil {
		return nil, err
	}

	return []any{vals}, nil
}

// bondDenom returns the bond denom from the staking module.
func (c *Contract) bondDenom(ctx context.Context) (string, error) {
	res, err := c.querier.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}

	return res.Params.BondDenom, nil
}
