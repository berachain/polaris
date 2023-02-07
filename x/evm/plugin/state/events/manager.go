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

package events

import (
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

// `Manager` is a controllable event manager that can be used to Cosmos events to the Eth logs
// journal.
type Manager struct {
	// `EventManager` is the underlying Cosmos SDK event manager floating around on the context.
	*sdk.EventManager

	// semaphore chan struct{}

	// `ldb` is used to add Eth logs.
	ldb vm.LogsDB
}

// `NewManager` creates and returns a controllable event manager from the given Cosmos SDK context.
func NewManagerFrom(em sdk.EventManagerI) *Manager {
	return &Manager{
		EventManager: utils.MustGetAs[*sdk.EventManager](em),
	}
}

// `BeginPrecompileExecution` is called when a precompile is about to be executed. This function
// sets the `LogsPlugin` to the given `ldb` so that the `EmitEvent` and `EmitEvents` methods can
// add logs to the journal.
func (m *Manager) BeginPrecompileExecution(ldb vm.LogsDB) {
	m.ldb = ldb
}

// `EndPrecompileExecution` is called when a precompile has finished executing. This function
// sets the `LogsPlugin` to nil so that the `EmitEvent` and `EmitEvents` methods don't add logs
// to the journal.
func (m *Manager) EndPrecompileExecution() {
	m.ldb = nil
}

// `EmitEvent` overrides the Cosmos SDK's `EventManager.EmitEvent` method to build Eth logs from
// the emitted event and add them to the journal.
func (m *Manager) EmitEvent(event sdk.Event) {
	if m.ldb != nil {
		m.ldb.AddLog(&coretypes.Log{})
	}
	m.EventManager.EmitEvent(event)
}

// `EmitEvents` overrides the Cosmos SDK's `EventManager.EmitEvents` method to build Eth logs from
// the emitted events and add them to the journal.
func (m *Manager) EmitEvents(events sdk.Events) {
	if m.ldb != nil {
		for range events {
			m.ldb.AddLog(&coretypes.Log{})
		}
	}
	m.EventManager.EmitEvents(events)
}

// `Registry` implements `libtypes.Registrable`.
func (m *Manager) RegistryKey() string {
	return managerRegistryKey
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (m *Manager) Snapshot() int {
	return len(m.Events())
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (m *Manager) RevertToSnapshot(id int) {
	// only get the events up to the snapshot id
	revertTo := m.Events()[:id]

	// modify the EventManager on the underlying Cosmos SDK context
	*m.EventManager = *sdk.NewEventManager()

	// don't add to the logs journal again as the Eth logs plugin will do that, so use the
	// underlying EventManager to reset the events.
	m.EventManager.EmitEvents(revertTo)
}

// `Finalize` implements `libtypes.Finalizable`.
func (m *Manager) Finalize() {
	// wait for semaphore to hit 0 --> should already be at 0 at this point
}
