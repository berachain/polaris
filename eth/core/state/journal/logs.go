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
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
)

// logs is a state plugin that tracks Ethereum logs.
type logs struct {
	journal   map[common.Hash]ds.Stack[*coretypes.Log] // txHash -> journal of logs in the tx
	flattened []*coretypes.Log                         // flattened logs for the block tx

	txHash  common.Hash
	txIndex int
}

// NewLogs returns a new `logs` journal.
//
//nolint:revive // only used as a `state.LogsJournal`.
func NewLogs() *logs {
	return &logs{
		journal:   make(map[common.Hash]ds.Stack[*coretypes.Log]),
		flattened: make([]*coretypes.Log, 0, initCapacity*initCapacity),
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// SetTxContext sets the transaction hash and index for the current transaction.
func (l *logs) SetTxContext(thash common.Hash, ti int) {
	l.txHash = thash
	l.txIndex = ti
	if l.journal[thash] == nil {
		l.journal[thash] = stack.New[*coretypes.Log](initCapacity)
	}
}

// TxIndex returns the index of the current tx in the block.
func (l *logs) TxIndex() int {
	return l.txIndex
}

// AddLog adds a log to the `Logs` store.
func (l *logs) AddLog(log *coretypes.Log) {
	// add relevant metadata
	log.TxHash = l.txHash
	log.TxIndex = uint(l.txIndex)
	log.Index = uint(len(l.flattened))

	// append to journal
	if l.journal[l.txHash] == nil {
		l.journal[l.txHash] = stack.New[*coretypes.Log](initCapacity)
	}
	l.journal[l.txHash].Push(log)

	// append to block logs
	l.flattened = append(l.flattened, log)
}

// Logs returns the logs for the current tx with the existing metadata.
func (l *logs) Logs() []*coretypes.Log {
	logs := l.journal[l.txHash]
	if logs == nil {
		return nil
	}

	size := logs.Size()
	buf := make([]*coretypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = logs.PeekAt(i)
	}
	return buf
}

// GetLogs returns the logs for the tx with the given metadata.
func (l *logs) GetLogs(txHash common.Hash, blockNumber uint64, blockHash common.Hash) []*coretypes.Log {
	logs := l.journal[txHash]
	if logs == nil {
		return nil
	}

	size := logs.Size()
	buf := make([]*coretypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = logs.PeekAt(i)
		buf[i].BlockHash = blockHash
		buf[i].BlockNumber = blockNumber
	}
	return buf
}

// GetBlockLogsAndClear returns the logs for the entire block with the given blockhash and prepares
// the journal for the next block.
func (l *logs) GetBlockLogsAndClear(blockHash common.Hash) []*coretypes.Log {
	blockLogs := l.flattened
	*l = *NewLogs()
	return blockLogs
}

// Snapshot takes a snapshot of the `Logs` store.
//
// Snapshot implements `libtypes.Snapshottable`.
func (l *logs) Snapshot() int {
	if l.journal[l.txHash] == nil {
		l.journal[l.txHash] = stack.New[*coretypes.Log](initCapacity)
	}
	return l.journal[l.txHash].Size()
}

// RevertToSnapshot reverts the `Logs` store to a given snapshot id.
//
// RevertToSnapshot implements `libtypes.Snapshottable`.
func (l *logs) RevertToSnapshot(id int) {
	// guaranteed that snapshot is called before revert, so the stack for the current txHash exists
	l.journal[l.txHash].PopToSize(id)
}

// Finalize is called at the end of every state transition.
//
// Finalize implements `libtypes.Controllable`.
func (l *logs) Finalize() {}
