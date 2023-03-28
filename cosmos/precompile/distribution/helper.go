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

package distribution

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
)

// `setWithdrawAddressHelper` is a helper function for the `SetWithdrawAddress` method.
func (c *Contract) setWithdrawAddressHelper(ctx context.Context, delegator, withdrawer sdk.AccAddress) ([]any, error) {
	_, err := c.msgServer.SetWithdrawAddress(ctx, &distributiontypes.MsgSetWithdrawAddress{
		DelegatorAddress: delegator.String(),
		WithdrawAddress:  withdrawer.String(),
	})
	return []any{err == nil}, err
}

func (c *Contract) getWithdrawAddrEnabled(ctx context.Context) ([]any, error) {
	res, err := c.querier.Params(ctx, &distributiontypes.QueryParamsRequest{})
	return []any{res.Params.WithdrawAddrEnabled}, err
}

// `withdrawDelegatorRewards` is a helper function for the `WithdrawDelegatorRewards` method.
func (c *Contract) withdrawDelegatorRewardsHelper(
	ctx context.Context,
	delegator sdk.AccAddress,
	validator sdk.ValAddress,
) ([]any, error) {
	res, err := c.msgServer.WithdrawDelegatorReward(ctx, &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: delegator.String(),
		ValidatorAddress: validator.String(),
	})
	if err != nil {
		return nil, err
	}

	amount := make([]generated.IBankModuleCoin, 0)
	for _, coin := range res.Amount {
		amount = append(amount, generated.IBankModuleCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Uint64(),
		})
	}

	return []any{amount}, nil
}
