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

type Querier struct {
	*Keeper
}

// ERC20AddressForDenom queries the ERC20 address for a given denom.
func (q Querier) ERC20AddressForDenom(
	ctx context.Context, req *types.ERC20AddressForDenomRequest,
) (*types.ERC20AddressForDenomResponse, error) {
	addr, err := q.DenomKVStore(sdk.UnwrapSDKContext(ctx)).GetAddressForDenom(req.Denom)
	if err != nil {
		return nil, err
	}

	return &types.ERC20AddressForDenomResponse{
		Address: cosmlib.AddressToAccAddress(addr).String(),
	}, nil
}

// DenomForERC20Address queries the denom for a given ERC20 address.
func (q Querier) DenomForERC20Address(
	ctx context.Context, req *types.DenomForERC20AddressRequest,
) (*types.DenomForERC20AddressResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	denom, err := q.DenomKVStore(sdk.UnwrapSDKContext(ctx)).GetDenomForAddress(
		cosmlib.AccAddressToEthAddress(addr),
	)
	if err != nil {
		return nil, err
	}

	return &types.DenomForERC20AddressResponse{
		Denom: denom,
	}, nil
}
