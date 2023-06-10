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

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// Compile-time check to ensure `Keeper` implements the `MsgServiceServer` interface.
var _ types.MsgServiceServer = &Keeper{}

// EthTransaction implements the MsgServiceServer interface. It processes an incoming request and
// applies it to the Polaris Chain.
func (k *Keeper) EthTransaction(
	ctx context.Context, msg *types.WrappedEthereumTransaction,
) (*types.WrappedEthereumTransactionResult, error) {
	// Process the transaction and return the result.
	result, err := k.ProcessTransaction(ctx, msg.AsTransaction())
	if err != nil {
		return nil, errorsmod.Wrapf(err, "failed to process transaction")
	}

	// Build the response.
	vmErr := ""
	if result.Err != nil {
		vmErr = result.Err.Error()
	}

	return &types.WrappedEthereumTransactionResult{
		GasUsed:    result.UsedGas,
		VmError:    vmErr,
		ReturnData: result.ReturnData,
	}, nil
}
