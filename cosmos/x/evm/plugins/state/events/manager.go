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

package events

import (
	"errors"

	errlib "github.com/berachain/polaris/lib/errors"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

// ErrEthEventNotRegistered is returned when an incoming event is not mapped to any registered
// Ethereum event.
var ErrEthEventNotRegistered = errors.New("no Ethereum event was registered for this Cosmos event")

// manager is a controllable event manager that supports snapshots and reverts for emitted Cosmos
// events. During precompile execution, it is also used to Cosmos events to the Eth logs journal.
type manager struct {
	// EventManager is the underlying Cosmos SDK event manager floating around on the context.
	*sdk.EventManager
	// ldb is the reference to the StateDB for adding Eth logs during precompile execution.
	ldb LogsDB
	// plf is used to build Eth logs from Cosmos events.
	plf PrecompileLogFactory
	// readOnly is true if the EVM is in read-only mode
	readOnly bool
}

// NewManager creates and returns a controllable event manager from the given Cosmos SDK context.
//
//nolint:revive // only used as a `state.ControllableEventManager`.
func NewManagerFrom(em sdk.EventManagerI, plf PrecompileLogFactory) *manager {
	return &manager{
		EventManager: utils.MustGetAs[*sdk.EventManager](em),
		plf:          plf,
	}
}

// IsReadOnly returns the current read-only mode.
func (m *manager) IsReadOnly() bool {
	return m.readOnly
}

// SetReadOnly sets the store to the given read-only mode.
func (m *manager) SetReadOnly(readOnly bool) {
	m.readOnly = readOnly
}

// BeginPrecompileExecution is called when a precompile is about to be executed. This function
// sets the `LogsDB` to the given `ldb` so that the `EmitEvent` and `EmitEvents` methods can
// add logs to the journal.
func (m *manager) BeginPrecompileExecution(ldb LogsDB) {
	m.ldb = ldb
}

// EndPrecompileExecution is called when a precompile has finished executing. This function
// sets the `LogsDB` to nil so that the `EmitEvent` and `EmitEvents` methods don't add logs
// to the journal.
func (m *manager) EndPrecompileExecution() {
	m.ldb = nil
}

// EmitEvent overrides the Cosmos SDK's `EventManager.EmitEvent` method to build Eth logs from
// the emitted event and add them to the journal.
func (m *manager) EmitEvent(event sdk.Event) {
	m.EventManager.EmitEvent(event)

	// add the event to the logs journal if in precompile execution
	if m.ldb != nil {
		if m.readOnly {
			panic(vm.ErrWriteProtection)
		}
		m.convertToLog(&event)
	}
}

// EmitEvents overrides the Cosmos SDK's `EventManager.EmitEvents` method to build Eth logs from
// the emitted events and add them to the journal.
func (m *manager) EmitEvents(events sdk.Events) {
	m.EventManager.EmitEvents(events)

	// add the events to the logs journal if in precompile execution
	if m.ldb != nil {
		if m.readOnly {
			panic(vm.ErrWriteProtection)
		}
		for i := range events {
			m.convertToLog(&events[i])
		}
	}
}

// Registry implements `libtypes.Registrable`.
func (m *manager) RegistryKey() string {
	return managerRegistryKey
}

// Snapshot implements `libtypes.Snapshottable`.
func (m *manager) Snapshot() int {
	return len(m.Events())
}

// RevertToSnapshot implements `libtypes.Snapshottable`.
func (m *manager) RevertToSnapshot(id int) {
	// only the events up to the snapshot id are remaining
	remaining := m.Events()[:id]

	// modify the EventManager on the underlying Cosmos SDK context
	*m.EventManager = *sdk.NewEventManager()

	// don't add to the logs db again as the Eth logs db will handle reverts, so just add the
	// remaining events on the underlying Cosmos EventManager.
	m.EventManager.EmitEvents(remaining)
}

// Finalize implements `libtypes.Finalizable`.
func (m *manager) Finalize() {}

// convertToLog builds an Eth log from the given Cosmos event and adds it to the logs journal.
func (m *manager) convertToLog(event *sdk.Event) {
	log, err := m.plf.Build(event)
	if err != nil {
		if errors.Is(err, ErrEthEventNotRegistered) {
			return
		}

		panic(errlib.Wrapf(err, "cannot convert Cosmos event %s to Eth log", event.Type))
	}

	m.ldb.AddLog(log)
}
