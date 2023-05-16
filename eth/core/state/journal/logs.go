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

package journal

import (
	"context"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
)

// logs is a state plugin that tracks Ethereum logs.
type logs struct {
	journal map[common.Hash]ds.Stack[*coretypes.Log] // txHash -> journal of logs in the tx

	logIndex uint // index of the log in the block
	txHash   common.Hash
	txIndex  int
}

// NewLogs returns a new `logs` journal.
//
//nolint:revive // only used as a `state.LogsJournal`.
func NewLogs() *logs {
	return &logs{
		journal: make(map[common.Hash]ds.Stack[*coretypes.Log]),
	}
}

// Prepare implements libtypes.Preparable.
func (l *logs) Prepare(context.Context) {
	*l = *NewLogs()
}

// RegistryKey implements `libtypes.Registrable`.
func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// SetTxContext sets the transaction hash and index for the current transaction.
func (l *logs) SetTxContext(thash common.Hash, ti int) {
	l.txHash = thash
	l.txIndex = ti
	l.journal[l.txHash] = stack.New[*coretypes.Log](initCapacity)
}

func (l *logs) TxIndex() int {
	return l.txIndex
}

// AddLog adds a log to the `Logs` store.
func (l *logs) AddLog(log *coretypes.Log) {
	log.TxHash = l.txHash
	log.TxIndex = uint(l.txIndex)
	log.Index = l.logIndex
	l.journal[l.txHash].Push(log)
	l.logIndex++
}

// Logs returns the logs for the current tx with the existing metadata.
func (l *logs) Logs() []*coretypes.Log {
	size := l.journal[l.txHash].Size()
	buf := make([]*coretypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = l.journal[l.txHash].PeekAt(i)
	}
	return buf
}

// GetLogs returns the logs for the tx with the given metadata.
func (l *logs) GetLogs(txHash common.Hash, blockNumber uint64, blockHash common.Hash) []*coretypes.Log {
	size := l.journal[txHash].Size()
	buf := make([]*coretypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = l.journal[txHash].PeekAt(i)
		buf[i].BlockHash = blockHash
		buf[i].BlockNumber = blockNumber
	}
	return buf
}

// Snapshot takes a snapshot of the `Logs` store.
//
// Snapshot implements `libtypes.Snapshottable`.
func (l *logs) Snapshot() int {
	return l.journal[l.txHash].Size()
}

// RevertToSnapshot reverts the `Logs` store to a given snapshot id.
//
// RevertToSnapshot implements `libtypes.Snapshottable`.
func (l *logs) RevertToSnapshot(id int) {
	l.journal[l.txHash].PopToSize(id)
}

// Finalize clears the journal of the tx logs.
//
// Finalize implements `libtypes.Controllable`.
func (l *logs) Finalize() {}
