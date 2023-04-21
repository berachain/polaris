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

package builder

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	builderkeeper "github.com/skip-mev/pob/x/builder/keeper"
	buildertypes "github.com/skip-mev/pob/x/builder/types"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the builder module.
type Contract struct {
	precompile.BaseContract

	msgServer   buildertypes.MsgServer
	queryServer buildertypes.QueryServer
	evmDenom    string
}

// NewPrecompileContract returns a new instance of the builder module precompile contract.
func NewPrecompileContract(bk *builderkeeper.Keeper, evmDenom string) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			bindings.BuilderModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(buildertypes.ModuleName)),
		),
		msgServer:   builderkeeper.NewMsgServerImpl(*bk),
		queryServer: builderkeeper.NewQueryServer(*bk),
		evmDenom:    evmDenom,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:      "auctionBid(uint256,bytes[],uint64)",
			Execute:     c.AuctionBid,
			RequiredGas: params.IdentityBaseGas,
		},
	}
}

// AuctionBid implements the `auctionBid(uint256,bytes[],uint64)` method.
func (c *Contract) AuctionBid(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bid, ok := utils.GetAs[*big.Int](args[0])
	if !ok {
		return nil, precompile.ErrInvalidCoin
	}
	bundleTxs, ok := utils.GetAs[[][]byte](args[1])
	if !ok {
		return nil, precompile.ErrInvalidAny
	}

	msgAuctionBid := &buildertypes.MsgAuctionBid{
		Bidder:       cosmlib.AddressToAccAddress(caller).String(),
		Bid:          sdk.NewCoin(c.evmDenom, sdk.NewIntFromBigInt(bid)),
		Transactions: bundleTxs,
	}

	// Do we need to add timeout logic here as well?

	_, err := c.msgServer.AuctionBid(ctx, msgAuctionBid)

	return []any{err == nil}, nil
}
