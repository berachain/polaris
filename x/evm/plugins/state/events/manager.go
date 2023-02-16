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
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

// `manager` is a controllable event manager that supports snapshots and reverts for emitted Cosmos
// events. During precompile execution, it is also used to Cosmos events to the Eth logs journal.
type manager struct {
	// `EventManager` is the underlying Cosmos SDK event manager floating around on the context.
	*sdk.EventManager
	// `ldb` is the reference to the StateDB for adding Eth logs during precompile execution.
	ldb LogsDB
	// `plf` is used to build Eth logs from Cosmos events.
	plf PrecompileLogFactory
}

// `NewManager` creates and returns a controllable event manager from the given Cosmos SDK context.
//
//nolint:revive // only used as a `state.ControllableEventManager`.
func NewManagerFrom(em sdk.EventManagerI, plf PrecompileLogFactory) *manager {
	return &manager{
		EventManager: utils.MustGetAs[*sdk.EventManager](em),
		plf:          plf,
	}
}

// `BeginPrecompileExecution` is called when a precompile is about to be executed. This function
// sets the `LogsDB` to the given `ldb` so that the `EmitEvent` and `EmitEvents` methods can
// add logs to the journal.
func (m *manager) BeginPrecompileExecution(ldb LogsDB) {
	m.ldb = ldb
}

// `EndPrecompileExecution` is called when a precompile has finished executing. This function
// sets the `LogsDB` to nil so that the `EmitEvent` and `EmitEvents` methods don't add logs
// to the journal.
func (m *manager) EndPrecompileExecution() {
	m.ldb = nil
}

// `EmitEvent` overrides the Cosmos SDK's `EventManager.EmitEvent` method to build Eth logs from
// the emitted event and add them to the journal.
func (m *manager) EmitEvent(event sdk.Event) {
	m.EventManager.EmitEvent(event)

	// add the event to the logs journal if in precompile execution
	if m.ldb != nil {
		m.convertToLog(&event)
	}
}

// `EmitEvents` overrides the Cosmos SDK's `EventManager.EmitEvents` method to build Eth logs from
// the emitted events and add them to the journal.
func (m *manager) EmitEvents(events sdk.Events) {
	m.EventManager.EmitEvents(events)

	// add the events to the logs journal if in precompile execution
	if m.ldb != nil {
		for i := range events {
			m.convertToLog(&events[i])
		}
	}
}

// `Registry` implements `libtypes.Registrable`.
func (m *manager) RegistryKey() string {
	return managerRegistryKey
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (m *manager) Snapshot() int {
	return len(m.Events())
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (m *manager) RevertToSnapshot(id int) {
	// only get the events up to the snapshot id
	revertTo := m.Events()[:id]

	// modify the EventManager on the underlying Cosmos SDK context
	*m.EventManager = *sdk.NewEventManager()

	// don't add to the logs journal again as the Eth logs plugin will do that, so use the
	// underlying EventManager to reset the events.
	m.EventManager.EmitEvents(revertTo)
}

// `Finalize` implements `libtypes.Finalizable`.
func (m *manager) Finalize() {}

// `convertToLog` builds an Eth log from the given Cosmos event and adds it to the logs journal.
func (m *manager) convertToLog(event *sdk.Event) {
	log, err := m.plf.Build(event)
	if err != nil {
		panic(errors.Wrapf(err, "cannot convert Cosmos event %s to Eth log", event.Type))
	}
	m.ldb.AddLog(log)
}
