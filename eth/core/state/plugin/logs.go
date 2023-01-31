// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package plugin

import (
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
)

const (
	// `initLogCapacity` is the initial capacity of the `logs` snapshot stack.
	initLogCapacity = 16
	logsRegistryKey = `logs`
)

// `logs` is a state plugin that tracks Ethereum logs.
type logs struct {
	// For the block.
	txHashToLogs map[common.Hash]ds.Stack[*coretypes.Log]
	logSize      uint

	// Reset every tx.
	currentTxHash common.Hash
	currenTxIndex uint
}

// `NewLogs` returns a new `Logs` store.
func NewLogs() *logs {
	return &logs{
		txHashToLogs:  make(map[common.Hash]ds.Stack[*coretypes.Log]),
		currentTxHash: common.Hash{},
		currenTxIndex: 0,
	}
}

func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// `Prepare` prepares the `Logs` store for a new transaction.
func (l *logs) Prepare(txHash common.Hash, ti uint) {
	l.currentTxHash = txHash
	l.currenTxIndex = ti
	l.txHashToLogs[l.currentTxHash] = stack.New[*coretypes.Log](initLogCapacity)
}

// `AddLog` adds a log to the `Logs` store.
func (l *logs) AddLog(log *coretypes.Log) {
	logs := l.txHashToLogs[l.currentTxHash]
	log.TxHash = l.currentTxHash
	log.TxIndex = l.currenTxIndex
	log.Index = l.logSize
	logs.Push(log)
	l.logSize++
}

// `GetLogs` returns the Logs for a given transaction hash.
func (l *logs) GetLogs(txHash common.Hash, blockHash common.Hash) []*coretypes.Log {
	logs := l.txHashToLogs[txHash]
	size := logs.Size()
	buf := make([]*coretypes.Log, size)
	for i := 0; i < logs.Size(); i++ {
		buf[i] = logs.PeekAt(i)
		buf[i].BlockHash = blockHash
	}
	return buf
}

// `Snapshot` takes a snapshot of the `Logs` store.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (l *logs) Snapshot() int {
	return l.txHashToLogs[l.currentTxHash].Size()
}

// `RevertToSnapshot` reverts the `Logs` store to a given snapshot.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (l *logs) RevertToSnapshot(i int) {
	l.txHashToLogs[l.currentTxHash].PopToSize(i)
}

func (l *logs) Finalize() error {
	return nil
}
