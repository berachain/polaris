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

package auth

import (
	"context"
	"math/big"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// sendGrantHelper is the helper method to call the grant method on the msgServer, with a send authorization.
func (c *Contract) sendGrantHelper(
	ctx context.Context,
	blocktime time.Time,
	granter, grantee sdk.AccAddress,
	limit sdk.Coins,
	expiration *big.Int,
) ([]any, error) {
	var (
		grant authz.Grant
		err   error
	)

	// Create the send authorization.
	sendAuth := banktypes.NewSendAuthorization(limit, []sdk.AccAddress{grantee})

	// If the expiration is 0, then the grant is valid forever, and can be nil.
	if expiration.Cmp(big.NewInt(0)) == 0 {
		grant, err = authz.NewGrant(blocktime, sendAuth, nil)
	} else {
		expirationTime := time.Unix(expiration.Int64(), 0)
		grant, err = authz.NewGrant(blocktime, sendAuth, &expirationTime)
	}

	// Assert that the grant is valid.
	if err != nil {
		return nil, err
	}

	_, err = c.msgServer.Grant(ctx, &authz.MsgGrant{
		Granter: granter.String(),
		Grantee: grantee.String(),
		Grant:   grant,
	})

	return []any{err == nil}, err
}

// getSendAllownace returns the highest allowance for a given coin denom.
func (c *Contract) getSendAllownaceHelper(
	ctx context.Context,
	blocktime time.Time,
	granter, grantee sdk.AccAddress,
	coinDenom string,
) ([]any, error) {
	// Get the grants from the query server.
	res, err := c.queryServer.Grants(ctx, &authz.QueryGrantsRequest{
		Granter:    granter.String(),
		Grantee:    grantee.String(),
		MsgTypeUrl: banktypes.SendAuthorization{}.MsgTypeURL(),
		Pagination: nil,
	})
	if err != nil {
		//nolint:nilerr // We want to ignore the error here.
		return []any{big.NewInt(0)}, nil
	}

	// If there are no grants, then the allowance is 0.
	if len(res.Grants) == 0 {
		return []any{big.NewInt(0)}, nil
	}

	// Map the grants to send authorizations, should have the same type since we filtered by msg type url.
	sendAuths, err := mapGrantToSendAuth(res.Grants, blocktime)
	if err != nil {
		return nil, err // Hard error here since this is a faliure in the precompiled contract.
	}

	// Get the highest allowance from the send authorizations.
	allowance := getHighestAllowance(sendAuths, coinDenom)

	return []any{allowance}, nil
}

// extractCoinsFromInput converts coins from input (of type any) into sdk.Coins.
func extractCoinsFromInput(coins any) (sdk.Coins, error) {
	// note: we have to use unnamed struct here, otherwise the compiler cannot cast
	// the any type input into IBankModuleCoin.
	amounts, ok := utils.GetAs[[]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}](coins)
	if !ok {
		return nil, precompile.ErrInvalidCoin
	}

	sdkCoins := sdk.NewCoins()
	for _, evmCoin := range amounts {
		sdkCoins = sdkCoins.Add(
			sdk.Coin{
				Amount: sdk.NewIntFromBigInt(evmCoin.Amount),
				Denom:  evmCoin.Denom,
			},
		)
	}
	return sdkCoins, nil
}

// mapGrantToSendAuth maps a list of grants to a list of send authorizations.
func mapGrantToSendAuth(grants []*authz.Grant, blocktime time.Time) ([]*banktypes.SendAuthorization, error) {
	var sendAuths []*banktypes.SendAuthorization
	for _, grant := range grants {
		// Check that the expiration is still valid.
		if grant.Expiration == nil || grant.Expiration.After(blocktime) {
			sendAuth, ok := grant.Authorization.GetCachedValue().(*banktypes.SendAuthorization)
			if !ok {
				return nil, precompile.ErrInvalidAny
			}
			sendAuths = append(sendAuths, sendAuth)
		}
	}
	return sendAuths, nil
}
