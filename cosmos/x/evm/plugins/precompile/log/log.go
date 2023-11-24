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

package log

import (
	"github.com/berachain/polaris/eth/accounts/abi"
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

// Compile-time assertion.
var _ libtypes.Registrable[string] = (*precompileLog)(nil)

// precompileLog contains the required data for a precompile contract to produce an Ethereum
// compatible event log.
type precompileLog struct {
	// eventType is the corresponding Cosmos event type for this precompile log.
	eventType string
	// address is the Ethereum address used as the `Address` field for the Ethereum log.
	precompileAddr common.Address
	// id is the Ethereum event ID, to be used as an Ethereum event's first topic.
	id common.Hash
	// indexedInputs holds an Ethereum event's indexed arguments, emitted as event topics.
	indexedInputs abi.Arguments
	// nonIndexedInputs holds an Ethereum event's non-indexed arguments, emitted as event data.
	nonIndexedInputs abi.Arguments
}

// newPrecompileLog returns a new `precompileLog` with the given `precompileAddress` and
// abiEvent. It separates the indexed and non-indexed arguments of the event.
func newPrecompileLog(precompileAddr common.Address, abiEvent abi.Event) *precompileLog {
	return &precompileLog{
		eventType:        abi.ToUnderScore(abiEvent.Name),
		precompileAddr:   precompileAddr,
		id:               abiEvent.ID,
		indexedInputs:    abi.GetIndexed(abiEvent.Inputs),
		nonIndexedInputs: abiEvent.Inputs.NonIndexed(),
	}
}

// RegistryKey returns the Cosmos event type for the precompile log.
//
// RegistryKey implements `libtypes.Registrable`.
func (l *precompileLog) RegistryKey() string {
	return l.eventType
}
