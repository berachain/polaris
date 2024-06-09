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
