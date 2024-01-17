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

package precompile

import (
	"errors"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core/vm"
)

// RecoveryHandler is used to recover from any WriteProtection and gas consumption panics that
// occur during precompile execution; the handler modifies the given error to be returned to the
// caller. Any other type of panic is propagated up to the caller via panic.
func RecoveryHandler(ctx sdk.Context, vmErr *error) {
	if panicked := recover(); panicked != nil {
		// NOTE: this only propagates an error back to the EVM if the type of the given panic
		// is ErrWriteProtection, Cosmos ErrorOutOfGas, Cosmos ErrorGasOverflow, or Cosmos
		// ErrorNegativeGasConsumed.
		switch {
		case utils.Implements[error](panicked) &&
			errors.Is(utils.MustGetAs[error](panicked), vm.ErrWriteProtection):
			*vmErr = vm.ErrWriteProtection
		case utils.Implements[storetypes.ErrorGasOverflow](panicked):
			fallthrough
		case utils.Implements[storetypes.ErrorOutOfGas](panicked):
			fallthrough
		case utils.Implements[storetypes.ErrorNegativeGasConsumed](panicked):
			*vmErr = vm.ErrOutOfGas
		case utils.Implements[error](panicked):
			// any other type of panic value is returned as a vm error: execution reverted
			// NOTE: precompile txs which panic will be included in the block as failed txs
			ctx.Logger().Error("panic recovered in precompile execution", "err", panicked)
			*vmErr = errors.Join(vm.ErrExecutionReverted, utils.MustGetAs[error](panicked))
		default:
			ctx.Logger().Error("panic recovered in precompile execution", "panic", panicked)
			*vmErr = vm.ErrExecutionReverted
		}
	}
}
