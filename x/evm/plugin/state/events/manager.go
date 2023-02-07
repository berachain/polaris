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
	"fmt"

	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/lib/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
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
	// semaphore chan struct{}
	// `ldb` is the reference to the StateDB for adding Eth logs during precompile execution.
	ldb precompile.LogsDB
	// `plf` is used to build Eth logs from Cosmos events.
	plf PrecompileLogFactory
	// `logger` is used to log errors when building Eth logs from Cosmos events.
	logger tmlog.Logger
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
// sets the `LogsPlugin` to the given `ldb` so that the `EmitEvent` and `EmitEvents` methods can
// add logs to the journal. It also sets the logger for logging errors from building Eth logs.
func (m *manager) BeginPrecompileExecution(ldb precompile.LogsDB, logger tmlog.Logger) {
	m.ldb = ldb
	m.logger = logger
}

// `EndPrecompileExecution` is called when a precompile has finished executing. This function
// sets the `LogsPlugin` to nil so that the `EmitEvent` and `EmitEvents` methods don't add logs
// to the journal. It also sets the logger to nil.
func (m *manager) EndPrecompileExecution() {
	m.ldb = nil
	m.logger = nil
}

// `EmitEvent` overrides the Cosmos SDK's `EventManager.EmitEvent` method to build Eth logs from
// the emitted event and add them to the journal.
func (m *manager) EmitEvent(event sdk.Event) {
	m.EventManager.EmitEvent(event)

	// add the event to the logs journal if in precompile execution
	if m.ldb != nil {
		log, err := m.plf.Build(&event)
		if err != nil {
			m.logger.Error(
				fmt.Sprintf("cannot convert Cosmos event %s to Eth log: %e\n", event.Type, err),
			)
		}
		m.ldb.AddLog(log)
	}
}

// `EmitEvents` overrides the Cosmos SDK's `EventManager.EmitEvents` method to build Eth logs from
// the emitted events and add them to the journal.
func (m *manager) EmitEvents(events sdk.Events) {
	m.EventManager.EmitEvents(events)

	// add the events to the logs journal if in precompile execution
	if m.ldb != nil {
		for i := range events {
			log, err := m.plf.Build(&events[i])
			if err != nil {
				m.logger.Error(
					fmt.Sprintf(
						"cannot convert Cosmos event %s to Eth log: %e\n", events[i].Type, err,
					),
				)
			}
			m.ldb.AddLog(log)
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
func (m *manager) Finalize() {
	// wait for semaphore to hit 0 --> should already be at 0 at this point
}
