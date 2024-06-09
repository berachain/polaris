// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
