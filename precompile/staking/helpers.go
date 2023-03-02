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
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"pkg.berachain.dev/stargazer/eth/common"
	evmutils "pkg.berachain.dev/stargazer/x/evm/utils"
)

// `delegationHelper` is the helper function for `getDelegation`.
func (c *Contract) getDelegationHelper(
	ctx context.Context,
	del sdk.AccAddress,
	val sdk.ValAddress,
) ([]any, error) {
	res, err := c.querier.Delegation(ctx, &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: del.String(),
		ValidatorAddr: val.String(),
	})
	if err != nil {
		return nil, err
	}

	delegation := res.GetDelegationResponse()
	if delegation == nil {
		return nil, err
	}

	return []any{delegation.Balance.Amount.BigInt()}, nil
}

// `getUnbondingDelegationHelper` is the helper function for `getUnbondingDelegation`.
func (c *Contract) getUnbondingDelegationHelper(
	ctx context.Context,
	caller common.Address,
	val sdk.ValAddress,
) ([]any, error) {
	res, err := c.querier.UnbondingDelegation(ctx, &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: evmutils.AddressToAccAddress(caller).String(),
		ValidatorAddr: val.String(),
	})
	if err != nil {
		return nil, errors.New("unbonding delegation not found")
	}

	return []any{res}, nil
}

// `getRedelegationsHelper` is the helper function for `getRedelegations.
func (c *Contract) getRedelegationsHelper(
	ctx context.Context,
	caller common.Address,
	srcValidator sdk.ValAddress,
	dstValidator sdk.ValAddress,
) ([]any, error) {
	rsp, err := c.querier.Redelegations(
		ctx,
		&stakingtypes.QueryRedelegationsRequest{
			DelegatorAddr:    evmutils.AddressToAccAddress(caller).String(),
			SrcValidatorAddr: srcValidator.String(),
			DstValidatorAddr: dstValidator.String(),
		},
	)

	var redelegationEntryResponses []stakingtypes.RedelegationEntryResponse
	for _, r := range rsp.GetRedelegationResponses() {
		redel := r.GetRedelegation()
		if redel.DelegatorAddress == evmutils.AddressToAccAddress(caller).String() &&
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

	return []any{redelegationEntries}, err
}

// `delegateHelper` is the helper function for `delegate`.
func (c *Contract) delegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	validatorAddress sdk.ValAddress,
) error {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return err
	}

	_, err = c.msgServer.Delegate(ctx, stakingtypes.NewMsgDelegate(
		evmutils.AddressToAccAddress(caller),
		validatorAddress,
		sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount)),
	))
	return err
}

// `undelegateHelper` is the helper function for `undelegate`.
func (c *Contract) undelegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	val sdk.ValAddress,
) error {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return err
	}

	_, err = c.msgServer.Undelegate(ctx, stakingtypes.NewMsgUndelegate(
		evmutils.AddressToAccAddress(caller),
		val,
		sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount)),
	))

	return err
}

// `beginRedelegateHelper` is the helper function for `beginRedelegate`.
func (c *Contract) beginRedelegateHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	srcVal, dstVal sdk.ValAddress,
) error {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return err
	}

	_, err = c.msgServer.BeginRedelegate(
		ctx,
		stakingtypes.NewMsgBeginRedelegate(
			evmutils.AddressToAccAddress(caller),
			srcVal,
			dstVal,
			sdk.NewCoin(bondDenom, sdk.NewIntFromBigInt(amount)),
		),
	)

	return err
}

// `cancelRedelegateHelper` is the helper function for `cancelRedelegate`.
func (c *Contract) cancelUnbondingDelegationHelper(
	ctx context.Context,
	caller common.Address,
	amount *big.Int,
	val sdk.ValAddress,
	creationHeight int64,
) error {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return err
	}

	_, err = c.msgServer.CancelUnbondingDelegation(
		ctx,
		stakingtypes.NewMsgCancelUnbondingDelegation(
			evmutils.AddressToAccAddress(caller),
			val,
			creationHeight,
			sdk.NewCoin(bondDenom, sdk.NewIntFromBigInt(amount)),
		),
	)

	return err
}

func (c *Contract) activeValidatorsHelper(ctx context.Context) ([]any, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status: stakingtypes.BondStatusBonded,
	})
	if err != nil {
		return nil, err
	}
	// Iterate over all validators and return their addresses.
	addrs := make([]common.Address, 0, len(res.Validators))
	for i, val := range res.Validators {
		var valAddr sdk.ValAddress
		valAddr, err = sdk.ValAddressFromBech32(val.OperatorAddress)
		if err != nil {
			return nil, err
		}
		addrs[i] = evmutils.ValAddressToEthAddress(valAddr)
	}

	return []any{addrs}, nil
}

// `bondDenom` returns the bond denom from the staking module.
func (c *Contract) bondDenom(ctx context.Context) (string, error) {
	res, err := c.querier.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}

	return res.Params.BondDenom, nil
}
