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

	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"
)

// RecoveryHandler is used to recover from any WriteProtection and OutOfGas panics that occur
// during precompile execution; the handler modifies the given error to be returned to the caller.
// Any other type of panic is propogated up to the caller.
func RecoveryHandler(err *error) {
	if panicked := recover(); panicked != nil {
		// NOTE: this only propagates an error back to the EVM if the type of the given panic
		// is ErrWriteProtection, Cosmos ErrorOutOfGas, or Cosmos ErrorGasOverflow
		switch {
		case utils.Implements[error](panicked):
			if errors.Is(utils.MustGetAs[error](panicked), vm.ErrWriteProtection) {
				*err = vm.ErrWriteProtection
			}
		case utils.Implements[storetypes.ErrorGasOverflow](panicked):
			fallthrough
		case utils.Implements[storetypes.ErrorOutOfGas](panicked):
			*err = vm.ErrOutOfGas
		default:
			// any other type of panic value is ignored and passed up the call stack
			panic(panicked)
		}
	}
}
