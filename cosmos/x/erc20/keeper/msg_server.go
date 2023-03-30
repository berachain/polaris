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
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/store"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
)

// Compile-time interface assertion.
var _ types.MsgServiceServer = MsgServer{}

type MsgServer struct {
	*Keeper
}

func (m MsgServer) ConvertERC20ToCosmos(
	ctx context.Context, msg *types.ConvertERC20ToCosmosRequest,
) (*types.ConvertERC20ToCosmosResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get the denom corresponding to a given ERC20 token.
	addr, err := sdk.AccAddressFromBech32(msg.Token)
	if err != nil {
		return nil, err
	}
	tokenAddr := lib.AccAddressToEthAddress(addr)
	denom, err := m.DenomKVStore(sdkCtx).GetDenomForAddress(tokenAddr)

	// If the denom is not found, we need to register it, this means that the token
	// began it's life as an ERC20.
	if errors.Is(err, store.ErrDenomNotFound) {
		m.RegisterDenomTokenPair(sdkCtx, tokenAddr)
	}

	// Mint the Cosmos SDK coins and send to the recipient.
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}
	err = lib.MintCoinsToAddress(
		sdkCtx,
		m.bankKeeper,
		lib.AccAddressToEthAddress(recipientAddr),
		denom,
		msg.Amount.BigInt(),
	)

	return &types.ConvertERC20ToCosmosResponse{Success: err == nil}, err
}

func (m MsgServer) ConvertCosmosToERC20(
	ctx context.Context, msg *types.ConvertCosmosToERC20Request,
) (*types.ConvertCosmosToERC20Response, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get the ERC20 token address coresponding to a given denom.
	tokenAddr, err := m.DenomKVStore(sdkCtx).GetAddressForDenom(msg.Denom)
	if err != nil {
		return nil, err
	}

}
