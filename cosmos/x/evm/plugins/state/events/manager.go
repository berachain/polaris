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
