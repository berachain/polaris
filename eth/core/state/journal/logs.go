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
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Log defines the interface for tracking logs created during a state transition.
type Log interface {
	// Log implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// Log implements `libtypes.Cloneable`.
	libtypes.Cloneable[Log]
	// SetTxContext sets the transaction hash and index for the current transaction.
	SetTxContext(thash common.Hash, ti int)
	// TxIndex returns the current transaction index.
	TxIndex() int
	// AddLog adds a log to the logs journal.
	AddLog(*ethtypes.Log)
	// Logs returns the logs of the tx with the existing metadata.
	Logs() []*ethtypes.Log
	// GetLogs returns the logs of the tx with the given metadata.
	GetLogs(hash common.Hash, blockNumber uint64, blockHash common.Hash) []*ethtypes.Log
}

// logs is a state plugin that tracks Ethereum logs.
type logs struct {
	baseJournal[*ethtypes.Log]
	txHash  common.Hash
	txIndex int
	logSize int
}

// NewLogs returns a new `logs` journal.
func NewLogs() Log {
	return &logs{
		baseJournal: newBaseJournal[*ethtypes.Log](initCapacity),
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// SetTxContext sets the transaction hash and index for the current transaction.
func (l *logs) SetTxContext(thash common.Hash, ti int) {
	l.baseJournal = newBaseJournal[*ethtypes.Log](initCapacity)
	// Set the transaction hash and index.
	l.txHash = thash
	l.txIndex = ti
}

// TxIndex returns the current transaction index.
func (l *logs) TxIndex() int {
	return l.txIndex
}

// AddLog adds a log to the `Logs` store.
func (l *logs) AddLog(log *ethtypes.Log) {
	log.TxHash = l.txHash
	log.TxIndex = uint(l.txIndex)
	log.Index = uint(l.logSize) + uint(l.Size())
	l.Push(log)
}

// Logs returns the logs for the current tx with the existing metadata.
func (l *logs) Logs() []*ethtypes.Log {
	size := l.Size()
	buf := make([]*ethtypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = l.PeekAt(i)
	}
	return buf
}

// GetLogs returns the logs for the tx with the given metadata.
func (l *logs) GetLogs(_ common.Hash, blockNumber uint64, blockHash common.Hash) []*ethtypes.Log {
	size := l.Size()
	buf := make([]*ethtypes.Log, size)
	for i := 0; i < size; i++ {
		buf[i] = l.PeekAt(i)
		buf[i].BlockHash = blockHash
		buf[i].BlockNumber = blockNumber
	}
	return buf
}

// Finalize clears the journal of the tx logs.
//
// Finalize implements `libtypes.Controllable`.
func (l *logs) Finalize() {
	// increase in finalize based on the final size of the
	// logs
	l.logSize += l.Size()
}

// Clone implements `libtypes.Cloneable`.
func (l *logs) Clone() Log {
	clone := &logs{
		baseJournal: newBaseJournal[*ethtypes.Log](l.Capacity()),
		txHash:      l.txHash,
		txIndex:     l.txIndex,
		logSize:     l.logSize,
	}

	// copy every individual log from the journal
	for i := 0; i < l.Size(); i++ {
		cpy := new(ethtypes.Log)
		*cpy = *l.PeekAt(i)
		clone.Push(cpy)
	}

	return clone
}
