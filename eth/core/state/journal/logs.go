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

package journal

import (
	"math"

	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
)

// `logs` is a state plugin that tracks Ethereum logs.
type logs struct {
	// Reset every tx.
	ds.Stack[*coretypes.Log] // journal of tx logs
}

// `NewLogs` returns a new `logs` journal.
//
//nolint:revive // only used as a `state.LogsJournal`.
func NewLogs() *logs {
	return &logs{
		Stack: stack.New[*coretypes.Log](initJournalCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (l *logs) RegistryKey() string {
	return logsRegistryKey
}

// `AddLog` adds a log to the `Logs` store.
func (l *logs) AddLog(log *coretypes.Log) {
	l.Push(log)
}

// `BuildLogsAndClear` returns the logs for the tx with the given metadata.
func (l *logs) BuildLogsAndClear(
	txHash common.Hash,
	blockHash common.Hash,
	txIndex uint,
	logIndex uint,
) []*coretypes.Log {
	size := l.Size()
	buf := make([]*coretypes.Log, size)
	for i := uint(size) - 1; i < math.MaxUint; i-- {
		buf[i] = l.Pop()
		buf[i].TxHash = txHash
		buf[i].BlockHash = blockHash
		buf[i].TxIndex = txIndex
		buf[i].Index = logIndex + i
	}
	return buf
}

// `Snapshot` takes a snapshot of the `Logs` store.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (l *logs) Snapshot() int {
	return l.Size()
}

// `RevertToSnapshot` reverts the `Logs` store to a given snapshot id.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (l *logs) RevertToSnapshot(id int) {
	l.PopToSize(id)
}

// `Finalize` implements `libtypes.Controllable`.
func (l *logs) Finalize() {
	l.Stack = stack.New[*coretypes.Log](initJournalCapacity)
}
