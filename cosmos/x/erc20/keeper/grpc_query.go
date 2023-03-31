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

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
)

// Compile-time interface assertion.
var _ types.QueryServiceServer = Querier{}

// Querier implements the QueryServer for the erc20 module.
type Querier struct {
	*Keeper
}

// ERC20AddressForCoinDenom queries the ERC20 token address for a given SDK coin denomination.
func (q Querier) ERC20AddressForCoinDenom(
	ctx context.Context, req *types.ERC20AddressForCoinDenomRequest,
) (*types.ERC20AddressForCoinDenomResponse, error) {
	return &types.ERC20AddressForCoinDenomResponse{
		Token: cosmlib.AddressToAccAddress(
			q.DenomKVStore(sdk.UnwrapSDKContext(ctx)).GetAddressForDenom(req.Denom),
		).String(),
	}, nil
}

// CoinDenomForERC20Address queries the SDK coin denomination for a given ERC20 token address.
func (q Querier) CoinDenomForERC20Address(
	ctx context.Context, req *types.CoinDenomForERC20AddressRequest,
) (*types.CoinDenomForERC20AddressResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Token)
	if err != nil {
		return nil, err
	}

	return &types.CoinDenomForERC20AddressResponse{
		Denom: q.DenomKVStore(sdk.UnwrapSDKContext(ctx)).GetDenomForAddress(
			cosmlib.AccAddressToEthAddress(addr),
		),
	}, nil
}
