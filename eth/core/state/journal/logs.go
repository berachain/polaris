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
