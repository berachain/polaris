// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

type controllableManager struct {
	// TODO: better name for base?
	base sdk.EventManager

	journal ds.Stack[*sdk.Event]
}

func NewControllableManagerFrom(ctx sdk.Context) *controllableManager {
	return &controllableManager{
		base:    ctx.EventManager(),
		journal: stack.New[*sdk.Event](initJournalCapacity),
	}
}

// `Events` implements `sdk.EventManager`.
func (cm *controllableManager) Events() sdk.Events {
	size := cm.journal.Size()
	events := make(sdk.Events, size)
	for i := 0; i < size; i++ {
		events[i] = *cm.journal.PeekAt(i)
	}
	return events
}

// `ABCIEvents` implements `sdk.EventManager`.
func (cm *controllableManager) ABCIEvents() []abci.Event {
	return cm.Events().ToABCIEvents()
}

// `EmitEvent` implements `sdk.EventManager`.
func (cm *controllableManager) EmitEvent(event sdk.Event) {
	cm.journal.Push(&event)
}

// `EmitEvents` implements `sdk.EventManager`.
func (cm *controllableManager) EmitEvents(events sdk.Events) {
	for _, event := range events {
		cm.journal.Push(&event)
	}
}

// `EmitTypedEvent` implements `sdk.EventManager`.
func (cm *controllableManager) EmitTypedEvent(tev proto.Message) error {
	event, err := sdk.TypedEventToEvent(tev)
	if err != nil {
		return err
	}
	cm.EmitEvent(event)
	return nil
}

// `EmitTypedEvents` implements `sdk.EventManager`.
func (cm *controllableManager) EmitTypedEvents(tevs ...proto.Message) error {
	events := make(sdk.Events, len(tevs))
	for i, tev := range tevs {
		res, err := sdk.TypedEventToEvent(tev)
		if err != nil {
			return err
		}
		events[i] = res
	}

	cm.EmitEvents(events)
	return nil
}

// `Registry` implements `libtypes.Registrable`.
func (cm *controllableManager) RegistryKey() string {
	return managerRegistryKey
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (cm *controllableManager) Snapshot() int {
	return cm.journal.Size()
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (cm *controllableManager) RevertToSnapshot(id int) {
	cm.journal.PopToSize(id)
}

// `Finalize` implements `libtypes.Finalizable`.
func (cm *controllableManager) Finalize() {
	size := cm.journal.Size()
	eventsToEmit := make(sdk.Events, size)
	for i := 0; i < size; i++ {
		eventsToEmit[i] = *cm.journal.PeekAt(i)
	}
	cm.base.EmitEvents(eventsToEmit)

	// clear journal
	cm.journal.PopToSize(0)
}
