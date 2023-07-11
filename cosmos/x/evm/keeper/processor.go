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

	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// ProcessTransaction is called during the DeliverTx processing of the ABCI lifecycle.
func (k *Keeper) ProcessTransaction(ctx context.Context, tx *coretypes.Transaction) (*core.ExecutionResult, error) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	// We zero-out the gas meter prior to evm execution in order to ensure that the receipt output
	// from the EVM is correct. In the future, we will revisit this to allow gas metering for more
	// complex operations prior to entering the EVM.
	sCtx.GasMeter().RefundGas(sCtx.GasMeter().GasConsumed(),
		"reset gas meter prior to ethereum state transition")

	// Process the transaction and return the EVM's execution result.
	execResult, err := k.polaris.ProcessTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	// We don't want the cosmos transaction to be marked as failed if the EVM reverts. But
	// its not the worst idea to log the error.
	if execResult.Err != nil {
		k.Logger(sdk.UnwrapSDKContext(ctx)).Error(
			"evm execution",
			"tx_hash", tx.Hash(),
			"error", execResult.Err,
			"gas_consumed", sCtx.GasMeter().GasConsumed(),
		)
	} else {
		k.Logger(sdk.UnwrapSDKContext(ctx)).Debug(
			"evm execution",
			"tx_hash", tx.Hash(),
			"gas_consumed", sCtx.GasMeter().GasConsumed(),
		)
	}

	// Return the execution result.
	return execResult, err
}
