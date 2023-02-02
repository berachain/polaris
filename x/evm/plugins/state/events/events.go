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
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

type controllableManager struct {
	// TODO: better names for these?
	base *sdk.EventManager // only used to finalize to "parent" ctx EventManager
	temp *sdk.EventManager // current, most valid event manager
}

// `NewControllableManagerFrom` creates and returns a controllable event manager from the given
// Cosmos SDK context `ctx`.
//
//nolint:revive // should only be used as a `state.ControllableEventManager`.
func NewControllableManager(base *sdk.EventManager) *controllableManager {
	return &controllableManager{
		base: base,
		temp: sdk.NewEventManager(),
	}
}

// `EventManager` implements `state.ControllableEventManager`.
func (cm *controllableManager) EventManager() *sdk.EventManager {
	return cm.temp
}

// `Registry` implements `libtypes.Registrable`.
func (cm *controllableManager) RegistryKey() string {
	return managerRegistryKey
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (cm *controllableManager) Snapshot() int {
	return len(cm.temp.Events())
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (cm *controllableManager) RevertToSnapshot(id int) {
	// create new event manager with only the first `id` events
	newTemp := sdk.NewEventManager()
	newTemp.EmitEvents(cm.temp.Events()[:id])
	cm.temp = newTemp
}

// `Finalize` implements `libtypes.Finalizable`.
func (cm *controllableManager) Finalize() {
	cm.base.EmitEvents(cm.temp.Events())
}
