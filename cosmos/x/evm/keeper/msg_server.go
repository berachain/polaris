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

	errorsmod "cosmossdk.io/errors"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/configuration"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Compile-time check to ensure `Keeper` implements the `MsgServiceServer` interface.
var _ types.MsgServiceServer = &Keeper{}

// EthTransaction implements the  MsgServiceServer interface. It processes an incoming request
// and applies it to the Polaris Chain.
func (k *Keeper) EthTransaction(
	ctx context.Context, msg *types.EthTransactionRequest,
) (*types.EthTransactionResponse, error) {
	// Process the transaction and return the result.
	receipt, err := k.ProcessTransaction(ctx, msg.AsTransaction())
	if err != nil {
		return nil, errorsmod.Wrapf(err, "failed to process transaction")
	}

	return &types.EthTransactionResponse{
		GasUsed: receipt.GasUsed,
	}, nil
}

// UpdateParams  processes an incoming request and applies it to the Configuration plugin to
// update things about both the Polaris Chain as well as the EVM Module.
func (k *Keeper) UpdateParams(
	ctx context.Context, req *types.UpdateParamsRequest,
) (*types.UpdateParamsResponse, error) {
	// Ensure the authority is valid.
	if k.authority != req.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority, expected %s, got %s", k.authority, req.Authority,
		)
	}

	// Update the params.
	cp := utils.MustGetAs[configuration.Plugin](k.host.GetConfigurationPlugin())
	cp.Prepare(ctx)
	cp.SetParams(&req.Params)
	return &types.UpdateParamsResponse{}, nil
}
