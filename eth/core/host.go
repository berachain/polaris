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

package core

import libtypes "github.com/berachain/stargazer/lib/types"

// =============================================================================
// Mandatory Plugins
// =============================================================================

// The following plugins MUST be implemented by the chain running Stargazer EVM and exposed via the
// `StargazerHostChain` interface. All plugins should be resettable with a given context.
type (
	// `GasPlugin` is an interface that allows the Stargazer EVM to consume gas on the host chain.
	GasPlugin interface {
		// `GasPlugin` implements `libtypes.Resettable`. Calling Reset() MUST reset the GasPlugin to a
		// default state.
		libtypes.Resettable

		// `ConsumeGas` MUST consume the supplied amount of gas. It MUST not panic due to a GasOverflow
		// and must return core.ErrOutOfGas if the amount of gas remaining is less than the amount
		// requested.
		ConsumeGas(uint64) error

		// `RefundGas` MUST refund the supplied amount of gas. It MUST not panic.
		RefundGas(uint64)

		// `GasRemaining` MUST return the amount of gas remaining. It MUST not panic.
		GasRemaining() uint64

		// `GasUsed` MUST return the amount of gas used during the current transaction. It MUST not panic.
		GasUsed() uint64

		// `CumulativeGasUsed` MUST return the amount of gas used during the current block. The value returned
		// MUST include any gas consumed during this transaction. It MUST not panic.
		CumulativeGasUsed() uint64

		// `MaxFeePerGas` MUST set the maximum amount of gas that can be consumed by the meter. It MUST not panic, but
		// instead, return an error, if the new gas limit is less than the currently consumed amount of gas.
		SetGasLimit(uint64) error
	}
)
